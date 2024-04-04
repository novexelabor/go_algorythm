package main

import "fmt"

type Node struct { //节点结构
	data  interface{}
	pNext *Node
}

type QueueLink struct {
	front *Node
	rear  *Node
}

type LinkQueue interface {
	Length() int
	Enqueue(data interface{})
	Dequeue() interface{}
}

func NewLinkQueue() *QueueLink {
	return &QueueLink{}
}

func (qlk *QueueLink) Length() int {
	length := 0
	qnext := qlk.front //链式队列front
	for qnext != nil {
		qnext = qnext.pNext
		length++
	}
	return length
}

func (qlk *QueueLink) Enqueue(data interface{}) {
	nlk := &Node{data, nil} // 新节点指针
	if qlk.front == nil {   //判断队列是否为空
		qlk.front = nlk
		qlk.rear = nlk
	} else {
		qlk.rear.pNext = nlk
		qlk.rear = qlk.rear.pNext
	}

}

func (qlk *QueueLink) Dequeue() interface{} {
	if qlk.front == nil {
		return nil
	}
	value := qlk.front.data
	if qlk.front == qlk.rear {
		qlk.front = nil
		qlk.rear = nil
	} else {
		qlk.front = qlk.front.pNext
	}
	return value
}

func main() {
	myq := NewLinkQueue()
	for i := 0; i <= 10; i++ {
		myq.Enqueue(i)
	}
	for data := myq.Dequeue(); data != nil; data = myq.Dequeue() {
		fmt.Println(data)
	}
}
