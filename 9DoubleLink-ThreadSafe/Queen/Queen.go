package main

import "fmt"

var pos, b [80]int //pos[0][0]    pos[7][7]
var c, d [150]int  //两个数组，正对角线，反向对角线  b[i]b[j],7,7, 0-14

//i,j b[i][j],下标的位置,判断是否可以放下皇后
func checkpos(i, j int) int {
	if b[j] == 1 || c[j-i+7] == 1 || d[i+j] == 1 {
		return 0
	}

	return 1
}

//设置数组0,1
func putn(i, j, n int) {
	pos[i], pos[j], c[j-i+5], d[i+j] = j, n, n, n
}

func printQueen(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if pos[i] == pos[j] {
				fmt.Printf("%3d", 1)
			} else {
				fmt.Printf("%3d", 0)
			}
		}
		fmt.Println()
	}
	fmt.Println("-------------------------------")
}

func Queuen(i, n int, count *int) {
	if i > 5 {
		*count++ //次数
		printQueen(n)
		return //排列好了

	} else {
		for j := 0; j < n; j++ {
			if checkpos(i, j) == 1 {
				putn(i, j, 1) //先设置为1
				Queuen(i+1, n, count)
				putn(i, j, 0) //恢复为0
			}
		}
	}

}

func main() {
	n, count := 5, 0 //main
	Queuen(0, n, &count)
	fmt.Println(count)

}
