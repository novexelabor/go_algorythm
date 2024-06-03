
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

import "fmt"

type ArrayQueue struct {
	a *arr
}

//无检测功能
func CreateArrayQueue(cap int) *ArrayQueue {
	return &ArrayQueue{CreateArray(cap)}
}

func (q *ArrayQueue) GetSize() int {
	return q.a.GetSize()
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.a.IsEmpty()
}

func (q *ArrayQueue) EnQueue(e interface{}) {
	q.a.AddLast(e)
}

func (q *ArrayQueue) DeQueue() interface{} {
	return q.a.RemoveFirst()
}

func (q *ArrayQueue) GetFront() interface{} {
	return q.a.GetFirst()
}
func (q *ArrayQueue) GetCap() int {
	return q.a.GetCap()
}
func (q *ArrayQueue) String() string {
	str := fmt.Sprint("Queue:")
	str += fmt.Sprint(("front ["))
	for i := 0; i < q.a.size; i++ {
		str += fmt.Sprint((*q).a.data[i])
		if i != q.a.size-1 {
			str += fmt.Sprint(",")
		}
	}
	str += fmt.Sprint("] tail")
//fmt.println（字符串）
	return str
}
