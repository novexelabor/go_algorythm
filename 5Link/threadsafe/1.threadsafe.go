package main

import (
	"fmt"
	"time"
	//"time"
)

//线程安全，多个线程访问同一个资源，产生资源竞争，最终结果不正确
var money int = 0

func add(pint *int) {
	for i := 0; i < 100000; i++ {
		*pint++
	}
}

func main() {
	for i := 0; i < 1000; i++ {
		add(&money)
	}
	time.Sleep(time.Second * 20)
	fmt.Println(money)
}
