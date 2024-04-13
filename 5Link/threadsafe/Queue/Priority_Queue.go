package Queue

type PriorityItem struct {
	value    interface{} //数据
	Priority int         //优先级
}

//新建队列的元素
func NewPriorityItem(value interface{}, Priority int) *PriorityItem {
	return &PriorityItem{value, Priority}
}

//比大小设定比较优先级
func (x PriorityItem) Less(than Item) bool { //PriorityItem类型实现了Item接口
	return x.Priority < than.(PriorityItem).Priority
}

//优先队列，基于堆
type PriorityQueue struct {
	data *Heap
}

func NewMaxPriorityQueue() *PriorityQueue {
	return &PriorityQueue{NewMax()}
}
func NewMinPriorityQueue() *PriorityQueue {
	return &PriorityQueue{NewMin()}
}

func (pq *PriorityQueue) Len() int {
	return pq.data.Len() //队列长度
}

func (pq *PriorityQueue) Insert(el PriorityItem) {
	pq.data.Insert(el)
}
func (pq *PriorityQueue) Extract() PriorityItem {
	return pq.data.Extract().(PriorityItem)
}
func (pq *PriorityQueue) ChangePriority(val interface{}, Priority int) {
	var storage = NewQueue() //队列备份数据
	popped := pq.Extract()   //拿出最小的数值
	for val != popped.value {
		if pq.Len() == 0 {
			return
		}
		storage.Push(popped) //压入数据
		popped = pq.Extract()
	}
	popped.Priority = Priority //修改优先级
	pq.data.Insert(popped)     //插入数据
	for storage.Len() > 0 {
		pq.data.Insert(storage.Shift().(Item)) //其余数据重新放入优先队列
	}
}
