package main

import "fmt"

type Table struct {
	data   []int
	Length int
	head   *Table
}

func mainTable() {
	t1 := &Table{[]int{1, 2, 3}, 3, &Table{[]int{1, 2, 3}, 3, nil}}
	fmt.Println(t1)
}
