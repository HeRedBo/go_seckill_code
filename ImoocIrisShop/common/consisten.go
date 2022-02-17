package common

import (
	"container/list"
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// 声明新切片类型
type units []uint32

// 返回切片差评
func (x units) Len() int {
	return len(x)
}

// 比较两个数大小
func (x units) Less(i, j int) _bool {
	return x[i] < x[j]
}

// 切片两个值交换
func (x units) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

var errEmpty = errors.New("Hash 环没有数据")

// 当hash 环没有数据时，提示错误
type Consistent struct {
	// hash 环 ，key 为hash 值，值存放节点的信息
	circle map[uint32]string

	sortedHashes units

	VirtualNode int

	sync.RWMutex
}

func NewConsistent() *Consistent {
	return &Consistent{
		circle: make(map[uint32]string),

		VirtualNode: 20,
	}
}

func (c *Consistent) generateKey(element string, index int) string {
	return element + strconv.Itoa(index)
}

func (c *Consistent) hashkey(key string) uint32 {
	if len(key) < 64 {
		//声明一个数组长度为64
		var srcatch [64]byte
		copy(srcatch[:], key)
		//使用IEEE 多项式返回数据的CRC-32校验和
		return crc32.ChecksumIEEE(srcatch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

// 更新排序 方便查找
func (c *Consistent) updateSortedHashes() {
	hashes := c.sortedHashes[:0]

	if cap(c.sortedHashes)/(c.VirtualNode*4) > len(c.circle) {
		hashes = nil
	}
	for k := range c.circle {
		hashes = append(hashes, k)
	}

	sort.Sort(hashes)
	c.sortedHashes = hashes
}

func (c *Consistent) Add(element string) {
	c.RLock()
	defer c.RUnlock()
	c.add(element)
}

func (c *Consistent) add(element string) {
	for i := 0; i < c.VirtualNode; i++ {
		c.circle[c.hashkey(c.generateKey(element, i))] = element
	}
	// 更新排序
	c.updateSortedHashes()
}

//删除节点
func (c *Consistent) remove(element string) {
	for i := 0; i < c.VirtualNode; i++ {
		delete(c.circle, c.hashkey(c.generateKey(element, i)))
	}
	c.updateSortedHashes()
}

//删除一个节点
func (c *Consistent) Remove(element string) {
	c.Lock()
	defer c.Unlock()
	c.remove(element)
}

//顺时针查找最近的节点
func (c *Consistent) search(key uint32) int {
	//查找算法
	f := func(x int) bool {
		return c.sortedHashes[x] > key
	}
	//使用"二分查找"算法来搜索指定切片满足条件的最小值
	i := sort.Search(len(c.sortedHashes), f)
	//如果超出范围则设置i=0
	if i >= len(c.sortedHashes) {
		i = 0
	}
	return i
}

//根据数据标示获取最近的服务器节点信息
func (c *Consistent) Get(name string) (string, error) {
	//添加锁
	c.RLock()
	//解锁
	defer c.RUnlock()
	//如果为零则返回错误
	if len(c.circle) == 0 {
		return "", errEmpty
	}
	//计算hash值
	key := c.hashkey(name)
	i := c.search(key)
	return c.circle[c.sortedHashes[i]], nil
}
