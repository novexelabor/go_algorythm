package main

import "fmt"

const M = 10
const N = 10

//1代表人，2代表障碍，0空地

var data [M][N]int = [M][N]int{
	{1, 0, 2, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 2, 0, 0, 2, 0, 0, 0},
	{0, 0, 0, 0, 2, 0, 0, 2, 2, 0},
	{0, 0, 0, 0, 2, 0, 2, 0, 0, 0},
	{0, 0, 0, 0, 2, 0, 0, 0, 0, 2},
	{2, 2, 2, 0, 0, 2, 0, 2, 0, 0},
	{0, 0, 0, 0, 0, 2, 0, 2, 0, 2},
	{0, 0, 0, 0, 0, 0, 0, 2, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 2, 2, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 2, 0},
}

var ipos, jpos int = 0, 0 //定义初始位置
var cangoout bool = false //假定走不出迷宫

func show(arr [M][N]int) {
	//fmt.Println("-----------------------------------------")
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%4d", arr[i][j])
		}
		fmt.Println("")
	}

}

//寻路，i,j =0,0
func FindOut(datax [M][N]int, i int, j int) bool {
	datax[i][j] = 3 //避免回头路
	if i == M-1 && j == N-1 {
		cangoout = true //代表可以走出来
		data = datax
		fmt.Println("迷宫走出来")

	} else {
		if j+1 <= N-1 && datax[i][j+1] < 2 && cangoout != true {
			FindOut(datax, i, j+1)
		}
		if j-1 >= 0 && datax[i][j-1] < 2 && cangoout != true {
			FindOut(datax, i, j-1)
		}
		if i+1 <= M-1 && datax[i+1][j] < 2 && cangoout != true {
			FindOut(datax, i+1, j)
		}
		if i-1 >= 0 && datax[i-1][j] < 2 && cangoout != true {
			FindOut(datax, i-1, j)
		}
	}

	return cangoout

}

func main() {
	show(data)
	isok := FindOut(data, 0, 0)
	if isok {
		fmt.Println("可以走出")
		show(data)
	}
}
