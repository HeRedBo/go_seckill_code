package repositories

import (
	"IrisMovies/datamodels"
	"errors"
	"sync"
)

// Query代表一种“访客”和它的查询动作
type Query func(datamodels.Movie) bool

// MovieRepository会处理一些关于movie实例的基本的操作 。
// 这是一个以测试为目的的接口，即是一个内存中的movie库
// 或是一个连接到数据库的实例。

type MovieRepository interface {

	//  执行方法
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	// 查询相关方法
	Select(query Query) (movie datamodels.Movie, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Movie)

	// 新增编辑删除
	InsertOrUpdate(movie datamodels.Movie) (updateMovie datamodels.Movie, err error)
	Delete(query Query, limit int) (deleted bool)

	GetMovieName() string
}

// movieMemoryRepository就是一个"MovieRepository"
// 它负责存储于内存中的实例数据(map)
type movieMemoryRepository struct {
	source map[int64]datamodels.Movie
	mu sync.RWMutex
}
// NewMovieRepository返回一个新的基于内存的movie库。
// 库的类型在我们的例子中是唯一的。

func NewMovieRepository(source map[int64]datamodels.Movie) MovieRepository {
	return &movieMemoryRepository{source:source}
}

const (
	// ReadOnlyMode will RLock(read) the data .
	ReadOnlyMode = iota
	// ReadWriteMode will Lock(read/write) the data.
	ReadWriteMode
)

// 电影仓库方法
func (m *movieMemoryRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0
	if mode == ReadOnlyMode {
		m.mu.RLock()
		defer m.mu.RUnlock()
	} else {
		m.mu.Lock()
		defer m.mu.Unlock()
	}

	for _, movie := range m.source {
		ok = query(movie)
		if ok {
			if action(movie) {
				loops++
				if actionLimit >= loops {
					break
				}
			}
		}
	}
	return
}

// Select方法会收到一个查询方法
// 这个方法给出一个单独的movie实例
// 直到这个功能返回为true时停止迭代。
//
// 它返回最后一次查询成功所找到的结果的值
// 和最后的movie模型
// 以减少caller之间的通信
func (r *movieMemoryRepository) Select(query Query) (movie datamodels.Movie, found bool) {
	found = r.Exec(query, func(m datamodels.Movie) bool {
		movie = m
		return true
	},1, ReadOnlyMode)

	// set an empty datamodels.Movie if not found at all.
	if !found {
		movie = datamodels.Movie{}
	}
	return
}

func (r *movieMemoryRepository) SelectMany(query Query, limit int) (results []datamodels.Movie)  {
	r.Exec(query, func(m datamodels.Movie) bool {
		results = append(results,m)
		return true
	},limit ,ReadOnlyMode)
	return
}

func (r *movieMemoryRepository) InsertOrUpdate(movie datamodels.Movie) (datamodels.Movie, error)  {
	id := movie.ID

	// 数据更新操作
	if id == 0 {
		var lastID int64
		// 找到最大的ID，避免重复。
		// 在实际使用时您可以使用第三方库去生成
		// 一个string类型的UUID
		r.mu.RLock()
		for _,item := range r.source {
			if item.ID > lastID {
				lastID = item.ID
			}
		}
		r.mu.RUnlock()

		id = lastID +1
		movie.ID = id
		// map-specific thing
		r.mu.Lock()
		r.source[id] = movie
		r.mu.Unlock()
		return movie,nil
	}

	current, exists := r.Select(func(m datamodels.Movie) bool {
		return m.ID == id
	})

	if !exists { // 当ID不存在时抛出一个error
		return datamodels.Movie{}, errors.New("failed to update a nonexistent movie")
	}

	if movie.Poster != "" {
		current.Poster = movie.Poster
	}

	if movie.Genre != "" {
		current.Genre = movie.Genre
	}

	// 数据替换
	r.mu.Lock()
	r.source[id] = current
	r.mu.Unlock()
	return movie ,nil
}


func (r *movieMemoryRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(m datamodels.Movie) bool {
		delete(r.source, m.ID)
		return true
	},limit, ReadWriteMode)
}


func (m *movieMemoryRepository ) GetMovieName() string {
	movie := &datamodels.Movie{Name:"胖墩墩奥运奇遇记"}
	return movie.Name
}

