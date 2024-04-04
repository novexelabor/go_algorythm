package main

import "fmt"

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func NewLinkStack() *Node {
	return &Node{} //返回一个节点指针
}

func (n *Node) IsEmpty() bool {
	if n.pNext == nil {
		return true
	} else {
		return false
	}
}

func (n *Node) Push(data interface{}) {
	//新建一个节点指针压入栈中
	newnode := &Node{data, nil}
	newnode.pNext = n.pNext //采用头插法，第一个节点功能性
	n.pNext = newnode
}

func (n *Node) Pop() interface{} {
	if n.IsEmpty() {
		return nil
	}
	value := n.pNext.data
	n.pNext = n.pNext.pNext //删除，自动回收
	return value
}

func (n *Node) Length() int {
	pnext := n //值拷贝
	length := 0
	for pnext.pNext != nil {
		length++
		pnext = pnext.pNext
	}
	return length
}

func main1() {
	node1 := new(Node)
	node2 := new(Node)
	node3 := new(Node)
	node4 := new(Node)
	node1.data = 1
	node1.pNext = node2
	node2.data = 2
	node2.pNext = node3
	node3.data = 3
	node3.pNext = node4
	node4.data = 4
	fmt.Println(node1.data)
	fmt.Println(node2.data)
	fmt.Println(node3.data)
	fmt.Println(node4.data)
	fmt.Println("------------------")
	fmt.Println(node1.pNext.pNext.pNext.data)
}

func main2() {
	mystack := NewLinkStack()
	for i := 0; i < 100; i++ {
		mystack.Push(i)
	}
	for data := mystack.Pop(); data != nil; data = mystack.Pop() {
		fmt.Println(data)
	}
}
