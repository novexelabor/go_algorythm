package main

import "fmt"

func cocktail(arr []int) []int {
	swap := false
	for i := 0; i < len(arr)/2; i++ {
		left := 0
		right := len(arr) - 1
		for left <= right {
			if arr[left] > arr[left+1] {
				arr[left], arr[left+1] = arr[left+1], arr[left]
				swap = true
			}
			left++
			if arr[right] < arr[right-1] {
				arr[right], arr[right-1] = arr[right-1], arr[right]
				swap = true
			}
			right--
		}
		fmt.Println(i, arr)
		if !swap {
			break
		}
	}
	return arr
}

//可以进一步的优化，如果是提前排好序了，添加一个交换的标志
//没有交换了，代表已经排好了，可以不用继续遍历了

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(cocktail(arr))
}
