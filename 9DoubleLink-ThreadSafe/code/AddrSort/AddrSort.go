package main

import "fmt"

func QuickSort(arr []int, addr []int) []int {
	length := len(addr) //数组长度
	if length <= 1 {
		return addr //一个元素的数组，直接返回
	} else {

		splitaddr := addr[0]                 //以第一个为基准
		lowaddr := make([]int, 0)            //存储比我小的
		highaddr := make([]int, 0)           //存储比我大的
		midaddr := make([]int, 0)            //存储与我相等
		midaddr = append(midaddr, splitaddr) //加入第一个相等

		split := arr[0]          //以第一个为基准
		low := make([]int, 0)    //存储比我小的
		high := make([]int, 0)   //存储比我大的
		mid := make([]int, 0)    //存储与我相等
		mid = append(mid, split) //加入第一个相等

		for i := 1; i < length; i++ {
			if arr[i] < split {
				lowaddr = append(lowaddr, addr[i])
				low = append(low, arr[i])
			} else if arr[i] > split {
				highaddr = append(highaddr, addr[i])
				high = append(high, arr[i])
			} else {

				midaddr = append(midaddr, splitaddr)
				mid = append(mid, split)
			}
		}
		lowaddr, highaddr = QuickSort(low, lowaddr), QuickSort(high, highaddr) //切割递归处理
		myarr := append(append(lowaddr, midaddr...), highaddr...)
		return myarr //返回地址
	}
}

func main() {

	arr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5, 10}
	arraddr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	fmt.Println(arraddr)
	fmt.Println("-------------")
	myarr := QuickSort(arr, arraddr)
	fmt.Println(arr)
	fmt.Println(myarr)
	fmt.Println("-------------")
	for _, v := range myarr {
		fmt.Println(arr[v])
	}

}
