package main

import (
	"fmt"
	"go_algorythm/5Link/threadsafe/Queue"
)

func main() {
	h := Queue.NewMinPriorityQueue()
	h.Insert(*Queue.NewPriorityItem(101, 11))
	h.Insert(*Queue.NewPriorityItem(102, 12))
	h.Insert(*Queue.NewPriorityItem(103, 15))
	h.Insert(*Queue.NewPriorityItem(104, 14))
	h.Insert(*Queue.NewPriorityItem(105, 13))
	h.Insert(*Queue.NewPriorityItem(106, 19))

	fmt.Println(h.Extract())
	fmt.Println(h.Extract())
	fmt.Println(h.Extract())

}
