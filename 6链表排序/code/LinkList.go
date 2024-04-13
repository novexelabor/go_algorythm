package main

import (
	"fmt"
	//"github.com/nntaoli-project/GoEx/aacoin"
)

const (
	ERROR = -21231231231231
)

//定义类型
type Element int64

//链表基本结构
type LinkNode struct {
	Data  Element
	pNext *LinkNode
}

type LinkNoder interface {
	Add(head *LinkNode, data Element, ishead bool)
	AddHead(head *LinkNode, data Element)
	Append(head *LinkNode, data Element)
	Show(head *LinkNode)
	Getlength(head *LinkNode) int
	Search(head *LinkNode, data Element) bool
	Getdata(head *LinkNode, id int) Element
	Delete(head *LinkNode, id int) Element
	clear(phead *LinkNode)
	Insert(phead *LinkNode, index int, data Element)

	Reverse(head *LinkNode)
	BubbleSort(head *LinkNode)
	SelectSort(head *LinkNode)

	InsertSort(head *LinkNode)
	MergeSort(head *LinkNode) //合并
	QucikSort(head *LinkNode)
}

//创建头部节点，不放任何数据
func NewLinkList() *LinkNode {
	return &LinkNode{Data: 0, pNext: nil}
}
func Add(head *LinkNode, data Element, ishead bool) {
	if ishead {
		AddHead(head, data)
	} else {
		Append(head, data)
	}
}
func AddHead(head *LinkNode, data Element) {
	node := &LinkNode{Data: data} //新建一个节点
	node.pNext = head.pNext
	head.pNext = node
}
func Append(head *LinkNode, data Element) {
	node := &LinkNode{Data: data} //新建一个节点
	if head.pNext == nil {
		head.pNext = node
	} else {
		pcur := head //拷贝一份
		for pcur.pNext != nil {
			pcur = pcur.pNext //循环到最后
		}
		pcur.pNext = node //尾部插入
	}
}
func show(head *LinkNode) {
	pahead := head.pNext //从头节点下一个节点开始
	for pahead != nil {
		fmt.Println(pahead.Data)
		pahead = pahead.pNext //循环
	}
	fmt.Println("show over")
}

func Getlength(head *LinkNode) int {
	pahead := head.pNext //从头节点下一个节点开始
	var length int = 0
	for pahead != nil {
		//fmt.Println(pahead.Data)
		length++
		pahead = pahead.pNext //循环
	}
	//fmt.Println("show over")
	return length
}
func Search(head *LinkNode, data Element) bool {
	pahead := head.pNext //从头节点下一个节点开始
	var isfind bool = false
	for pahead != nil {
		if pahead.Data == data {
			isfind = true
			break
		}
		pahead = pahead.pNext //循环
	}
	//fmt.Println("show over")
	return isfind
}
func Getdata(head *LinkNode, id int) Element {
	if id <= 0 || id > Getlength(head) {
		return ERROR
	} else {
		pgo := head
		for i := 0; i < id; i++ {
			pgo = pgo.pNext //0对应的是第一个
		}
		return pgo.Data
	}
}
func Delete(head *LinkNode, id int) Element { //个数从1开始的
	if id <= 0 || id > Getlength(head) {
		return ERROR
	} else {
		pgo := head
		for i := 0; i < id-1; i++ { //该位置，找到的是id了，找id的前一个,所以是id-1
			pgo = pgo.pNext
		}
		//要删除的位置
		data := pgo.pNext.Data

		pgo.pNext = pgo.pNext.pNext //删除

		return data

	}
}
func clear(phead *LinkNode) {
	phead = nil
}
func Insert(head *LinkNode, id int, data Element) {
	if id < 0 || id > Getlength(head) {
		return
	} else {
		pgo := head
		for i := 0; i < id-1; i++ { //i=0是第一个,找到前一个
			pgo = pgo.pNext
		}
		var node LinkNode //新节点
		node.Data = data  //插入
		node.pNext = pgo.pNext
		pgo.pNext = &node

	}
}
func reverseNode(head *LinkNode) *LinkNode {
	var pnow *LinkNode = head.pNext //第一个节点
	var pre *LinkNode = nil
	var pnext *LinkNode = nil
	var phead *LinkNode = head
	//head=head.pNext
	for pnow != nil {
		pnext = pnow.pNext //保存当前节点下一个节点
		if pnext == nil {
			phead.pNext = pnow //记录最后一个节点
		}
		pnow.pNext = pre
		pre = pnow
		pnow = pnext
	}

	return phead

}

func reverseNodeList(head *LinkNode, index int) *LinkNode {
	j := index //记录第几个
	//如果链表为空或者链表中只有一个元素
	if head == nil || head.pNext == nil {
		return head
	} else {
		//先反转后面的链表，走到链表的末端结点
		var newhead *LinkNode = reverseNodeList(head.pNext, j+1)
		if j == 0 {
			head.pNext = newhead
			return head
		} else {
			head.pNext.pNext = head //实现反转,不用循环遍历了
			head.pNext = nil
			return newhead

		}
	}
}
func reverseNodeList1(head *LinkNode, index int) *LinkNode {
	j := index //记录第几个
	//如果链表为空或者链表中只有一个元素
	if head == nil || head.pNext == nil {
		return head
	} else {
		//先反转后面的链表，走到链表的末端结点
		var newhead *LinkNode = reverseNodeList1(head.pNext, j+1)
		if j == 0 {
			head.pNext = newhead
			return head
		} else {
			tail := newhead
			for tail.pNext != nil {
				tail = tail.pNext
			}
			tail.pNext = head
			head.pNext = nil
			return newhead

		}
	}
}

//冒泡排序
func BubbleSort(head *LinkNode) {
	for phead1 := head.pNext; phead1.pNext != nil; phead1 = phead1.pNext {
		for phead2 := head.pNext; phead2.pNext != nil; phead2 = phead2.pNext {
			if phead2.Data > phead2.pNext.Data {
				phead2.pNext.Data, phead2.Data = phead2.Data, phead2.pNext.Data
			}
		}
	}
}

//选择
//13579468
//9  1357    468
func selectsort(head *LinkNode) *LinkNode {

	if head == nil || head.pNext == nil {
		return head
	} else {
		fmt.Println("select start ")
		//假定头节点为空，
		for newhead := head; newhead != nil; newhead = newhead.pNext {

			//fmt.Println(newhead.Data)
			pahead := newhead //从头节点下一个节点开始
			maxnode := newhead
			for pahead != nil {
				if pahead.Data < maxnode.Data {
					maxnode = pahead
				}
				pahead = pahead.pNext //循环
			}
			maxnode.Data, newhead.Data = newhead.Data, maxnode.Data

		}
		/*
			pahead:=head.pNext  //从头节点下一个节点开始
			maxnode:=head.pNext
			for pahead!=nil{
				if pahead.Data>maxnode.Data{
					maxnode=pahead
				}
				pahead=pahead.pNext//循环
			}
			maxnode.Data,head.pNext.Data=head.pNext.Data,maxnode.Data
		*/
		return head
	}

}

// 3   1479
//13   479
//134  79
func InsertSort(head *LinkNode) *LinkNode {

	if head == nil || head.pNext == nil {
		return head
	} else {
		var pstart *LinkNode = new(LinkNode)
		var pend *LinkNode = head    //排好序的最后一个节点
		var p *LinkNode = head.pNext //第一个节点
		pstart = head                //存储头节点

		for p != nil { //待排序链表循环遍历的节点
			var ptmp *LinkNode = pstart.pNext      //第一个节点
			var pre *LinkNode = pstart             //记录前一个位置，方便删除与插入
			for ptmp != p && p.Data >= ptmp.Data { //找到要插入的位置
				ptmp = ptmp.pNext //ptmp 是第二个循环遍历的节点
				pre = pre.pNext
			}
			if ptmp == p { //插入,p之前是已经排列好的
				pend = p //排序好的最后一个节点
			} else {
				pend.pNext = p.pNext
				p.pNext = ptmp //在ptmp之前插入节点
				pre.pNext = p  //其实没有意义了，是保持了pre指向前一个节点
			}
			p = pend.pNext
			//show(head)
		}

		//head = pstart.pNext

		return pstart
	}

}
func MergeSort(head *LinkNode) *LinkNode {
	if head == nil || head.pNext == nil {
		return head
	} else {
		var phead *LinkNode = head
		var qhead *LinkNode = head
		var pre *LinkNode = nil
		//归并位置，
		//7  1  2
		//8
		// 1,3,      5,   7,9
		for qhead != nil && qhead.pNext != nil {
			qhead = qhead.pNext.pNext //走两步
			pre = phead               //记录中间位置的上一步
			phead = phead.pNext       //中间
		}
		pre.pNext = nil //拆分两端

		//这样才能1，2；1，2的合并的
		var left *LinkNode = MergeSort(head)   //继续拆分左边
		var right *LinkNode = MergeSort(phead) //继续拆分右边
		return Merge(left, right)
	}
}

func Merge(left *LinkNode, right *LinkNode) *LinkNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	var pres *LinkNode = new(LinkNode)
	var ptmp *LinkNode = pres //开辟节点，备份,pres头指针

	for left != nil && right != nil {
		if left.Data < right.Data {
			ptmp.pNext = left
			ptmp = ptmp.pNext //更新下一个节点
			left = left.pNext
		} else {
			ptmp.pNext = right
			ptmp = ptmp.pNext
			right = right.pNext
		}
	}
	if left != nil {
		ptmp.pNext = left
	} else {
		ptmp.pNext = right
	}
	ptmp = pres.pNext //回到最开始的位置

	return ptmp //没有头指针的已经排好序的链表

}

func QuickSort(head *LinkNode) *LinkNode {
	if head == nil || head.pNext == nil {
		return head
	} else {
		var tmphead LinkNode = LinkNode{0, nil}
		qsortlist(&tmphead, head, nil)
		return head
	}

}
func qsortlist(headpre, head, tail *LinkNode) {
	if head != tail && head.pNext != tail {
		var mid *LinkNode = partition(headpre, head, tail) //递归
		qsortlist(headpre, headpre.pNext, mid)
		qsortlist(mid, mid.pNext, tail)
	}
}

// 3       8 9 2 1
//3       8 29  1
//3       1 2      89

func partition(lowpre, low, high *LinkNode) *LinkNode {
	key := low.Data //取得中间数据
	var node1 LinkNode = LinkNode{0, nil}
	var node2 LinkNode = LinkNode{0, nil}
	var small, big *LinkNode = &node1, &node2

	//遍历不到high节点
	for i := low.pNext; i != high; i = i.pNext {
		if i.Data < key {
			small.pNext = i
			small = i
		} else {
			big.pNext = i
			big = i
		}
	}
	big.pNext = high           //第一次的high节点是nil
	small.pNext = low          //第一次的low是head节点
	low.pNext = node2.pNext    //存放较大值,把较低的连接上较大者
	lowpre.pNext = node1.pNext //存放较小者
	return low
}

func reversrList(head *LinkNode) *LinkNode {
	cur := head
	var pre *LinkNode = nil
	for cur != nil {
		pre, cur, cur.pNext = cur, cur.pNext, pre //这句话最重要
	}
	return pre
}

//-------------------------非常完美解决该问题
func QuickSort1(head *LinkNode) {
	if head == nil || head.pNext == nil {
		return
	} else {
		tail := head
		for tail.pNext != nil {
			tail = tail.pNext
		}
		qsortlist1(head, tail)
	}

}

func qsortlist1(head, tail *LinkNode) {
	if head == tail {
		return
	} else {
		midpre := partition1(head, tail)
		if midpre == nil {
			if head.pNext != tail {
				qsortlist1(head.pNext, tail)
			}
		} else {
			if head != midpre {
				qsortlist1(head, midpre)
			}
			if midpre.pNext != tail && midpre.pNext.pNext != tail && midpre.pNext.pNext != nil {
				qsortlist1(midpre.pNext.pNext, tail)
			}
		}
	}

}

func partition1(low, high *LinkNode) *LinkNode {

	key := low.Data
	tem := low //记录当前key的节点
	var pre *LinkNode = nil

	node1 := low //记录小于的一部分
	for i := low; i != high.pNext; i = i.pNext {
		if i.Data < key {
			node1.Data, i.Data = i.Data, node1.Data
			pre = node1         //记录前一个节点
			node1 = node1.pNext //下一次要交换的节点
			tem = i
		}
	}
	if tem != node1 { //下次要交换的位置不相同，将key交换到分割节点处
		tem.Data, node1.Data = node1.Data, tem.Data
	}
	return pre
}

// //--------------------------------------

func main() {
	var phead *LinkNode = NewLinkList()
	Append(phead, 9)
	Append(phead, 3)
	Append(phead, 2)
	Append(phead, -19)
	Append(phead, 10)
	Append(phead, 25)
	//Append(phead, -98)

	show(phead)

	// Delete(phead, 2) //delete修改正确
	// show(phead)

	// Insert(phead, 2, 2) //insert修改正确
	// show(phead)

	//selectsort(phead)
	QuickSort1(phead.pNext) //完美解决该问题

	//phead=reverseNode(phead)
	//phead=reverseNode(phead)

	// fmt.Println("----------------")
	// phead = reverseNode(phead) //修改完成
	// show(phead)

	//传入第一个节点，不包括头指针即可
	//phead.pNext = MergeSort(phead.pNext)
	//phead = InsertSort(phead)  //可以成功

	// fmt.Println("reverseNodeList1") //修改完成，完美
	// phead = reverseNodeList1(phead, 0)
	show(phead)

	// fmt.Println("reverseNodeList") //修改完成，完美
	// phead = reverseNodeList1(phead, 0)
	// show(phead)

	//fmt.Println(phead.Data,phead.pNext.Data,phead.pNext.pNext.Data)
	//fmt.Println("----------------")
	//fmt.Println(phead.Data)

}
