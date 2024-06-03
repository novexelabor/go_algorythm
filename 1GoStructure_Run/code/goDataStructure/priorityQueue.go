
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

type PriorityQueue struct {
	mh *MaxHeap
}

func CreatePriorityQueue() *PriorityQueue {

	return &PriorityQueue{
		CreateMaxHeap(),
	}
}

func (pq *PriorityQueue) DeQueue() interface{} {
	return pq.mh.Remove()
}

func (pq *PriorityQueue) EnQueue(e interface{}) {
	ce, ok := e.(Comparable)
	if !ok {
		panic("e is no a comparable ver")
	}
	pq.mh.Add(ce)
}

func (pq *PriorityQueue) GetFront() interface{} {
	return pq.mh.getMax()
}

func (pq *PriorityQueue) GetSize() int {
	return pq.mh.GetSize()
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.mh.IsEmpty()
}

func (pq *PriorityQueue) String() string {
	return pq.mh.String()
}
