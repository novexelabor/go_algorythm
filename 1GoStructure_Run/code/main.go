package main

import (
	"./goDataStructure"
	"fmt"
)



func main(){
	myq:=goDataStructure.CreateArrayQueue(1000000)
	fmt.Println(goDataStructure.TestQueue(myq, 1000000))



}
