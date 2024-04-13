package main

import (
	"fmt"
	"go_algorythm/5Link/threadsafe/Queue"
)

func main() {
	fmt.Println("go")
	h := Queue.NewMin()
	h.Insert(Queue.Int(8))

	h.Insert(Queue.Int(9))
	h.Insert(Queue.Int(7))
	h.Insert(Queue.Int(5))
	h.Insert(Queue.Int(6))
	h.Insert(Queue.Int(4))
	h.Insert(Queue.Int(2))
	h.Insert(Queue.Int(3))
	fmt.Println("insert")
	fmt.Println(h.Extract().(Queue.Int))
	fmt.Println(h.Extract().(Queue.Int))
	fmt.Println(h.Extract().(Queue.Int))
	fmt.Println(h.Extract().(Queue.Int))

}
