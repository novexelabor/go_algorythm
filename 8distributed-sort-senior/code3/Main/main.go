package main

import (
	"fmt"
	"time"
)

func add(ch chan int) { //向channel中写入数据
	for i := 0; i < 3; i++ {
		ch <- i
	}
	time.Sleep(10 * time.Second)
	ch <- 4
}

func main1() {
	ch := make(chan int, 10)
	go add(ch) //协程写入数据
	for {
		select { //那个管道有值读取那个，都有值就随机处理
		case <-time.After(2 * time.Second):
			fmt.Println("timeout")
			return
		case v := <-ch:
			fmt.Println(v) // if ch not empty, time.After will nerver exec
			fmt.Println("sleep one seconds ...")
			time.Sleep(1 * time.Second)
			fmt.Println("sleep one seconds end...")
		}
	}
}

func main() {

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	ch := make(chan int, 10)
	go add(ch)
	for {
		select {
		case v := <-ch:
			fmt.Println(v) // if ch not empty, time.After will nerver exec
			fmt.Println("sleep one seconds ...")
			time.Sleep(1 * time.Second)
			fmt.Println("sleep one seconds end...")
		default: // forbid block
		}

		select {
		case <-ticker.C:
			fmt.Println("timeout")
			return
		default: // forbid block
		}
	}

}
