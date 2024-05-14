package main

import "fmt"

const Num int = 3 //控制皇后问题
var count int = 1 //次数
var queens [Num][Num]int

func show() {
	fmt.Printf("第%d种解法\n", count)
	for i := 0; i < Num; i++ {
		for j := 0; j < Num; j++ {
			if queens[i][j] == 1 {
				fmt.Printf("%s", "■")
			} else {
				fmt.Printf("%s", "□")
			}
		}
		fmt.Println()
	}
}

func setQueen(row, col int) bool {
	if row == 0 { //第一个放入
		queens[row][col] = 1 //设置为1
		return true
	}
	for i := 0; i < Num; i++ {
		if queens[row][i] == 1 { //列有一个为1，无法放下
			return false
		}
	}
	for i := 0; i < Num; i++ {
		if queens[i][col] == 1 { //行有一个为1，无法放下
			return false
		}
	}
	for i, j := row, col; i < Num && j < Num; i, j = i+1, j+1 {
		if queens[i][j] == 1 { //对角线无法放下，
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if queens[i][j] == 1 { //对角线无法放下，
			return false
		}
	}
	for i, j := row, col; i < Num && j >= 0; i, j = i+1, j-1 {
		if queens[i][j] == 1 { //对角线无法放下，
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < Num; i, j = i-1, j+1 {
		if queens[i][j] == 1 { //对角线无法放下，
			return false
		}
	}
	queens[row][col] = 1 //设置为1
	return true
}

func solveQueen(row int) { //一直递归下去是成功的，回退的重置为0,n*n的棋盘
	if row == Num {
		show()
		count++
		return
	}
	for i := 0; i < Num; i++ {
		if setQueen(row, i) {
			solveQueen(row + 1)
		}
		queens[row][i] = 0 //回退solveQueen(row+1) //循环
	}
}

func main() {
	solveQueen(0)
}
