package main

import (
	"fmt"
	//"math/rand"
	//"testing"
	//"time"
)

//B树的节点
type BtreeNode struct {
	Leaf     bool         //是否叶子
	N        int          //分支的数量
	keys     []int        //存储数据
	Children []*BtreeNode //指向自己的多个分支节点

}

//新建一个节点
func NewBtreeNode(n int, branch int, leaf bool) *BtreeNode {
	return &BtreeNode{leaf,
		n,
		make([]int, branch*2-1), //n个branch对应2n,root  2n-1 ，确保最大的存储容量
		make([]*BtreeNode, branch*2)}
}

//搜索B树枝的节点
func (btreenode *BtreeNode) Search(key int) (mynode *BtreeNode, idx int) {
	i := 0
	//找到合适的位置，找到最后一个小于key的,i之后的就是大于等于
	for i < btreenode.N && btreenode.keys[i] < key {
		i += 1
	}
	if i < btreenode.N && btreenode.keys[i] == key {
		mynode, idx = btreenode, i //找到 ,节点内的索引,节点内的数据是有序递增的
	} else if btreenode.Leaf == false {
		//进入孩子叶子继续搜索
		mynode, idx = btreenode.Children[i].Search(key) //小于btreenode.keys[i]的值
	}
	return
}
func (parent *BtreeNode) Split(branch int, idx int) {
	full := parent.Children[idx]                         //孩子节点
	newnode := NewBtreeNode(branch-1, branch, full.Leaf) //新建一个节点，备份2*n-1,0_n-2;n_2n-2
	for i := 0; i < branch-1; i++ {
		newnode.keys[i] = full.keys[i+branch] //数据移动，跳过一个分支
		newnode.Children[i] = full.Children[i+branch]
	}
	newnode.Children[branch-1] = full.Children[branch*2-1] //处理最后,一个节点分裂成两个,新节点保存最后的孩子分支
	full.N = branch - 1                                    //减少了branch-1个分支
	//x新增一个key到children
	for i := parent.N; i > idx; i-- { //此处与原代码不同，做了修改,将children与keys对换了
		parent.Children[i+1] = parent.Children[i] //让下标是idx的位置腾出一个空位置
		parent.keys[i] = parent.keys[i-1]         //从后往前移动,相应的关键值向后移动
	}
	parent.keys[idx] = full.keys[branch-1]
	parent.Children[idx+1] = newnode //插入数据，增加总量
	parent.N++

}

//节点插入数据
func (btreenode *BtreeNode) InsertNonFull(branch int, key int) {
	if btreenode == nil {
		return
	}
	i := btreenode.N    //记录叶子节点的总量
	if btreenode.Leaf { //是叶子节点
		for i > 0 && key < btreenode.keys[i-1] {
			btreenode.keys[i] = btreenode.keys[i-1] //从后往前移动，
			i--                                     //i从后往前移动
		}
		btreenode.keys[i] = key //插入数据
		btreenode.N++           //总量加一

	} else { //不是叶子节点
		for i > 0 && key < btreenode.keys[i-1] {
			i-- //i从后往前移动
		}
		c := btreenode.Children[i] //找到下标,将key的值插入到该节点的keys切片中
		if c != nil && c.N == 2*branch-1 {
			btreenode.Split(branch, i) //切割，i是children中的对应的下标
			if key > btreenode.keys[i] {
				i++
			}
		}
		btreenode.Children[i].InsertNonFull(branch, key) //递归插入孩子叶子

	}

}

//节点显示为字符串
func (btreenode *BtreeNode) String() string {
	return fmt.Sprintf("{n=%d,leaf=%v,Children=%v}\n", btreenode.N, btreenode.keys, btreenode.Children)
}

//B树
type Btree struct {
	Root   *BtreeNode //根节点
	branch int        //分支的数量
}

//插入
func (tree *Btree) Insert(key int) {
	root := tree.Root //根节点
	if root.N == 2*tree.branch-1 {
		s := NewBtreeNode(0, tree.branch, false)
		tree.Root = s //新建一个节点备份根节点
		s.Children[0] = root
		s.Split(tree.branch, 0) //拆分整合
		root.InsertNonFull(tree.branch, key)
	} else {
		root.InsertNonFull(tree.branch, key)
	}
}

//查找
func (tree *Btree) Search(key int) (n *BtreeNode, idx int) {
	return tree.Root.Search(key)
}

//返回字符串
func (tree *Btree) String() string {

	return tree.Root.String() //返回树的字符串
}

//新建B树
func NewBtree(branch int) *Btree {
	return &Btree{NewBtreeNode(0, branch, true), branch}
}

func main() {
	// mybtree := NewBtree(10000)
	// for i := 10000; i > 0; i-- {
	// 	mybtree.Insert(rand.Int() % 10000)
	// }
	// fmt.Println(mybtree.String())
	// for i := 0; i < 10; i++ {
	// 	starttime := time.Now()
	// 	fmt.Println(mybtree.Search(i))
	// 	fmt.Println("一共用了", time.Since(starttime))
	// }

	//测试看结构是否正确
	parent := NewBtreeNode(0, 4, true)
	btreetest := &Btree{parent, 4}
	for i := 0; i < 10; i++ {
		btreetest.Insert(i)
	}

	fmt.Println(btreetest.String())
	// 	fmt.Println(btreetest.Search(2))
}
