package main

import "fmt"

func ShellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start + gap; i < length; i += gap { //插入排序的变种
		backup := arr[i]
		j := i - gap
		for j >= 0 && arr[j] > backup {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = backup
	}
}

func ShellSort(arr []int) []int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		gap := length / 2 //gap是除以2更迭的
		for gap > 0 {
			for i := 0; i < gap; i++ { //处理每个元素的步长
				ShellSortStep(arr, i, gap)
			}
			//gap-- //gap--
			gap /= 2
		}

	}

	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	//fmt.Println(SelectSortMax(arr))
	fmt.Println(ShellSort(arr))

}
