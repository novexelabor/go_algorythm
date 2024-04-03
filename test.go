package main

import (
	"fmt"
	"go_algorythm/ArrayList"
	"go_algorythm/CircleQueue"
	"go_algorythm/Queue"
	"go_algorythm/StackArray"
)

func main1() {
	//定义接口对象，赋值的对象必须实现接口的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	for i := 0; i < 55; i++ {
		list.Insert(1, "x5")
		fmt.Println(list)
	}
}

func main2() {
	//定义接口对象，赋值的对象必须实现接口的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	list.Append("d3")
	list.Append("f3")
	fmt.Println(list.Size()) //共有多少个值
	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next("111111")
		if item == "d3" {
			it.Remove()
		}
		fmt.Println(item)
	}
	fmt.Println(list)
	fmt.Println(list.Size())
}

func main3() {
	mystack := ArrayList.NewArrayListStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}

func main4() {
	mystack := ArrayList.NewArrayListStackX()
	mystack.Push(1)
	fmt.Println(mystack.Pop())
	mystack.Push(2)
	fmt.Println(mystack.Pop())
	mystack.Push(3)
	fmt.Println(mystack.Pop())
	mystack.Push(4)
	fmt.Println(mystack.Pop())
	mystack.Push(11)
	mystack.Push(22)
	mystack.Push(33)
	mystack.Push(44)

	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Pop())

	for it := mystack.Myit; it.HasNext(); { //通过迭代遍历数组
		item, _ := it.Next("111111")
		fmt.Println(item)
	}
}

func main5() {
	mystack := StackArray.NewStack()
	mystack.Push(1)
	mystack.Push(2)
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}

func main6() {
	myq := Queue.NewQueue()
	myq.EnQueue(1)
	myq.EnQueue(2)
	myq.EnQueue(3)
	myq.EnQueue(4)
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	myq.EnQueue(14)
	myq.EnQueue(114)
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	myq.EnQueue(11114)
	fmt.Println(myq.DeQueue())
}

func main7() {
	var myq CircleQueue.CircleQueue
	CircleQueue.InitQueue(&myq)
	CircleQueue.EnQueue(&myq, 1)
	CircleQueue.EnQueue(&myq, 2)
	CircleQueue.EnQueue(&myq, 3)
	CircleQueue.EnQueue(&myq, 4)
	CircleQueue.EnQueue(&myq, 5)
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
	fmt.Println(CircleQueue.DeQueue(&myq))
}

func FAB(num int) int { //斐波那契递归计算
	if num == 1 || num == 2 {
		return 1
	} else {
		return FAB(num-1) + FAB(num-2)
	}
}

//使用栈实现斐波那契非递归计算
func main() {
	mystack := StackArray.NewStack()
	mystack.Push(7)
	last := 0
	for !mystack.IsEmpty() {
		data := mystack.Pop()
		if data == 1 || data == 2 {
			last += 1
		} else {
			mystack.Push(data.(int) - 1)
			mystack.Push(data.(int) - 2)
		}
	}
	fmt.Println(last)
}
