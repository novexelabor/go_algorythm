package main

import "fmt"

func main() {
	rbtree := NewRBTree()
	for i := 0; i < 1000000; i++ {
		rbtree.Insert(Int(i))
	}
	for i := 0; i < 900000; i++ {
		//rbtree
		rbtree.Delete(Int(i))
	}
	fmt.Println(rbtree.GetDepth())
}
