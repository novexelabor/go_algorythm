package main

import "fmt"

//单链表的节点
type node struct {
	num  int
	next *node
}

//环链表 ,标记头部，结尾
var head, tail *node

func Addnode(n *node) { //向单环链表中添加节点
	if tail == nil { //没有节点的情况下
		head = n
		n.next = head
		tail = n
	} else {
		tail.next = n
		n.next = head
		tail = n
	}
}
func showlist(head *node) {
	if head == nil {
		return
	} else {
		//循环环链表
		for head.next != nil && head != tail {
			fmt.Println(head.num)
			head = head.next
		}
		fmt.Println(head.num) //只要非空，便打印节点数值
	}
}

//从第K个，循环起第num个，留下最后一个
func jose(k, num int) { //需要把head和tail当成参数传递进来
	count := 1                  //记录次数
	for i := 0; i <= k-1; i++ { //到第K个位置处
		head = head.next //表示有头节点的
		tail = tail.next //循环到起点，head 和 tail同时往后移动
	}
	for {
		count++ //开始记录次数
		head = head.next
		tail = tail.next //循环到起点

		if count == num {
			fmt.Println(head.num, "出局")
			tail.next = head.next //删除一个,为了形成环链，tail的next要指向head，修改tail的指向
			head = head.next      //删除掉当前的head节点

			count = 1 //清零
		}
		if head == tail { //相等意味着仅剩一个
			fmt.Println(head.num, "最后一个")
			break
		}

	}

}

func main() {
	for i := 0; i < 10; i++ {
		n := &node{i, nil}
		Addnode(n)
	}
	showlist(head)
	jose(3, 3)
	//012 5  8
	//showlist(head)
}
