package main

import "fmt"

//10亿人的身高，0-300 可以使用计数排序了

func SelectSortMaxy(arr []int) int {
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
func Countsort(arr []int) []int {
	max := SelectSortMaxy(arr) //寻找最大值

	sortedarr := make([]int, len(arr)) //排序之后存储

	countsarr := make([]int, max+1) //统计次数  用最大值申请的空间

	for _, v := range arr {
		countsarr[v]++
	}
	fmt.Println("第一次统计次数", countsarr) //统计次数

	//计数排序重要的一步就是叠加记录位置
	for i := 1; i <= max; i++ {
		countsarr[i] += countsarr[i-1] //叠加 记录位置
	}
	fmt.Println("次数叠加", countsarr) //统计次数

	for _, v := range arr {
		sortedarr[countsarr[v]-1] = v //展开数据
		countsarr[v]--                //递减

		fmt.Println("zkcount", countsarr)
		fmt.Println("zk", sortedarr)
	}
	return sortedarr

}

//1 2 3 4 5
//2 3 2 2 1

//2 5 7 9 10

// 1 2 3 4 5
// 11  222   33 44   15
//11 222 33 44 5

func main() {
	arr := []int{1, 2, 3, 4, 4, 3, 2, 1, 2, 5, 5, 3, 4, 3, 2, 1}

	fmt.Println(Countsort(arr))

}
