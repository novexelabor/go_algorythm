package main

import "fmt"

//最大的沉到底部,像冒泡一样
func Bubblesort(arr []int) []int {
	length := len(arr) //求数组长度
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ { //只剩一个，不需要冒泡了
			isneedexchange := false
			for j := 0; j < length-1-i; j++ { //每次循环，确定的位置是最后的位置
				if arr[j] > arr[j+1] { //两两比较
					arr[j], arr[j+1] = arr[j+1], arr[j]
					isneedexchange = true
				}
			}
			if !isneedexchange {
				break
			}
			fmt.Println(arr)

		}

		return arr
	}
}

func main() {
	arr := []int{11, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	//fmt.Println(BubbleFindMax(arr))
	fmt.Println(Bubblesort(arr))

}
