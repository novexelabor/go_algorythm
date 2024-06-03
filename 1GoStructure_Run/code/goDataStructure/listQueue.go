
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

import "fmt"

type ListQueue struct {
	il *ImproveList
}

func CreateListQueue() *ListQueue {
	return &ListQueue{
		il: CreateImproveList(),
	}
}

func (lq *ListQueue) DeQueue() interface{} {
	return lq.il.RemoveFirst()
}

func (lq *ListQueue) EnQueue(e interface{}) {
	lq.il.AddLast(e)
}

func (lq *ListQueue) GetFront() interface{} {
	return lq.il.GetFirst()
}

func (lq *ListQueue) GetSize() int {
	return lq.il.GetSize()
}

func (lq *ListQueue) IsEmpty() bool {
	return lq.il.IsEmpty()
}

func (lq *ListQueue) String() string {
	str := fmt.Sprint("Queue:")
	str += fmt.Sprint("front [")
	head := lq.il.head
	for {
		str += fmt.Sprint(head.value)
		if head.next != nil {
			str += fmt.Sprint("->")
			head = head.next
		} else {
			break
		}
	}
	str += fmt.Sprint("] tail")
//fmt.println（字符串）
	return str
}
