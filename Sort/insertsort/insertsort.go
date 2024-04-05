package main

import "fmt"

//插入排序，比较之后，从前往后移动

func InsertSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 1; i < length; i++ {
			backup := arr[i]
			j := i - 1
			for j >= 0 && backup < arr[j] { //只设置条件
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = backup
			fmt.Println(arr)
		}
		return arr
	}
}

func main() {
	arr := []int{1, 19, 29, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(InsertSort(arr))
}
