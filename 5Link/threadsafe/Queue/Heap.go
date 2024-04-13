package Queue

import (
	"sync"
)

type Int int

func (x Int) Less(than Item) bool {
	return x < than.(Int)
}

//接口实现比大小
type Item interface { //能够比较大小的数据类型，就实现了该接口
	Less(than Item) bool
}

//最小堆  /最大堆
type Heap struct {
	lock *sync.Mutex //线程安全
	data []Item      //数组，Int类型实现了Item接口,PriorityItem类型
	min  bool        //是否
}

//标准堆
func NewHeap() *Heap {
	return &Heap{new(sync.Mutex), make([]Item, 0), true}
}

//最小堆
func NewMin() *Heap {
	return &Heap{new(sync.Mutex), make([]Item, 0), true}
}

//最大堆
func NewMax() *Heap {
	return &Heap{new(sync.Mutex), make([]Item, 0), false}
}

//判断是否为空
func (h *Heap) isEmpty() bool {
	return len(h.data) == 0
}

//求长度
func (h *Heap) Len() int {
	return len(h.data)
}

//抓取数据
func (h *Heap) Get(index int) Item {
	return h.data[index]
}

//插入数据
func (h *Heap) Insert(it Item) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.data = append(h.data, it) //插入数据
	h.siftUp()                  //插入数据,无论最大堆还是最小堆都是向上调整，每次插入也是调整堆
	return
}

//根据类型返回比大小,根据最大最小标志，来进行比较数据的大小
func (h *Heap) Less(a, b Item) bool {
	if h.min {
		return a.Less(b)
	} else {
		return b.Less(a) //b是否比a小
	}
}

//压缩，弹出一个
func (h *Heap) Extract() Item {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.Len() == 0 {
		return nil //长度为0.不需要处理
	}
	el := h.data[0]
	last := h.data[h.Len()-1] //最后一个
	if h.Len() == 1 {
		h.data = nil //重新分配内存
		return nil
	}
	h.data = append([]Item{last}, h.data[1:h.Len()-1]...) //将最后一个数据放到最前面
	h.siftDown()                                          //调整堆,同样无论最大堆还是最小堆都是，弹出调整
	return el                                             //返回第一个数据

}

//弹出一个极大值,从最后的节点出发的
func (h *Heap) siftUp() {
	//堆排序的循环过程  n,2n+1
	for i, parent := h.Len()-1, h.Len()-1; i > 0; i = parent {
		parent = i / 2
		if h.Less(h.Get(i), h.Get(parent)) { //处理数据
			h.data[parent], h.data[i] = h.data[i], h.data[parent]
		} else {
			break
		}

	}
}

//弹出一个极小值，从第一个出发的,比较左右孩子节点大小
func (h *Heap) siftDown() {
	//堆排序的循环过程  n,2n+1
	for i, child := 0, 1; i < h.Len() && i*2+1 < h.Len(); i = child {
		child = i*2 + 1
		if child+1 <= h.Len()-1 && h.Less(h.Get(child+1), h.Get(child)) {
			child++ //循环左右节点过程,有孩子节点更小
		}
		if h.Less(h.Get(i), h.Get(child)) {
			break
		}
		h.data[i], h.data[child] = h.data[child], h.data[i] //处理数据交换
	}
}
