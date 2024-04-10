package Double_Link

import (
	"fmt"
	"strings"
)

//双链表的基本结构
type DoubleLinkList struct {
	head   *DoubleLinkNode
	length int
}

//新建一个双链表
func NewDoubleLinkList() *DoubleLinkList {
	head := NewDoubleLinkNode(nil)
	return &DoubleLinkList{head, 0}

}

//返回链表长度
func (dlist *DoubleLinkList) Getlength() int {
	return dlist.length
}

//返回第一个节点
func (dlist *DoubleLinkList) GetFirstNode() *DoubleLinkNode {
	return dlist.head.next
}

func (dlist *DoubleLinkList) InsertHead(node *DoubleLinkNode) {
	phead := dlist.head //头节点
	if phead.next == nil {
		node.next = nil   //nil
		phead.next = node //只有一个节点直接链接上
		node.prev = phead //上一个节点

		dlist.length++
	} else {
		phead.next.prev = node //标记上一个节点
		node.next = phead.next //下一个节点

		phead.next = node //标记头部节点的下一个
		node.prev = phead

		dlist.length++
	}

}

func (dlist *DoubleLinkList) InsertBack(node *DoubleLinkNode) {
	phead := dlist.head //头节点
	if phead.next == nil {
		node.next = nil   //nil
		phead.next = node //只有一个节点直接链接上
		node.prev = phead //上一个节点
		dlist.length++

	} else {
		for phead.next != nil {
			phead = phead.next //循环下去
		}
		phead.next = node //后缀
		node.prev = phead //前缀

		dlist.length++
	}
}

func (dlist *DoubleLinkList) String() string {
	var listString1 string
	var listString2 string
	phead := dlist.head
	for phead.next != nil {
		//正向循环
		listString1 += fmt.Sprintf("%v-->", phead.next.value)
		phead = phead.next
	}
	listString1 += fmt.Sprintf("nil")
	listString1 += "\n" //--> 最后的一个节点了,从最后的节点开始往前prev
	for phead != dlist.head {
		//反向循环
		listString2 += fmt.Sprintf("<--%v", phead.value)
		phead = phead.prev
	}
	listString1 += fmt.Sprintf("nil")

	return listString1 + listString2 + "\n" //打印链表字符串
}

func (dlist *DoubleLinkList) InsertValueBack(dest *DoubleLinkNode, node *DoubleLinkNode) bool {
	phead := dlist.head
	for phead.next != nil && phead.next != dest { //循环查找
		phead = phead.next
	}
	if phead.next == dest {
		if phead.next.next != nil {
			phead.next.next.prev = node //500100
		}
		node.next = phead.next.next //300300
		phead.next.next = node      //500100
		node.prev = phead.next      //300200

		//dlist.head=phead
		dlist.length++
		return true
	} else {

		return false
	}

}
func (dlist *DoubleLinkList) InsertValueHead(dest *DoubleLinkNode, node *DoubleLinkNode) bool {

	phead := dlist.head
	for phead.next != nil && phead.next != dest { //循环查找
		phead = phead.next
	}
	if phead.next == dest {
		if phead.next != nil { //不是最后节点
			phead.next.prev = node //400500
		}
		node.next = phead.next //300100
		node.prev = phead      //300200
		phead.next = node      //400500

		//dlist.head=phead
		dlist.length++
		return true
	} else {

		return false
	}
}
func (dlist *DoubleLinkList) InsertValueBackByValue(dest interface{}, node *DoubleLinkNode) bool {
	phead := dlist.head
	for phead.next != nil && phead.next.value != dest { //循环查找
		phead = phead.next
	}
	if phead.next.value == dest {
		if phead.next.next != nil {
			phead.next.next.prev = node //500100
		}
		node.next = phead.next.next //300300
		phead.next.next = node      //500100
		node.prev = phead.next      //300200

		//dlist.head=phead
		dlist.length++
		return true
	} else {

		return false
	}

}
func (dlist *DoubleLinkList) InsertValueHeadByValue(dest interface{}, node *DoubleLinkNode) bool {

	phead := dlist.head
	for phead.next != nil && phead.next.value != dest { //循环查找
		phead = phead.next
	}
	if phead.next.value == dest {
		if phead.next != nil {
			phead.next.prev = node //400500
		}
		node.next = phead.next //300100
		node.prev = phead      //300200
		phead.next = node      //400500

		//dlist.head=phead
		dlist.length++
		return true
	} else {

		return false
	}
}

func (dlist *DoubleLinkList) GetNodeAtindex(index int) *DoubleLinkNode {
	if index > dlist.length-1 || index < 0 {
		return nil
	}
	phead := dlist.head
	for index > -1 {
		phead = phead.next
		index-- //计算位置
	}
	return phead
}

func (dlist *DoubleLinkList) DeleteNode(node *DoubleLinkNode) bool {
	if node == nil {
		return false
	} else {
		phead := dlist.head
		for phead.next != nil && phead.next != node {
			phead = phead.next //循环查找
		}
		if phead.next == node {
			if phead.next.next != nil {
				phead.next.next.prev = phead //设置pre
			}
			phead.next = phead.next.next //设置next

			dlist.length--
			return true
		} else {
			return false
		}

		return true

	}

}

func (dlist *DoubleLinkList) DeleteNodeAtindex(index int) bool {
	if index > dlist.length-1 || index < 0 {
		return false
	}
	phead := dlist.head
	for index > 0 { //找到前一个节点
		phead = phead.next
		index-- //计算位置
	}

	if phead.next.next != nil {
		phead.next.next.prev = phead //设置pre
	}
	phead.next = phead.next.next //设置next

	dlist.length--
	return true

}
func (dlist *DoubleLinkList) FindString(inputstr string) {
	phead := dlist.head.next
	for phead != nil {
		//正向循环
		if strings.Contains(phead.value.(string), inputstr) {
			fmt.Println(phead.value.(string))
		}

		phead = phead.next
	}
}
