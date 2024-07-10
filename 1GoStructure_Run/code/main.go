package main

import (
	"fmt"
	"go_algorythm/1GoStructure_Run/code/goDataStructure"
)

func main() {
	myq := goDataStructure.CreateArrayQueue(1000000)
	fmt.Println(goDataStructure.TestQueue(myq, 1000000))

}
