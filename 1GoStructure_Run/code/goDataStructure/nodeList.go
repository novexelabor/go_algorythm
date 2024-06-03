
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

import (
	"fmt"
)

type NodeList struct {
	dummyHead *node
	size      int
}
type node struct {
	value interface{}
	next  *node
}

func CreateNodeList() *NodeList {
	return &NodeList{
		dummyHead: &node{
			value: nil,
			next:  nil,
		},
		size: 0,
	}
}

func (nl *NodeList) GetSize() int {
	return nl.size
}

func (nl *NodeList) IsEmpty() bool {
	return nl.size == 0
}

func (nl *NodeList) AddHead(e interface{}) {
	nl.dummyHead.next = &node{
		value: e,
		next:  nl.dummyHead.next,
	}
	nl.size++
}

func (nl *NodeList) AddLast(e interface{}) {
	nl.AddIndex(nl.size, e)
}

//func（nl*nodelist）addindex（index int，e interface）//循环实现
//如果index<0_index>nl.size_
//恐慌（“费法伟志”）
//}
//上一个：=&nl.dummyhead
//对于i：=0；i<索引；i++
//prev=上一个下一个
//}
//新节点：=节点
//值：
//下一个：上一个下一个，
//}
//prev.next=新建节点（&N）
//大小+ +
//}
func (nl *NodeList) AddIndex(index int, e interface{}) { //附带条件
	if index < 0 || index > nl.size {
		panic("fei fa wei zhi")
	}
	addIndex(nl.dummyHead, index, e)
	nl.size++
}
func addIndex(prev *node, index int, e interface{}) {
	if index == 0 {
		prev.next = &node{e, prev.next}
		return
	}
	addIndex(prev.next, index-1, e)
}

//func（nl*nodelist）get（index int）interface//循环实现
//如果index<0_index>nl.size_
//恐慌（“费法伟志”）
//}
//当前：=nl.dummyhead.next
//对于i：=0；i<索引；i++
//当前=当前。下一步
//}
//返回当前值
//}
func (nl *NodeList) Get(index int) interface{} { //附带条件
	if index < 0 || index > nl.size {
		panic("fei fa wei zhi")
	}
	return get(nl.dummyHead.next, index)
}
func get(cur *node, index int) interface{} {
	if index == 0 {
		return cur.value
	}
	return get(cur.next, index-1)
}
func (nl *NodeList) GetFirst() interface{} {
	return nl.Get(0)
}
func (nl *NodeList) GetLast() interface{} {
	return nl.Get(nl.size - 1)
}

//func（nl*nodelist）set（index int，e interface）//循环实现
//如果index<0_index>nl.size_
//恐慌（“费法伟志”）
//}
//当前：=nl.dummyhead.next
//对于i：=0；i<索引；i++
//当前=当前。下一步
//}
//当前值=E
//}
func (nl *NodeList) Set(index int, e interface{}) {
	if index < 0 || index > nl.size {
		panic("fei fa wei zhi")
	}
	set(nl.dummyHead.next, index, e)
}
func set(cur *node, index int, e interface{}) {
	if index == 0 {
		cur.value = e
		return
	}
	set(cur.next, index-1, e)
}
func (nl *NodeList) SetFirst(e interface{}) {
	nl.Set(0, e)
}
func (nl *NodeList) SetLast(e interface{}) {
	nl.Set(nl.size-1, e)
}
func (nl *NodeList) Contains(e interface{}) bool {
	current := nl.dummyHead.next
	for current.value != nil {
		if current.value == e {
return true //附带条件
		}
		current = current.next
	}
	return false
}

//func（nl*nodelist）删除（index int）接口
//如果index<0_index>nl.size_
//恐慌（“费法伟志”）
//}
//上一个：=&nl.dummyhead
//对于i：=0；i<索引；i++
//prev=上一个下一个
//}
//返回：=上一个下一个
//prev.next=返回下一个
//Real.NET= nIL
//
//N.L. siZe-
//返回ret.value
//}
func (nl *NodeList) Remove(index int) interface{} {
	if index < 0 || index > nl.size {
		panic("fei fa wei zhi")
	}
	nl.size--
	return removeNode(nl.dummyHead, index)
}
func removeNode(prev *node, index int) interface{} {
	if index == 0 {
		ret := prev.next
		prev.next = ret.next
		ret.next = nil
		return ret.value
	}
	return removeNode(prev.next, index-1)
}
func (nl *NodeList) RemoveFirst() interface{} {
	return nl.Remove(0)
}
func (nl *NodeList) RemoveLast() interface{} {
	return nl.Remove(nl.size - 1)
}
func (nl *NodeList) String() string {
	str := fmt.Sprint("List:")
	str += fmt.Sprint("Head ")
	current := nl.dummyHead.next
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
