package main

import "fmt"

//1,9,2,8,3,   7,4,6,5,10
//1,9,2,   8,3,        7,4,6,5,10
//1,2,   9,
//3,8
//1  2  3  8  9     //  456 7 10
//12345678910
func merge(leftarr []int, rightarr []int) []int {
	leftindex := 0     //左边索引
	rightindex := 0    //右边索引
	lastarr := []int{} //最终的数组  空间复杂度为n
	for leftindex < len(leftarr) && rightindex < len(rightarr) {
		if leftarr[leftindex] < rightarr[rightindex] {
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++

		} else if leftarr[leftindex] > rightarr[rightindex] {
			lastarr = append(lastarr, rightarr[rightindex])
			rightindex++
		} else {
			lastarr = append(lastarr, rightarr[rightindex])
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++
			rightindex++
		}
	}
	for leftindex < len(leftarr) { //把没有结束的归并过来
		lastarr = append(lastarr, leftarr[leftindex])
		leftindex++
	}
	for rightindex < len(rightarr) { //把没有结束的归并过来
		lastarr = append(lastarr, rightarr[rightindex])
		rightindex++
	}
	return lastarr
}

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr //小与10改用插入排序
	} else if length > 1 && length < 5 {
		return InsertSortX(arr)
	} else {
		mid := length / 2
		leftarr := MergeSort(arr[:mid])
		rightarr := MergeSort(arr[mid:])
		//递归合并
		return merge(leftarr, rightarr)
	}
}

func InsertSortX(arr []int) []int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		for i := 1; i < length; i++ { //跳过第一个
			backup := arr[i] //备份插入的数据
			j := i - 1       //上一个位置循环找到位置插入
			for j >= 0 && backup < arr[j] {
				arr[j+1] = arr[j] //从前往后移动
				j--
			}
			arr[j+1] = backup //插入
			fmt.Println(arr)
		}

		return arr

	}
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(MergeSort(arr))
}
