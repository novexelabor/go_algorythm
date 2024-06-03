package main

import (
	"container/heap"
	"fmt"
)

type HuffmanTree interface {
	Freq() int //哈夫曼树的接口
}

//哈夫曼叶子类型
type HuffmanLeaf struct {
	freq  int  //频率
	value rune //int32
}

//哈夫曼树的类型
type HuffmanNode struct {
	freq        int         //频率
	left, right HuffmanTree //具体结构体类型实现了Freq(),即可
}

func (self HuffmanLeaf) Freq() int { //频率
	return self.freq
}
func (self HuffmanNode) Freq() int { //频率
	return self.freq
}

type treeHeap []HuffmanTree

//求长度
func (th treeHeap) Len() int {
	return len(th) //内置函数
}

//比较函数
func (th treeHeap) Less(i int, j int) bool {
	return th[i].Freq() < th[j].Freq()
}

//压入
func (th *treeHeap) Push(ele interface{}) {
	*th = append(*th, ele.(HuffmanTree))
}

//弹出
func (th *treeHeap) Pop() (po interface{}) {
	po = (*th)[len(*th)-1]
	*th = (*th)[:len(*th)-1] //二次指针,*访问,删除操作
	return
}

func (th treeHeap) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
}

func BuildTree(symFreqs map[rune]int) HuffmanTree {
	var trees treeHeap
	for c, f := range symFreqs {
		trees = append(trees, HuffmanLeaf{f, c}) //叠加数据
	}
	heap.Init(&trees) //开始使用堆。切片堆栈的使用
	//该哈夫曼树的构造并不是严谨的，不是最小的代价
	for trees.Len() > 1 {
		a := heap.Pop(&trees).(HuffmanTree)
		b := heap.Pop(&trees).(HuffmanTree)
		heap.Push(&trees, HuffmanNode{a.Freq() + b.Freq(), a, b})
	} //构造哈夫曼树
	return heap.Pop(&trees).(HuffmanTree)

}
func showtimes(tree HuffmanTree, prefix []byte) {
	switch i := tree.(type) {
	case HuffmanLeaf:
		fmt.Printf("%c\t%d\n", i.value, i.freq) //打印数据与频率
	case HuffmanNode:
		prefix = append(prefix, '0')
		showtimes(i.left, prefix)       //递归到左子树
		prefix = prefix[:len(prefix)-1] //删除最后一个

		prefix = append(prefix, '1')
		showtimes(i.right, prefix) //递归到右子树
		prefix = prefix[:len(prefix)-1]

	}
}

func showcodes(tree HuffmanTree, prefix []byte) {
	switch i := tree.(type) {
	case HuffmanLeaf:
		fmt.Printf("%s---\n", string(prefix)) //打印数据与频率
	case HuffmanNode:
		prefix = append(prefix, '0')
		showcodes(i.left, prefix)       //递归到左子树
		prefix = prefix[:len(prefix)-1] //删除最后一个

		prefix = append(prefix, '1')
		showcodes(i.right, prefix) //递归到右子树
		prefix = prefix[:len(prefix)-1]

	}
}

func main() {
	stringcode := "aaaaaaabbccccddddeffgggg"
	fmt.Println("stringcode:", stringcode)
	symFreqs := make(map[rune]int)
	for _, c := range stringcode { //使用字典来统计重复字符的个数rune类型
		symFreqs[c]++ //统计频率
	}
	trees := BuildTree(symFreqs)
	showtimes(trees, []byte{})
	fmt.Printf("-------------------\n")
	showcodes(trees, []byte{})
	//fmt.Println(symFreqs)
	//fmt.Println("a",int('a'))
	//fmt.Println("a",int('b'))
}
