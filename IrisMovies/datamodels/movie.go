package datamodels

// Movie 是基本数据的的结构体.
// 请注意公共标签（适用于我们的 web 应用）
// 应保存在 「web / viewmodels / movie.go」等其他文件中
// 它可以通过嵌入数据模型进行换行。
//电影或声明新的字段，但我们将使用此数据模型作为我们的应用程序
//中唯一的一个电影模型，为了简单起见。
type Movie struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Year int `json:"year"`
	Genre string `json:"genre"`
	Poster string `json:"poster"`
}
