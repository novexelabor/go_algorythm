package main

import "fmt"

func main() {
	bst := NewBinaryTree()

	node1 := &Node{4, nil, nil}
	node2 := &Node{2, nil, nil}
	node3 := &Node{6, nil, nil}
	node4 := &Node{1, nil, nil}
	node5 := &Node{3, nil, nil}
	node6 := &Node{5, nil, nil}
	node7 := &Node{7, nil, nil}
	node8 := &Node{17, nil, nil}
	bst.Root = node1

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	node6.Left = node8

	bst.Size = 8
	nodelast := bst.FindlowerstAncestor(bst.Root, node6, node7)
	fmt.Println(nodelast)
	fmt.Println(bst.GetDepth(bst.Root))

}

func main2() {
	bst := NewBinaryTree()

	node1 := &Node{4, nil, nil}
	node2 := &Node{2, nil, nil}
	node3 := &Node{6, nil, nil}
	node4 := &Node{1, nil, nil}
	node5 := &Node{3, nil, nil}
	node6 := &Node{5, nil, nil}
	node7 := &Node{7, nil, nil}
	bst.Root = node1

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7
	bst.Size = 7
	fmt.Println("中序----------------------")
	bst.InOrder()
	fmt.Println(bst.InOrderNoRecursion())
	fmt.Println("前序----------------------")
	bst.PreOrder()
	fmt.Println(bst.PreOrderNoRecursion())
	fmt.Println("后序----------------------")
	bst.PostOrder()
	fmt.Println(bst.PostOrderNoRecursion())
	fmt.Println("level----------------------")
	//bst.Levelshow()
	bst.Stackshow(bst.Root)
}
func main1() {
	bst := NewBinaryTree()

	node1 := &Node{4, nil, nil}
	node2 := &Node{2, nil, nil}
	node3 := &Node{6, nil, nil}
	node4 := &Node{1, nil, nil}
	node5 := &Node{3, nil, nil}
	node6 := &Node{5, nil, nil}
	node7 := &Node{7, nil, nil}
	bst.Root = node1

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7
	bst.Size = 7

	//for i:=1;i<=7;i++{
	//	bst.Add(i)
	//}
	//bst.Add(4)
	//bst.Add(6)
	//bst.Add(5)
	//bst.Add(7)
	//bst.Add(2)
	//bst.Add(1)
	//bst.Add(3)
	fmt.Println(bst.FindMax(), "max")
	fmt.Println(bst.FindMin(), "min")
	//fmt.Println(bst.Isin(3))
	//fmt.Println(bst.Isin(31))
	//fmt.Println(bst.RemoveMin())
	//fmt.Println(bst.RemoveMax())
	bst.Remove(4)
	fmt.Println("中序----------------------")
	bst.InOrder()
	fmt.Println("前序----------------------")
	//bst.PreOrder()
	fmt.Println("后序----------------------")
	//bst.PostOrder()
	fmt.Println("last ----------------------")
	fmt.Println(bst.String())
}
