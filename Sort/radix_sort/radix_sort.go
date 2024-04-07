package main

import "fmt"

// 基数排序又称桶排序，空间复杂度是数组的个数和桶数相关，个数n和桶数m
// 故空间复杂度是(n+m)
//此处利用的是 以空间换取时间

func SelectSortMaxx(arr []int) int {
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

func RadixSort(arr []int) []int {
	max := SelectSortMaxx(arr) //寻找数组的极大值
	//进行个十百千位的排序
	for bit := 1; max/bit > 0; bit *= 10 {
		arr = BitSort(arr, bit)
		fmt.Println(arr)
	}
	return arr
}

func BitSort(arr []int, bit int) []int {
	length := len(arr)            //数组长度
	bitcounts := make([]int, 10)  //统计长度0,1,2,3,4,5,6,7,8,9
	for i := 0; i < length; i++ { //个十百。。。
		num := (arr[i] / bit) % 10 //分层处理，bit=1000的，三位数不参与排序了，bit=10000的四位数不参与排序
		bitcounts[num]++           //统计余数相等个数
	}
	fmt.Println(bitcounts)
	//  0 1 2 3  4 5
	//  1 0 3 0  0  1
	//  1 1 4 4  4  5
	for i := 1; i < 10; i++ {
		bitcounts[i] += bitcounts[i-1] //叠加，计算位置  *****
	}
	fmt.Println(bitcounts)

	//重点： 这个空间复杂度是（n+m）， 即与数组的个数相关

	tmp := make([]int, 10) //开辟临时数组  和排序数的个数有关系
	for i := length - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		tmp[bitcounts[num]-1] = arr[i] //计算排序的位置,下标为零
		bitcounts[num]--
	}

	for i := 0; i < length; i++ {
		arr[i] = tmp[i] //保存数组
	}
	return arr

}

func main() {
	arr := []int{11, 91, 222, 878, 348, 7123, 4213, 6232, 5123, 111011}
	//11  222  7123
	//91
	//

	//fmt.Println(SelectSortMax(arr))
	fmt.Println(RadixSort(arr))

}
