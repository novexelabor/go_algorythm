package main

import "fmt"

//7  0  8  -1  16  4
//0 7   8  -1  16  4
//0 7  -1   8 16  4
// -1 0       4  7  8 16
//顺序的移动排序
func GnomeSort(arr []int) []int {
	i := 1
	for i < len(arr) {
		if arr[i] >= arr[i-1] {
			i++ //符合顺序，继续前进
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			if i > 1 { //确定下标最小的范围
				i--
			}
			fmt.Println(arr)

		}
	}
	return arr

}

func main() {
	//快速排序排列100000个，插入3个数据，在排序用gnome,
	arr := []int{11, 2, 3, 23, 33, 3, 13, 4, 15, 6, 6, 61, 6, 17, 9, 10}
	GnomeSort(arr)
	fmt.Println(arr)

}
