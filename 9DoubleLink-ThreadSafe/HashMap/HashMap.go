package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

type rindex []uint32 //hash环索引

type ring struct {
	rmap      map[uint32]string //环结构体,字典map
	rindexarr rindex            //索引
	lock      *sync.RWMutex     //线程安全
}

//比大小
func (ridx rindex) Less(i, j int) bool {
	return ridx[i] < ridx[j]
}

//长度
func (ridx rindex) Len() int {
	return len(ridx)
}

//交换
func (ridx rindex) Swap(i, j int) {
	ridx[i], ridx[j] = ridx[j], ridx[i]
}

func (ridx *ring) Addnode(nodename string) {
	ridx.lock.Lock()
	defer ridx.lock.Unlock()
	//uint32
	index := crc32.ChecksumIEEE([]byte(nodename)) //sha256加密算法
	if _, ok := ridx.rmap[index]; ok {            //表示已存在
		return //返回
	}
	ridx.rmap[index] = nodename                    //赋值
	ridx.rindexarr = append(ridx.rindexarr, index) //加载索引
	sort.Sort(ridx.rindexarr)                      //排序

}
func (ridx *ring) Removenode(nodename string) {
	ridx.lock.Lock()
	defer ridx.lock.Unlock()
	index := crc32.ChecksumIEEE([]byte(nodename)) //sha256
	if _, ok := ridx.rmap[index]; !ok {
		return //返回
	}
	delete(ridx.rmap, index) //删除map内置数据
	ridx.rindexarr = rindex{}
	for k := range ridx.rmap { //重新遍历再排序
		ridx.rindexarr = append(ridx.rindexarr, k) //插入数据
	}
	sort.Sort(ridx.rindexarr) //排序
}
func (ridx *ring) Getnode(nodename string) string {
	ridx.lock.RLock() //其他线程可以读取，不可以修改
	defer ridx.lock.RUnlock()
	//查找
	hash := crc32.ChecksumIEEE([]byte(nodename))
	//sort.Search函数
	i := sort.Search(len(ridx.rindexarr), func(i int) bool {
		return ridx.rindexarr[i] == hash
	})
	if i < 0 || i > len(ridx.rindexarr)-1 {
		return ""
	}
	node := ridx.rmap[ridx.rindexarr[i]] //取得节点
	return node

}
func main() {
	filelist := []string{"123", "456", "789"}
	//创建哈希字典
	hashmap := &ring{map[uint32]string{}, rindex{}, new(sync.RWMutex)}
	fmt.Println(filelist, hashmap)
	for _, v := range filelist {
		index := crc32.ChecksumIEEE([]byte(v))               //循环索引
		hashmap.rmap[index] = v                              //设定索引
		hashmap.rindexarr = append(hashmap.rindexarr, index) //插入

	}
	//处理索引数组
	sort.Sort(hashmap.rindexarr)
	fmt.Println(hashmap)

	//fmt.Println(hashmap.Getnode("xadasd"))

	hashmap.Addnode("xyz123")
	fmt.Println(hashmap)

	hashmap.Removenode("xyz123")
	fmt.Println(hashmap)

	mystr := hashmap.Getnode("789123")
	if mystr == "" {
		fmt.Println("找不到")
	} else {
		fmt.Println("找到", mystr)
	}
}
