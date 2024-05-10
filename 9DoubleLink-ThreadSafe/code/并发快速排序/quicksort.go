package main

import (
	"fmt"
	"runtime"
)

//    arr需要排序的数组，lastarr排序排列号，级别，线程数量
func QuickSortThread(arr []int, lastarr chan int, level int, threads int) {
	level = level * 2 //每加深一个级，多一个线程,最下面一层的个数2^n,2的幂
	if len(arr) == 0 {
		close(lastarr) //关闭管道
		return
	} else if len(arr) == 1 {
		lastarr <- arr[0] //为一个数据放入管道
		close(lastarr)    //关闭管道
		return
	} else {
		less := make([]int, 0)        //比我小的数据
		greater := make([]int, 0)     //比我大的数据
		midder := make([]int, 0)      //与我相等的数据
		left := arr[0]                //取得第一个数据
		midder = append(midder, left) //中间存放相等数据
		for i := 1; i < len(arr); i++ {
			if arr[i] < left {
				less = append(less, arr[i]) //处理小于的
			} else if arr[i] > left {
				greater = append(greater, arr[i]) //处理小于的
			} else {
				midder = append(midder, arr[i]) //处理等于
			}
		}
		left_ch := make(chan int, len(less))
		right_ch := make(chan int, len(greater)) //存放数组的管道
		fmt.Println("level", level)
		if level <= threads { //如果线程超过执行数量，顺序调用，否则并发调用
			go QuickSortThread(less, left_ch, level, threads)
			go QuickSortThread(greater, right_ch, level, threads)
		} else {
			QuickSortThread(less, left_ch, level, threads)
			QuickSortThread(greater, right_ch, level, threads)
		}
		//数据压入管道,当前的线程获取左右channel的值
		for i := range left_ch {
			lastarr <- i //向当前线程的channel里传送数值
		}
		for _, v := range midder {
			lastarr <- v
		}
		for i := range right_ch {
			lastarr <- i
		}
		close(lastarr) //关闭管道

		return

	}
}

func main() {

	arr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5, 10}

	num := runtime.NumCPU()
	fmt.Println("最大的线程数", num) //用的CPU的核数

	lastarr := make(chan int) //管道
	go QuickSortThread(arr, lastarr, 1, 1)
	for v := range lastarr { //显示管道的每一个数据
		fmt.Println(v)
	}

}
