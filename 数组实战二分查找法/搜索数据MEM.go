package main

import "fmt"

// 3  2 9  1  5  7
//3
//21  3  957
//1 2  3

//57 9

func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr //只有一元素无需排序
	} else {
		splitdata := arr[0] //第一个数字
		low := make([]int, 0, 0)
		high := make([]int, 0, 0)
		mid := make([]int, 0, 0)
		mid = append(mid, splitdata) //保存分离的数据
		//数据分为三段处理，分别是大于，等于，小于
		for i := 1; i < length; i++ {
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high)   //递归循环
		myarr := append(append(low, mid...), high...) //数据归并
		return myarr

	}
}

func bin_search(arr []int, data int) int {
	left := 0
	right := len(arr) - 1 //最下最上面
	for left < right {
		mid := (left + right) / 2
		if arr[mid] > data {
			right = mid - 1
		} else if arr[mid] < data {
			left = mid + 1 //移动
		} else {
			return mid //找到
		}
	}
	return -1
}

func main() {
	arr := []int{1, 19, 4, 8, 3, 5, 4, 6, 19, 0}
	fmt.Println("未曾排序", arr)
	fmt.Println("已经排序", QuickSort(arr))
	index := bin_search(arr, 19)
	if index == -1 {
		fmt.Println("没有找到")
	} else {
		fmt.Println(arr[index], index, "找到")
	}
}
