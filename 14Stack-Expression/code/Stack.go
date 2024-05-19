package main

import "container/list"

type Stack struct {
	list *list.List //当作栈的结构
}

func NewStack() *Stack {
	list := list.New()  //新建内存
	return &Stack{list} //新建一个栈
}

//入栈
func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

//出栈
func (stack *Stack) Pop() interface{} {
	element := stack.list.Back() //取得最后一个数据
	if element != nil {
		stack.list.Remove(element) ///删除元素
		return element.Value       //返回数据
	}
	return nil
}

//取得数据但是不删除
func (stack *Stack) Peak() interface{} {
	element := stack.list.Back() //取得最后一个数据
	if element != nil {
		//stack.list.Remove(element)///删除元素
		return element.Value //返回数据
	}
	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len() //返回长度
}
func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0 //是否为空
}
