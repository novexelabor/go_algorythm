package main

import (
	"fmt"
	//"github.com/mosaicnetworks/babble/src/peers"
)

type CircleLink struct {
	Id   int         //数据编号
	Data interface{} //数据
	Prev *CircleLink //上一个节点，下一个节点
	Next *CircleLink
}

//初始化头节点
func InitHeadNode(data interface{}) *CircleLink {
	return &CircleLink{1, data, nil, nil}
}

//重置头节点,查找到链表为空的时候，进行重置
func (head *CircleLink) ResetHeadNode(data interface{}) {
	if head.Id == 0 {
		head.Id = 1
	}
	head.Data = data
}

//判断头节点是否为空,只有一个数据
func (head *CircleLink) IsHeadEmpty() bool {
	return head.Next == nil && head.Prev == nil
}

//判断链表是否为空
func (head *CircleLink) IsEmpty() bool {
	return head.Data == nil && head.Next == nil && head.Prev == nil
}

//抓取最后元素
func (head *CircleLink) GetLastNode() *CircleLink {
	curnode := head //对于链表的操作，先复制
	if !head.IsHeadEmpty() {
		for {
			if curnode.Next == head { //循环到了最后
				break
			}
			curnode = curnode.Next //循环
		}
	}
	return curnode
}
func (head *CircleLink) Addnode(newnode *CircleLink) {
	if head.IsHeadEmpty() { //只有一个节点，互为前后
		head.Next = newnode
		head.Prev = newnode
		newnode.Prev = head
		newnode.Next = head
		return
	}
	curnode := head //备份第一个数据
	flag := false   //标志，数据添加末尾
	for {
		if curnode == head.Prev {
			break //已经是最后一个节点，退出
		} else if curnode.Next.Id > newnode.Id {
			flag = true //标志下数据应该插入到前列
			break
		} else if curnode.Next.Id == newnode.Id {
			fmt.Printf("数据已经存在\n")
			return
		}
		curnode = curnode.Next //数据循环前进

	}
	if flag {
		//最后一个节点，前面插入,相当于中间插入
		newnode.Next = curnode.Next
		newnode.Prev = curnode

		curnode.Next.Prev = newnode
		curnode.Next = newnode

	} else {
		//z最后一个后面插入
		newnode.Prev = curnode      //300200
		newnode.Next = curnode.Next //300100
		curnode.Next = newnode
		head.Prev = newnode //头节点Prev指向最后一个

	}

}

//双环链表数据查找
func (head *CircleLink) Findnodebyid(id int) (*CircleLink, bool) {
	if head.IsHeadEmpty() && head.Id == id {
		return head, true
	} else if head.IsHeadEmpty() && head.Id != id {
		return &CircleLink{}, false
	}
	curnode := head
	flag := false
	for {
		if curnode.Id == id {
			flag = true //找到
			break
		}
		if curnode == head.Prev { //循环到最后
			break
		}
		curnode = curnode.Next

	}
	if !flag {
		return &CircleLink{}, false
	}
	return curnode, true
}
func (head *CircleLink) Deletenodebyid(id int) bool {
	if head.IsEmpty() {
		fmt.Printf("空链表无法删除\n")
		return false
	}
	node, isok := head.Findnodebyid(id) //搜索判断是否存在
	if isok {
		//删除第一个节点
		if node == head {
			//只有一个节点
			if head.IsHeadEmpty() {
				head.Next = nil
				head.Prev = nil
				head.Data = nil
				head.Id = 0
				return true
			}
			//只有两个节点
			if head.Next.Next == head {
				nextnode := head.Next
				head.Id = nextnode.Id
				head.Data = nextnode.Data
				head.Prev = nil
				head.Next = nil
				return true

			}
			//双环链表，始终保留第一个节点
			//移动下一个节点作为头节点
			nextNodetmp := head.Next
			head.Data = nextNodetmp.Data //把head后一个值复制，相当于删除了head
			head.Id = nextNodetmp.Id

			head.Next = nextNodetmp.Next
			nextNodetmp.Next.Prev = head
			return true

		}
		//删除最后一个节点
		if node == head.GetLastNode() {
			if node.Prev == head && node.Next == head {
				//只有两个元素
				head.Prev = nil
				head.Next = nil
				return true
			}
			head.Prev = node.Prev
			node.Prev.Next = head
			return true
		}
		//处理中间节点
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		return true

	} else {
		fmt.Printf("数据找不到所以无法删除\n")
	}
	return isok

}
func (head *CircleLink) Changenodebyid(id int, data interface{}) bool {
	node, isok := head.Findnodebyid(id)
	if isok {
		node.Data = data //修改数据
	} else {
		fmt.Printf("数据找达不到\n")
	}
	return isok
}

func (head *CircleLink) Showall() {
	if head.IsEmpty() {
		fmt.Printf("空链表无法显示\n")
		return
	}
	if head.IsHeadEmpty() {
		fmt.Println(head.Id, head.Data, head.Prev, head.Next)
		return
	}
	curnode := head
	for {
		fmt.Println(curnode.Id, curnode.Data, curnode.Prev, curnode.Next)
		if curnode == head.Prev {
			break
		}
		curnode = curnode.Next //节点循环继续
	}
}
func main() {
	linknode := InitHeadNode("a")
	linknode.Showall()
	fmt.Println("---------------")
	node1 := &CircleLink{3, "b", nil, nil}
	node2 := &CircleLink{2, "c", nil, nil}
	node3 := &CircleLink{5, "d", nil, nil}
	node4 := &CircleLink{4, "e", nil, nil}
	linknode.Addnode(node1)
	linknode.Showall()
	fmt.Println("---------------")
	linknode.Addnode(node2)
	linknode.Showall()
	fmt.Println("---------------")
	linknode.Addnode(node3)
	linknode.Showall()
	fmt.Println("---------------")
	linknode.Addnode(node4)
	linknode.Showall()
	fmt.Println("---------------")
	linknode.Deletenodebyid(2)
	linknode.Showall()
	linknode.Changenodebyid(3, "x")
	fmt.Println("---------------")
	linknode.Showall()

}
