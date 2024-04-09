package main

import "fmt"

//A找到第一个等于3的
//B找到最后一个等于3
//C找到第一个大于等于2
//D找到最后一个小于7的数据
//arr:=[]int {1,2,3,3,4,5,6,6,6,6,7,9,10}

//B找到最后一个等于3
func bin_searchB(arr []int, data int) int {
	low := 0             //最下方
	high := len(arr) - 1 //最上方
	index := -1          //索引index

	for low <= high { //循环的终止条件

		mid := (low + high) / 2

		if arr[mid] > data {
			high = mid - 1
		} else if arr[mid] < data {
			low = mid + 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] != data {
				index = mid
				fmt.Println("mid", mid)
				break
			} else {
				low = mid + 1 //递归继续查找 ,直到最后一个相等的
			}

		}
	}
	return index

}

////A找到第一个等于3的
func bin_searchA(arr []int, data int) int {
	low := 0             //最下方
	high := len(arr) - 1 //最上方
	index := -1          //索引index

	for low <= high { //循环的终止条件

		mid := (low + high) / 2

		if arr[mid] > data {
			high = mid - 1
		} else if arr[mid] < data {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != data {
				index = mid
				fmt.Println("mid", mid)
				break
			} else {
				high = mid - 1 //递归继续查找
			}

		}
	}
	return index

}

//C找到第一个大于等于3
//arr:=[]int {1,2,3,3,3,3,3,4,5,6,6,6,6,7,9,10}
func bin_searchC(arr []int, data int) int {
	low := 0             //最下方
	high := len(arr) - 1 //最上方
	index := -1          //索引index

	for low <= high { //循环的终止条件

		mid := (low + high) / 2

		if arr[mid] < data {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] < data { //临界点判断
				index = mid
				break
			} else {
				high = mid - 1 //正常遍历
			}

		}
	}
	return index

}

//D找到最后一个小于6的数据
func bin_searchD(arr []int, data int) int {
	low := 0             //最下方
	high := len(arr) - 1 //最上方
	index := -1          //索引index

	for low <= high { //循环的终止条件

		mid := (low + high) / 2

		if arr[mid] > data {
			high = mid - 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] > data { //临界点判断
				index = mid
				break
			} else {
				low = mid + 1
			}

		}
	}
	return index

}
func main() {
	//          0 1 2 3 4 5 6 7 8 9101112
	arr := []int{1, 2, 3, 3, 3, 3, 3, 4, 5, 6, 6, 6, 6, 7, 9, 10}
	for i := 0; i < len(arr); i++ {
		fmt.Println("index", i, arr[i])
	}
	fmt.Println(bin_searchD(arr, 7))
	fmt.Println(arr[2])
}
