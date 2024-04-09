package main

import "fmt"

//分布式编程
//----   ----   ----
func ThirdSearch(arr []int, data int) int {
	low := 0
	high := len(arr) - 1 //确定底部与高步
	i := 0
	for low <= high {
		mid1 := low + int((high-low)/3)
		mid2 := high - int((high-low)/3) // 中间两个节点下标

		i++
		fmt.Println("third", i) //分裂一次加一

		middata1 := arr[mid1]
		middata2 := arr[mid2]

		if middata1 == data {
			return mid1
		} else if middata2 == data {
			return mid2
		}

		if middata1 < data {
			low = mid1 + 1
		} else if middata2 > data {
			high = mid2 - 1
		} else {
			low = low + 1 //该部分是处理寻找不到的时候，退出循环条件low > high情况
			high = high - 1
		}

	}

	return -1

}

func bin_searchq(arr []int, data int) int {
	low := 0             //最下方
	high := len(arr) - 1 //最上方
	i := 0
	for low <= high { //循环的终止条件
		i++
		fmt.Println("bin", i)
		mid := (low + high) / 2

		if arr[mid] > data {
			high = mid - 1
		} else if arr[mid] < data {
			low = mid + 1
		} else {
			return mid //找到
		}
	}
	return -1

}

func main() {
	arr := make([]int, 1000, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	fmt.Println(ThirdSearch(arr, 999))
	fmt.Println(bin_searchq(arr, 999))
}
