package main

import "fmt"

//70亿人
// 160   -------
//161  ------------------------
//162  ---------------------------------
//11
//222
//333
//44

func BuckerSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		num := 4

		//[[] [] [] []]
		buckets := make([][]int, num) //创造二维数组
		for i := 0; i < length; i++ {
			buckets[arr[i]-1] = append(buckets[arr[i]-1], arr[i]) //木桶计数加1
		}
		fmt.Println(buckets)
		tmppose := 0 //木桶排序
		for i := 0; i < num; i++ {
			bucketslen := len(buckets[i]) //求某一段的长度
			if bucketslen > 0 {
				copy(arr[tmppose:], buckets[i]) //拷贝数据
				tmppose += bucketslen           //重新定位
			}
		}

		return arr

	}

}

func SelectSortMax(arr []int) int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr[0] //一个元素的数组，直接返回
	} else {
		max := arr[0] //假定第一个最大
		for i := 1; i < length; i++ {
			if arr[i] > max { //任何一个比我的大的数，最大的
				max = arr[i]
			}
		}
		return max
	}

}
func SelectSort(arr []int) []int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		for i := 0; i < length-1; i++ { //只剩一个元素不需要挑选，
			min := i                          //标记索引
			for j := i + 1; j < length; j++ { //每次选出一个极小值
				if arr[min] > arr[j] {
					min = j //保存极小值的索引
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i] //数据交换
			}
			fmt.Println(arr)

		}

		return arr

	}

}

//基数排序又称为木桶排序
func BuckerSortX(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		num := length
		max := SelectSortMax(arr)     //极大值
		buckets := make([][]int, num) //创造二维数组
		index := 0                    //索引
		for i := 0; i < length; i++ {
			index = arr[i] * (num - 1) / max                //木桶的自动分配算法,index最大值为length-1
			buckets[index] = append(buckets[index], arr[i]) //木桶计数加1
		}
		fmt.Println(buckets)
		tmppose := 0 //木桶排序
		for i := 0; i < num; i++ {
			bucketslen := len(buckets[i]) //求某一段的长度
			if bucketslen > 0 {
				buckets[i] = SelectSort(buckets[i]) //木桶内部数据排序，单个数组进行排序,不仅仅是选择排序了
				copy(arr[tmppose:], buckets[i])     //拷贝数据
				tmppose += bucketslen               //定位
			}
		}

		return arr

	}

}

//1850 2019
func main() {
	arr := []int{1, 2, 3, 4, 4, 3, 2, 2, 3, 1}
	fmt.Println(arr)
	arr = BuckerSortX(arr)
	fmt.Println(arr)
}
