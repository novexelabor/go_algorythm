package main

import "fmt"

func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else { //选择逻辑
		for i := 0; i < length-1; i++ {
			min := i
			for j := i; j < length; j++ {
				if arr[min] > arr[j] {
					min = j
				}
			}
			if min != i {
				arr[i], arr[min] = arr[min], arr[i] //交换数据
			}
			fmt.Println(arr)
		}
		return arr
	}

}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	//fmt.Println(SelectSortMax(arr))
	fmt.Println(SelectSort(arr))
}
