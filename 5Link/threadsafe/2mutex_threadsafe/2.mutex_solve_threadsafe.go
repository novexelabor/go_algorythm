package main

import (
	"fmt"
	"sync"
	"time"
)

//线程安全，多个线程访问同一个资源，产生资源竞争，最终结果不正确
var money int = 0
var lock *sync.RWMutex = new(sync.RWMutex) //初始化

func add(pint *int) {
	lock.Lock() //多个协程调用的全局变量，需要添加互斥量mutex，避免出错
	for i := 0; i < 100000; i++ {
		*pint++
	}
	lock.Unlock() //执行完之后解锁
}

func main() {
	for i := 0; i < 1000; i++ {
		go add(&money) //开启协程，调用函数
	}
	time.Sleep(time.Second * 20)
	fmt.Println(money)
}
