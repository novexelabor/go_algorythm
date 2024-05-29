package main

import (
	"fmt"

	"github.com/pkg/errors"
)

//AVL树适用于没有删除的情况
//红黑树的增删查改最优先，
type AVLnode struct {
	Data   interface{} //数据
	Left   *AVLnode    //指向左边
	Right  *AVLnode    //指向右边
	height int         //高度

}

//comparator 函数指针类型
type comparator func(a, b interface{}) int

//compare函数指针
var compare comparator

//比大小
func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func NewNode(data interface{}) *AVLnode {
	node := new(AVLnode) //新建节点,初始化
	node.Data = data
	node.Left = nil
	node.Right = nil
	node.height = 1
	return node
}

//新建AVLｔｒｅｅ
func NewAVLTree(data interface{}, myfunc comparator) (*AVLnode, error) {
	if data == nil && myfunc == nil {
		return nil, errors.New("不能为空")
	}
	compare = myfunc
	return NewNode(data), nil
}
func (avlnode *AVLnode) Getall() []interface{} {
	values := make([]interface{}, 0)
	return AddValues(values, avlnode)
}

func AddValues(values []interface{}, avlnode *AVLnode) []interface{} {
	if avlnode != nil {
		values = AddValues(values, avlnode.Left)
		values = append(values, avlnode.Data)
		fmt.Println(avlnode.Data, avlnode.height)
		values = AddValues(values, avlnode.Right)
	}
	return values
}

//左旋，逆时针
func (avlnode *AVLnode) LeftRotate() *AVLnode {
	headnode := avlnode.Right
	avlnode.Right = headnode.Left
	headnode.Left = avlnode

	//更新高度
	avlnode.height = Max(avlnode.Left.GetHeight(), avlnode.Right.GetHeight()) + 1
	headnode.height = Max(headnode.Left.GetHeight(), headnode.Right.GetHeight()) + 1
	return headnode

}

//右旋，顺时针
func (avlnode *AVLnode) RightRotate() *AVLnode {
	headnode := avlnode.Left //左边节点
	avlnode.Left = headnode.Right
	headnode.Right = avlnode

	//更新高度
	avlnode.height = Max(avlnode.Left.GetHeight(), avlnode.Right.GetHeight()) + 1
	headnode.height = Max(headnode.Left.GetHeight(), headnode.Right.GetHeight()) + 1
	return headnode

}

//两次左旋
//两次右旋

//先左旋再右旋
func (avlnode *AVLnode) LeftThenRightRotate() *AVLnode {
	sonheadnode := avlnode.Left.LeftRotate()
	avlnode.Left = sonheadnode
	return avlnode.RightRotate()
}

//先右旋再左旋
func (avlnode *AVLnode) RightThenLeftRotate() *AVLnode {
	sonheadnode := avlnode.Right.RightRotate()
	avlnode.Right = sonheadnode
	return avlnode.LeftRotate()
}

//自动处理不平衡，差距为1平衡，差距2不平衡
func (avlnode *AVLnode) adjust() *AVLnode {
	if avlnode.Right.GetHeight()-avlnode.Left.GetHeight() == 2 {
		if avlnode.Right.Right.GetHeight() > avlnode.Right.Left.GetHeight() {
			avlnode = avlnode.LeftRotate()
		} else {
			avlnode = avlnode.RightThenLeftRotate()
		}

	} else if avlnode.Left.GetHeight()-avlnode.Right.GetHeight() == 2 {
		if avlnode.Left.Left.GetHeight() > avlnode.Left.Right.GetHeight() {
			avlnode = avlnode.RightRotate()
		} else {
			avlnode = avlnode.LeftThenRightRotate()
		}
	}
	return avlnode
}

//数据插入
func (avlnode *AVLnode) Insert(value interface{}) *AVLnode {
	if avlnode == nil {
		newNode := &AVLnode{value, nil, nil, 1} //插入节点
		return newNode
	}
	switch compare(value, avlnode.Data) {
	case -1:
		avlnode.Left = avlnode.Left.Insert(value)
		avlnode = avlnode.adjust() //自动平衡
	case 1:
		avlnode.Right = avlnode.Right.Insert(value)
		avlnode = avlnode.adjust() //自动平衡
	case 0:
		fmt.Println("数据已经存在")
	}
	//当avlnode节点没有调整时，avlnode高度加1
	avlnode.height = Max(avlnode.Left.GetHeight(), avlnode.Right.GetHeight()) + 1
	return avlnode
}

//删除
func (avlnode *AVLnode) Delete(value interface{}) *AVLnode {
	if avlnode == nil {
		return nil
	}
	switch compare(value, avlnode.Data) {
	case -1:
		avlnode.Left = avlnode.Left.Delete(value)
	case 1:
		avlnode.Right = avlnode.Right.Delete(value)
	case 0:
		//删除在这里

		if avlnode.Left != nil && avlnode.Right != nil { //左右都有节点
			avlnode.Data = avlnode.Right.FindMin().Data
			avlnode.Right = avlnode.Right.Delete(avlnode.Data)
		} else if avlnode.Left != nil { //左孩子存在，右孩子存在或者不存在
			avlnode = avlnode.Left
		} else { //只有一个右孩子，或者无孩子
			avlnode = avlnode.Right
		}

	}
	if avlnode != nil {
		avlnode.height = Max(avlnode.Left.GetHeight(), avlnode.Right.GetHeight()) + 1
		avlnode = avlnode.adjust() //自动平衡
	}

	return avlnode
}

func (avlnode *AVLnode) Find(data interface{}) *AVLnode {
	var finded *AVLnode = nil
	switch compare(data, avlnode.Data) {
	case -1:
		finded = avlnode.Left.Find(data)
	case 1:
		finded = avlnode.Right.Find(data)
	case 0:
		return avlnode
	}
	return finded

}
func (avlnode *AVLnode) FindMin() *AVLnode {
	var finded *AVLnode
	if avlnode.Left != nil {
		finded = avlnode.Left.FindMin() //递归调用
	} else {
		finded = avlnode
	}
	return finded
}
func (avlnode *AVLnode) FindMax() *AVLnode {
	var finded *AVLnode
	if avlnode.Right != nil {
		finded = avlnode.Right.FindMax() //递归调用
	} else {
		finded = avlnode
	}
	return finded
}

//抓取数据
func (avlnode *AVLnode) Getdata() interface{} {
	return avlnode.Data
}

//设置
func (avlnode *AVLnode) Setdata(data interface{}) {

	avlnode.Data = data
}
func (avlnode *AVLnode) GetLeft() *AVLnode {
	if avlnode == nil {
		return nil
	}
	return avlnode.Left
}
func (avlnode *AVLnode) GetHeight() int {
	if avlnode == nil {
		return 0
	}
	return avlnode.height
}
func (avlnode *AVLnode) GetRight() *AVLnode {
	if avlnode == nil {
		return nil
	}
	return avlnode.Right
}
