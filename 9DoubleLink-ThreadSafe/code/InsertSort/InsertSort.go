package main

import "fmt"

const Int_MAX = int(^uint(0) >> 1) //位操作

type Node struct {
	value int ///数据
	next  int //下一个索引
}

var NL []Node // 集合

func InitList(arr []int) {
	var node Node
	node = Node{Int_MAX, 1} //第一个节点
	NL = append(NL, node)   //插入第一个节点，第一个节点最大，编号1
	for i := 1; i <= len(arr); i++ {
		node = Node{arr[i-1], 0} //插入数据
		NL = append(NL, node)
	}
	fmt.Println(NL)

}
func ListSort() {
	var i, low, high int
	for i = 2; i < len(NL); i++ {
		low = 0
		high = NL[0].next
		for NL[high].value < NL[i].value { //寻找一个邻居的数据 NL[max] NL[i]，插入NL[min]
			low = high
			high = NL[high].next //前后的两个下标，顺序记录从小到大
		}
		NL[low].next = i
		NL[i].next = high //插入数据到中间
	}
	fmt.Println(NL)
}

//地址排序，插入排序
func Arrange() {
	p := NL[0].next
	for i := 1; i < len(NL); i++ {
		for p < i { //i之前都是排序q:=好的
			p = NL[p].next
		}
		q := NL[p].next //下一个要排序的记录
		if p != i {
			NL[p].value, NL[i].value = NL[i].value, NL[p].value //数据交换
			NL[p].next = NL[i].next                             //修改next
			NL[i].next = p                                      //当前的最小的后一个就是交换的下标                                   //地址插入

		}
		p = q //循环下一个

	}
	for i := 1; i < len(NL); i++ {
		fmt.Println(NL[i].value)
	}
}
func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5, 10}
	InitList(arr) //初始化
	ListSort()    //排序
	Arrange()     //综合处理
}
