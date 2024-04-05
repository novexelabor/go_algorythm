package main

import "fmt"

//1,7,3,8,2,9,5,6,4,0
//1   3   2   5   4
//  7   8   9   6   0

// 9  2   1  6 0 7
// 2  9   1  6 0 7  所有奇数位与身后的相邻的偶数位比较交换
// 2   1  9  0  6 7所有偶数位与身后的相邻的奇数位比较交换
//1   2   09   6 7  所有奇数位与身后的相邻的偶数位比较交换
//1  0    2 6  9 7所有偶数位与身后的相邻的奇数位比较交换
//0  1   2  6  7 9  所有奇数位与身后的相邻的偶数位比较交换

//不停的奇偶位交换，直到不再交换完成排序
//for循环奇偶位交换，先一边奇数位交换，然后一次偶数位交换

func OddEven(arr []int) []int {

	isSorted := false //奇数位，偶数位都不需要排序的有序

	for isSorted == false {
		isSorted = true
		for i := 1; i < len(arr)-1; i = i + 2 { //奇数位
			if arr[i] > arr[i+1] { //需要交换换
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		fmt.Println("1", arr)
		for i := 0; i < len(arr)-1; i = i + 2 { //偶数位
			if arr[i] > arr[i+1] { //需要交换换
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		fmt.Println("0", arr)

	}
	return arr

}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	//fmt.Println(SelectSortMax(arr))
	fmt.Println(OddEven(arr))

}
