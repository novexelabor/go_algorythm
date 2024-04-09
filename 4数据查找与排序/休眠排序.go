package main

import (
	"fmt"
	"time"
)

//  5 1  3  2  4
//  5 1  3   2 4
//写入 1 us
//1亿数据
//多线程，分布式
var flag bool
var container chan bool //通道传递的值类型,变量声明
var count int
var arr chan int

//全局变量

func main() {

	var array []int = []int{16, 8, 1, 24, 30}
	var arr1 []int
	arr1 = make([]int, 0)

	flag = true                    //标识，区分
	container = make(chan bool, 5) //5个管道
	arr = make(chan int, len(array))
	for i := 0; i < len(array); i++ {
		go tosleep(array[i])
	}
	go listen(len(array), &arr1) //并发执行
	for flag {
		time.Sleep(1 * time.Second)
	}
	//time.Sleep(5 * time.Second) //等待1秒，读取完通道的数值
	fmt.Println(arr1)
}

func listen(size int, arr1 *[]int) { //channel通道通信

	for flag { //循环处理
		//获取通道数值
		a := <-arr
		*arr1 = append(*arr1, a)

		select { //select通道监听
		case <-container:
			count++            //计数器
			if count >= size { //等待5个数字采集完成就退出
				flag = false
				break
			}
		}
	}

}

func tosleep(data int) {
	time.Sleep(time.Duration(data) * time.Microsecond * 1000) //通过并发执行，多线程，通过睡眠时间的长短来排序
	fmt.Println("sleep", data)
	arr <- data
	container <- true //管道输入ok
}
