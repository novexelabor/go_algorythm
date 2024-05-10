package main

import "fmt"

type TreeNode struct {
	element interface{} //数据
	left    *TreeNode   //左右节点
	right   *TreeNode
	npl     int //级别
}
type PQ *TreeNode //优先队列,重命名了

//开辟一个节点，这个是没有头节点的
func NewLeftHeap(element interface{}) PQ {
	head := new(TreeNode)  //开辟内存
	head.element = element //数据初始化
	head.left = nil
	head.right = nil
	head.npl = 0
	return PQ(head)
}

//H1.element.(int) < H2.element.(int)---求的是极小值啊

//此处函数的目的就是将H2插入到H1合适的位置上，显然判断左右子树

func MergeSort(H1, H2 PQ) PQ { //返回的是左边的第一个
	if H1.left == nil { //此处是左节点为nil，右节点一定是nil
		H1.left = H2 //直接插入
	} else { //当左节点非空时，右节点插入了数值，左右皆有节点时，层级加一
		H1.right = Merge(H1.right, H2)  //递归合并,这里是根据数值
		if H1.left.npl < H1.right.npl { //处理层级互相切换，左右节点之间，那个层级低，下次插入节点放在那个子树上
			H1.left, H1.right = H1.right, H1.left
		}
		H1.npl = H1.right.npl + 1 //级递增 层,
	}
	return H1 //返回左边第一个节点参数
}

//确保有序
func Merge(H1, H2 PQ) PQ { //合并
	if H1 == nil {
		return H2
	}
	if H2 == nil {
		return H1
	}
	//递归排序,>取得极大，<取得极小
	if H1.element.(int) < H2.element.(int) { //类型转换
		return MergeSort(H1, H2) //不会确保左右之间的大小，均是小于父亲节点的

	} else {
		return MergeSort(H2, H1) //插入的是H1

	}
}

func Insert(data interface{}, H PQ) PQ {
	insertnode := new(TreeNode) //新建一个节点
	insertnode.element = data
	insertnode.left = nil
	insertnode.right = nil
	insertnode.npl = 0
	//插入用归并实现
	H = Merge(insertnode, H)
	return H

}
func DeleteMin(H PQ) (PQ, interface{}) { //每次取根节点的数值，左右子树merge()
	if H == nil {
		return nil, nil
	} else {
		leftHeap := H.left   //左边节点
		rightHeap := H.right //右边节点
		value := H.element   //取出顶端的数据
		H = nil
		return Merge(leftHeap, rightHeap), value
	}
}

//遍历大树
func PrintHQ(H PQ) {
	if H == nil {
		return
	}
	PrintHQ(H.left)
	fmt.Println(H.element, "  ")
	PrintHQ(H.right)
	//fmt.Println(H.element, "  ") //后序遍历树节点
}

func main() {
	H := NewLeftHeap(3)
	H = Insert(2, H)
	H = Insert(1, H)
	H = Insert(12, H)
	H = Insert(4, H)
	H = Insert(5, H)
	PrintHQ(H)
	H, data := DeleteMin(H)
	fmt.Println("min", data)
	PrintHQ(H)
	H, data = DeleteMin(H)
	fmt.Println("min", data)
	PrintHQ(H)

}
