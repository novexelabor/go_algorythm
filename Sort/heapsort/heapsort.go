package main

import "fmt"

// 堆排序： 子：2*i+1 ， 2*i+2
//         父：(i-1)/2
// 		i代表的是深度
// 		下标是从0开始的

//
//从数组的前一半开始遍历所有有孩子的节点
//从父节点往孩子节点判断，逐层往上

func HeapSortMax(arr []int, length int) []int {
	//length是数组的长度
	if length <= 1 {
		return arr
	} else {
		depth := length/2 - 1         //下标从零开始，求解深度，完全二叉树
		for i := depth; i >= 0; i-- { //循环遍历所有的三节点
			topmax := i //从下标零开始
			leftchild := 2*i + 1
			rightchild := 2*i + 2
			if leftchild <= length-1 && arr[leftchild] > arr[topmax] {
				topmax = leftchild
			}
			if rightchild <= length-1 && arr[rightchild] > arr[topmax] {
				topmax = rightchild
			}
			if topmax != i {
				arr[i], arr[topmax] = arr[topmax], arr[i] //交换数据
			}
		}
		return arr
	}
}

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastmax := length - i
		HeapSortMax(arr, lastmax) //一次循环遍历，arr[0]是最大值
		if i < length {
			//每次遍历确定的最大值放到最后
			arr[0], arr[lastmax-1] = arr[lastmax-1], arr[0]
		}
	}
	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(HeapSort(arr))
}
