package main

import (
	"fmt"
	"math/rand"
)

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
		//递归处理，利用返回 排序好的数组 ---带有返回值的递归
		low, high = QuickSort(low), QuickSort(high) //切割递归处理
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}

func QuickSort2(arr []int) []int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		//n:=length-1//n>=0  && n<=length-1
		n := rand.Int() % length     //随机获取一个数值,计算得到随地的数组基准
		splitdata := arr[n]          //以第一个为基准
		low := make([]int, 0, 0)     //存储比我小的
		high := make([]int, 0, 0)    //存储比我大的
		mid := make([]int, 0, 0)     //存储与我相等
		mid = append(mid, splitdata) //加入第一个相等

		for i := 0; i < length; i++ {
			if i == n {
				continue
			}
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

func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10, 67}
	//fmt.Println(SelectSortMax(arr))
	fmt.Println(QuickSort2(arr))

}
