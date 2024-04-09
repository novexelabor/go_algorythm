package main

import "fmt"

func QuickSort(arr []int) []int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		splitdata := arr[0]          //以第一个为基准
		low := make([]int, 0, 0)     //存储比我小的
		high := make([]int, 0, 0)    //存储比我大的
		mid := make([]int, 0, 0)     //存储与我相等
		mid = append(mid, splitdata) //加入第一个相等

		for i := 1; i < length; i++ {
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high) //切割递归处理
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}
func bin_search(arr []int, data int) int {
	low := 0             //最下方
	high := len(arr) - 1 //最上方

	for low <= high { //循环的终止条件
		fmt.Println(arr[low : high+1])
		mid := (low + high) / 2
		fmt.Println("mid", mid)
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
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	//fmt.Println(SelectSortMax(arr))
	//fmt.Println(QuickSort(arr))
	fmt.Println("未排序", arr)
	arr = QuickSort(arr)
	fmt.Println("已排序", arr)
	index := bin_search(arr, 4)
	fmt.Println("index", index)
	if index == -1 {
		fmt.Println("找不到")
	} else {
		fmt.Println("找到")
	}
}
