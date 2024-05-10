package main

//import "fmt"

type MyQueue interface {
	Size() int                //大小
	Front() interface{}       //第一个元素
	End() interface{}         //最后一个元素
	IsEmpty() bool            //是否为空
	Enqueue(data interface{}) //入队
	Dequeue() interface{}     //出对
	Clear()                   //清空
}

type Queue struct {
	datastore []*pos
	theSize   int
}

func (myqueue *Queue) Clear() {
	myqueue.datastore = make([]*pos, 0) //开辟内存
	myqueue.theSize = 0
}

func NewQueue() *Queue {
	myqueue := new(Queue)
	myqueue.Clear()
	return myqueue

}
func (myqueue *Queue) Size() int {
	return myqueue.theSize //大小
}
func (myqueue *Queue) Front() *pos {
	if myqueue.Size() == 0 { //判断是否为空
		return nil
	}
	return myqueue.datastore[0]
}
func (myqueue *Queue) End() *pos {
	if myqueue.Size() == 0 { //判断是否为空
		return nil
	}
	return myqueue.datastore[myqueue.theSize-1]
}
func (myqueue *Queue) IsEmpty() bool {
	return myqueue.theSize == 0
}
func (myqueue *Queue) Enqueue(data *pos) {
	myqueue.datastore = append(myqueue.datastore, data) //入队
	//fmt.Println()
	//myqueue.datastore[myqueue.theSize]=data
	myqueue.theSize++
}
func (myqueue *Queue) Dequeue() *pos {
	if myqueue.Size() == 0 { //判断是否为空
		return nil
	}
	data := myqueue.datastore[0]
	if myqueue.Size() > 1 {
		myqueue.datastore = myqueue.datastore[1:myqueue.theSize] //截取
	}
	myqueue.theSize--
	return data
}
