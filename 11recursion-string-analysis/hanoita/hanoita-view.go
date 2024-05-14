package main

import "fmt"

const N = 10

var arr [3][N]int = [3][N]int{{0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0}}

func show() {
	fmt.Printf("%5s%5s%5s\n", "A", "B", "C")
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%5d", arr[j][i])
		}
		fmt.Println()
	}
}

func datainit(n int) {
	for i := 0; i < n; i++ {
		arr[0][N-1-i] = n - i
	}
}
func move(X, Y string) {
	var m int = int(X[0]) - 65 //转化数组0，1，2方面计算
	var n int = int(Y[0]) - 65
	//fmt.Println("m=",m)
	var imove int = -1 //imove保存第一个不等于0的索引
	for i := 1; i < 10; i++ {
		if arr[m][i] != 0 {
			imove = i
			break
		}
	}
	//fmt.Println("imove",arr[m][imove+2])
	var jmove int
	if arr[n][N-1] == 0 {
		jmove = N - 1
	} else {
		jmove = N
		for i := 0; i < 10; i++ {
			if arr[n][i] != 0 {
				jmove = i
				break
			}
		}
		jmove -= 1 //找到不等于0的上一个空位置
	}
	//交换数据
	arr[m][imove], arr[n][jmove] = arr[n][jmove], arr[m][imove]
}

func HanIO(n int, A, B, C string) {
	if n < 1 {
		return
	}
	if n == 1 {
		fmt.Printf("%s->%s\n", A, C)
		move(A, C)
		show()
	} else {
		//递归，移动的简单逻辑，剩下的交给递归完成
		HanIO(n-1, A, C, B)
		fmt.Printf("%s->%s\n", A, C)
		move(A, C)
		show()
		HanIO(n-1, B, A, C)
	}

}

func main() {
	datainit(5)
	HanIO(5, "A", "B", "C")
}

func main1z() {
	//HanIO(3,"A","B","C")
	datainit(3)
	show()
	move("A", "C")
	move("A", "B")
	move("C", "B")
	show()
}
