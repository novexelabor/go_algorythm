package main

import "fmt"

func main() {

	isok := AIoutStack(AIdata, 0, 0)
	if isok {
		fmt.Println("可以走出")
		show(AIdata)
		fmt.Println("---------开始移动-----------")
		//AImoveOut()
		//show(data)
	} else {
		fmt.Println("走不出")
	}
}

func main1x() {
	show(data)
	for {
		var inputstr string
		fmt.Scanln(&inputstr)
		run(inputstr)
	}
}
