package main

import "fmt"

// QuickSort 快速排序, 不稳定的
// 最好时间复杂度O(nlogn), 最坏时间复杂度O(n^2), 平均时间复杂度O(nlogn), 空间复杂度O(1)
func QuickSortK(arr []int) {
	quickSortGo(arr, 0, len(arr)-1)
}

func Swap(arr []int, i int, j int) { //数据交换
	arr[i], arr[j] = arr[j], arr[i]
}

// quickSortGo 快速排序的算法部分
func quickSortGo(arr []int, left, right int) {
	// 当q=0时, 下一个分区是(0,-1), 所以出现p小于r的情况,说明分区只有一个元素了可以直接返回.
	if left >= right {
		//fmt.Println()
		return
	}
	q := partition(arr, left, right) //返回的切段数据

	quickSortGo(arr, left, q-1)  //前段
	quickSortGo(arr, q+1, right) //后段
}

// partition 分区函数. 对元素进行分区, 将大于等于(小于等于)arr[r]的元素, 全部放在arr[r]的左边, 小于(大于)arr[r]的元素放在右边
// @return 返回分区标志的下标
func partition(arr []int, left, right int) int {
	//注: 以从递增排序为例注释.
	//将最后一个元素作为分区标志
	pivot := right

	//12542
	//12 54 2
	//i始终指向从左数起, 第一个大于arr[pivot]的元素下标
	//游标j从左向右移动, 如果遇到arr[j]小于arr[pivot]的, 则交换arr[j]和arr[i]的位置, 同时i++, 这样i会一直标记着从左数起, 第一个大于arr[pivot]的元素下标
	i := left
	for j := left; j < pivot; j++ { //一次遍历，只记录左侧大于的
		//由于这里存在元素交换, 所以快排是不稳定算法(因为你永远的不知道被交换的元素, 后面是不是也要相同的, 这样他们的相对顺序就可能发生变化了)
		if arr[j] > arr[pivot] {
			Swap(arr, i, j) //大于的放左边
			i++             //记录的大于标志的最前面的一个小于或者等于的下标
		}
	}
	//最后交换arr[pivot]和arr[i]的位置, 则可以确保arr[pivot]左边的元素都比自己小, 右边的元素都比自己大
	Swap(arr, i, pivot) //标志的划分后的下标
	return i
}

func findKlargest(arr []int, k int) int {
	return findKlargestgo(arr, 0, len(arr)-1, k) //查找第K大的数
}

//    1  2 3
func findKlargestgo(arr []int, left int, right int, k int) int {
	if left >= right {
		return arr[left] //只有一个元素，这个元素就是最大
	}
	query := partition(arr, left, right) //切割
	if query+1 == k {
		return arr[query] //第K大的数字
	}
	if k < query+1 {
		return findKlargestgo(arr, left, query-1, k) //递归一直操作到区间为1
	}
	//一个一个的增加
	return findKlargestgo(arr, left, query+1, k) //递归一直操作到区间为1
}

func main() {
	//           0   1  2  3   4   5  6   7 8   9  10  11  12 13 14  15
	arr := []int{11, 2, 3, 23, 33, 3, 13, 4, 15, 6, 6, 61, 6, 17, 9, 10}
	fmt.Println(findKlargest(arr, 1))
	//提取QQ中间，最大的100个，最小的100个

}
func main1x() {
	//          0 1 2 3 4 5 6 7 8 9101112
	arr := []int{11, 2, 3, 23, 33, 3, 13, 4, 15, 6, 6, 61, 6, 17, 9, 10}
	QuickSortK(arr)
	fmt.Println(arr)
	// [61 33 23 17 15 13 11 10 9 6 6 6 4 3 3 2]

	fmt.Println(arr[len(arr)-2])

	//1  61
	//2  33
	//3  23
}
