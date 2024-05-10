package main

import (
	"fmt"
	"sync"
)

//循环双端队列
type Deque struct {
	array       []interface{}
	left, right int //记录前后的下标
	num         int
	lock        *sync.RWMutex //线程安全
}

func NewDeque(cap int) *Deque {
	if cap <= 0 {
		panic("队列容量必须大于0")
	}
	deq := new(Deque)                    //新建一个队列结构体，初始化
	deq.array = make([]interface{}, cap) //开辟内存
	deq.num = 0
	deq.left = 0
	deq.right = 0
	deq.lock = new(sync.RWMutex)
	return deq

}

//-----------
//左端添加数据
func (deq *Deque) Addleft(data interface{}) {
	if deq.num == len(deq.array) {
		panic("overflow") //超出了
	}
	deq.array[deq.left] = data
	deq.num++
	deq.left = deq.left - 1
	if deq.left == -1 {
		deq.left = len(deq.array) - 1 //循环双端队列，变成切片数组的下标最大
	}

}
func (deq *Deque) Addright(data interface{}) {
	if deq.num == len(deq.array) {
		panic("overflow")
	}
	deq.array[deq.right] = data
	deq.num++
	deq.right = deq.right + 1 //循环，没有处理超出的情况
	if deq.right == len(deq.array) {
		deq.right = 0
	}

}
func (deq *Deque) Delleft() interface{} {
	if deq.num == 0 {
		panic("overflow")
	}
	deq.left = deq.left + 1
	if deq.left == len(deq.array) { //循环到尾部
		deq.left = 0
	}
	data := deq.array[deq.left]
	deq.num--
	return data
}
func (deq *Deque) Delright() interface{} {
	if deq.num == 0 {
		panic("overflow")
	}
	deq.right = deq.right - 1 //循环
	if deq.right == -1 {
		deq.right = len(deq.array) - 1 //循环
	}
	data := deq.array[deq.right] //返回数据
	deq.num--
	return data
}

func main() {
	deq := NewDeque(3)
	//deq.Addleft(1)
	deq.Addright(1)
	fmt.Println(deq.left, deq.right, deq.array)
	//deq.Addleft(2)
	deq.Addright(2)
	fmt.Println(deq.left, deq.right, deq.array)
	//deq.Addleft(3)
	deq.Addright(3)
	fmt.Println(deq.left, deq.right, deq.array)

	// deq.Addright(4)
	// fmt.Println(deq.left, deq.right, deq.array)

	//deq.Delleft()
	fmt.Println(deq.Delright())
	fmt.Println(deq.left, deq.right, deq.array)
	//deq.Delleft()
	fmt.Println(deq.Delright())
	fmt.Println(deq.left, deq.right, deq.array)
	//deq.Delleft()
	fmt.Println(deq.Delright())
	fmt.Println(deq.left, deq.right, deq.array)
}
