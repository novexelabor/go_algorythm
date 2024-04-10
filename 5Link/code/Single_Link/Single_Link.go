package Single_Link

import (
	"fmt"
	"strings"
)

//单链表的接口，
type SingleLink interface {
	//增删查改
	GetFirstNode() *SingleLinkNode        //抓取头部节点
	InsertNodeFront(node *SingleLinkNode) //头部插入
	InsertNodeBack(node *SingleLinkNode)  //尾部插入

	//在一个节点之前插入或者一个节点之后插入
	InsertNodeValueBack(dest interface{}, node *SingleLinkNode) bool
	InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool

	GetNodeAtIndex(index int) *SingleLinkNode //根据索引抓取指定位置的节点

	DeleteNode(dest *SingleLinkNode) bool //删除一个节点
	Deleteatindex(index int)              //删除指定位置的节点

	String() string //返回链表字符串

}

//链表的结构
type SingleLinkList struct {
	head   *SingleLinkNode //链表的头指针
	length int             //链表的长度
}

//创建链表
func NewSingleLinkList() *SingleLinkList {
	head := NewSingleLinkNode(nil) //头部指针
	return &SingleLinkList{head, 0}
}

//返回第一个数据节点
func (list *SingleLinkList) GetFirstNode() *SingleLinkNode {
	return list.head.pNext
}

//头部插入
func (list *SingleLinkList) InsertNodeFront(node *SingleLinkNode) {
	if list.head == nil { //头指针是空，直接插入head后面
		list.head.pNext = node
		node.pNext = nil
		list.length++ //插入节点，数据追加
	} else {
		bak := list.head       //4000500
		node.pNext = bak.pNext //400500
		bak.pNext = node
		list.length++ //插入节点，数据追加
	}
}

//尾部插入
func (list *SingleLinkList) InsertNodeBack(node *SingleLinkNode) {
	if list.head == nil { //链表为空时，直接插入到head后面
		list.head.pNext = node
		node.pNext = nil
		list.length++ //插入节点，数据追加
	} else {
		bak := list.head //4000500
		for bak.pNext != nil {
			bak = bak.pNext //循环到最后
		}
		bak.pNext = node
		list.length++ //插入节点，数据追加
	}
}
func (list *SingleLinkList) String() string {
	var listString string
	p := list.head //头部节点
	for p.pNext != nil {
		listString += fmt.Sprintf("%v-->", p.pNext.value)
		p = p.pNext //循环
	}
	listString += fmt.Sprintf("nil")
	return listString //打印链表字符串
}

func (list *SingleLinkList) InsertNodeValueBack(dest interface{}, node *SingleLinkNode) bool {
	phead := list.head
	isfind := false //是否找到数据
	for phead.pNext != nil {
		if phead.value == dest { //找到
			isfind = true
			break
		}
		phead = phead.pNext
	}
	if isfind {
		//尾部插入
		node.pNext = phead.pNext
		phead.pNext = node
		list.length++
		return true
	} else {

		return false
	}

}
func (list *SingleLinkList) InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool {

	phead := list.head
	isfind := false //是否找到数据
	for phead.pNext != nil {
		if phead.pNext.value == dest { //找到
			isfind = true
			break
		}
		phead = phead.pNext
	}
	if isfind {
		node.pNext = phead.pNext
		phead.pNext = node
		list.length++

		return isfind
	} else {
		return isfind
	}

}

func (list *SingleLinkList) FindString(data string) {

	phead := list.head.pNext //指定头部
	for phead != nil {       //循环所有数据
		if strings.Contains(phead.value.(string), data) { //包含
			fmt.Println(phead.value)
		}
		phead = phead.pNext
	}

}

func (list *SingleLinkList) GetNodeAtIndex(index int) *SingleLinkNode {
	if index > list.length-1 || index < 0 {
		return nil
	} else {
		phead := list.head
		for index > -1 {
			phead = phead.pNext //向后循环
			index--             //向后循环过程
		}
		return phead

	}
}

func (list *SingleLinkList) DeleteNode(dest *SingleLinkNode) bool {
	if dest == nil {
		return false
	}
	phead := list.head
	for phead.pNext != nil && phead.pNext != dest {
		phead = phead.pNext //循环下去
	}
	if phead.pNext == dest {
		phead.pNext = phead.pNext.pNext
		list.length--
		return true
	} else {
		return false
	}

}
func (list *SingleLinkList) Deleteatindex(index int) {
	if index > list.length-1 || index < 0 {
		return
	} else {
		phead := list.head
		for index > 0 { //找到index的前一个节点
			phead = phead.pNext //向后循环
			index--             //向后循环过程
		}

		phead.pNext = phead.pNext.pNext
		list.length--
		return

	}
}

//1>2>3>4>5
//2>3>4>5
//--   --  /
//- -
//求链表分数位置

func (list *SingleLinkList) GetMid() *SingleLinkNode {
	if list.head.pNext == nil {
		return nil
	} else {
		phead1 := list.head
		phead2 := list.head
		for phead2 != nil && phead2.pNext != nil {
			phead1 = phead1.pNext
			phead2 = phead2.pNext.pNext
		}
		return phead1 //中间节点
	}

}

//链表反转
func (list *SingleLinkList) ReverseList() {
	if list.head == nil || list.head.pNext == nil {
		return //链表为空或者链表只有一个节点
	} else {
		var pre *SingleLinkNode                   //前面节点
		var cur *SingleLinkNode = list.head.pNext //当前节点
		for cur != nil {
			curNext := cur.pNext // 后续节点
			cur.pNext = pre      //反转第一步

			pre = cur     //持续推进
			cur = curNext //持续推进
		}
		fmt.Println(pre)
		list.head.pNext.pNext = nil //让第一个pNext=nil
		list.head.pNext = pre

	}
}
