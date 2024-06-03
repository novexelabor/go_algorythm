
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

import "fmt"

type ImproveList struct {
	head *node
	tail *node
	size int
}

func CreateImproveList() *ImproveList {
	return &ImproveList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (il *ImproveList) AddHead(e interface{}) {
	il.head = &node{
		value: e,
		next:  il.head,
	}
	if il.IsEmpty() {
		il.tail = il.head
	}
	il.size++
}
func (il *ImproveList) AddLast(e interface{}) {
	if il.IsEmpty() {
		il.AddHead(e)
		return
	}
	il.tail.next = &node{
		value: e,
		next:  nil,
	}
	il.tail = il.tail.next
	il.size++
}

func (il *ImproveList) RemoveFirst() interface{} {
	if il.IsEmpty() {
		panic("lian biao wei kong")
	}
	ret := il.head
	il.head = ret.next
	ret.next = nil
	il.size--
	if il.IsEmpty() {
		il.tail = nil
	}
	return ret.value
}
func (il *ImproveList) GetFirst() interface{} {
	return il.head.value
}
func (il *ImproveList) GetLast() interface{} {
	return il.tail.value
}
func (il *ImproveList) GetSize() int {
	return il.size
}

func (il *ImproveList) IsEmpty() bool {
	return il.size == 0
}

func (il *ImproveList) Contains(e interface{}) bool {
	current := il.head
	for current != nil {
		if current.value == e {
return true //附带条件
		}
		current = current.next
	}
	return false
}
func (il *ImproveList) String() string {
	if il.IsEmpty() {
		return fmt.Sprint("List: nil")
	}
	str := fmt.Sprint("List:")
	str += fmt.Sprint("Head ")
	current := il.head
	for {
		str += fmt.Sprint(current.value)
		if current.next != nil {
			str += fmt.Sprint("->")
			current = current.next
		} else {
			break
		}
	}
	str += fmt.Sprint("->nil")
	return str
}
