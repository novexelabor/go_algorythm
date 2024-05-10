package main

import "github.com/pkg/errors"

type pos struct {
	x int
	y int
}

type Node struct {
	data *pos
	next *Node
}

type LinkStack interface { //采用的是链栈
	IsEmpty() bool
	Push(value string)
	Pop() (string, error)
	Length() int
}

func NewStack() *Node {
	return &Node{}
}

func (n *Node) IsEmpty() bool { //判断是否为空,带有头节点的
	return n.next == nil
}
func (n *Node) Push(value *pos) {
	newnode := &Node{data: value} //初始化
	newnode.next = n.next
	n.next = newnode
}
func (n *Node) Pop() (*pos, error) {
	if n.IsEmpty() == true {
		return nil, errors.New("bug")
	}
	value := n.next.data
	n.next = n.next.next
	return value, nil
}
func (n *Node) Length() int {
	pnext := n
	length := 0
	for pnext.next != nil { //返回长度
		pnext = pnext.next
		length++
	}
	return length
}
