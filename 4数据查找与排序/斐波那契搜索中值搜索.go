package main

import "fmt"

func bin_search(arr []int, data int) int {
	low := 0             //最下方
	high := len(arr) - 1 //最上方

	i := 0
	for low <= high { //循环的终止条件
		i++
		fmt.Println("第N次", i)
		mid := (low + high) / 2
		if arr[mid] > data {
			high = mid - 1
		} else if arr[mid] < data {
			low = mid + 1
		} else {
			return mid //找到
		}
	}
	return -1

}

func bin_searchMid(arr []int, data int) int {
	low := 0             //最下方
	high := len(arr) - 1 //最上方

	i := 0
	for low <= high { //循环的终止条件
		i++
		fmt.Println("第N次", i)

		leftv := float64(data - arr[low])     //大段
		allv := float64(arr[high] - arr[low]) //整段
		diff := float64(high - low)
		mid := int(float64(low) + leftv/allv*diff) //计算中间值
		//mid:=(low+high)/2

		if mid < 0 || mid >= len(arr) {
			return -1
		}

		if arr[mid] > data {
			high = mid - 1
		} else if arr[mid] < data {
			low = mid + 1
		} else {
			return mid //找到
		}
	}
	return -1

}

func main1111() {
	arr := make([]int, 1000, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}
	fmt.Println(arr)

	index := bin_searchMid(arr, 12)
	fmt.Println("index", index)
	if index == -1 {
		fmt.Println("找不到")
	} else {
		fmt.Println("找到")
	}
}

func makeFabArray(arr []int) []int {
	length := len(arr) //数组长度
	flblen := 2
	first, secord, third := 1, 2, 3

	fmt.Println(length) //打印长度

	for third < length { //找出最接近的菲薄拿起
		third, first, secord = first+secord, secord, third //叠加计算菲薄纳妾
		flblen++
		fmt.Println(third, flblen) //斐波那契数组定制开辟多少个数
	}
	fb := make([]int, flblen) //开辟数组
	fb[0] = 1
	fb[1] = 1
	for i := 2; i < flblen; i++ { //叠加计算
		fb[i] = fb[i-1] + fb[i-2]
	}
	fmt.Println(flblen) //打印该长度
	return fb
}

func fab_search(arr []int, val int) int {
	length := len(arr)          //数组长度
	fabArr := makeFabArray(arr) //定制匹配的斐波那契额数组
	fmt.Println(fabArr)
	filllength := fabArr[len(fabArr)-1] //填充长度

	fillArr := make([]int, filllength) //填充的数组
	for i, v := range arr {
		fillArr[i] = v
	}

	lastdata := arr[length-1] //填充最后一个大数
	for i := length; i < filllength; i++ {
		fillArr[i] = lastdata //填充数据
	}
	fmt.Println(fillArr, length) //填充以后的数组

	left, mid, right := 0, 0, length //类似二分查找
	kindex := len(fabArr) - 1        //游标

	for left <= right {
		//fabArr[kindex]是超过length了
		mid = left + fabArr[kindex-1] - 1 //菲薄纳妾切割
		if val < fillArr[mid] {
			right = mid - 1 //类似二分查找
			kindex--

		} else if val > fillArr[mid] {
			left = mid + 1
			kindex -= 2 //减2

		} else {
			if mid > right { //filllength是大于length的
				return right //越界
			} else {
				return mid
			}
		}

	}

	return -1
}

func main() {
	arr := make([]int, 1000, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	//fmt.Println(makeFabArray([]int{1,2,3,4,5,6,7,8,9,10}))
	//fmt.Println(fab_search(arr,12 ))

	fmt.Println(makeFabArray(arr))

	index := fab_search(arr, 113)
	fmt.Println("index", index)
	if index == -1 {
		fmt.Println("找不到")
	} else {
		fmt.Println("找到")
	}
}
