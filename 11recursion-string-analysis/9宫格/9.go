package main

import "fmt"

var pos [9]int
var subnum []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var zhinum []int = []int{1, 2, 3, 5, 7, 11, 13, 17, 19}
var down, up int = 0, 9

//判断两个数之和是否质数
func iszhi(n int) bool {
	for i := 0; i < 9; i++ {
		if n == zhinum[i] {
			return true
		}
	}
	return false
}

//var pos[9]int
//0  1  2
//3  4  5
//6  7  8
func check(i, n int) bool {
	//纵
	if i-3 >= 0 { //跳过0，1，2
		if iszhi(pos[i]+pos[i-3]) == false {
			return false
		}
	}
	//横
	if i%3 != 0 { //跳过0 3  6
		if iszhi(pos[i]+pos[i-1]) == false {
			return false
		}
	}
	return true
}

func fillbox(i, n, r int, count *int) {
	if i == n {
		(*count)++
		fmt.Printf("--------------%d\n", *count)
		for i := 0; i < r; i++ {
			for j := 0; j < r; j++ {
				fmt.Printf("%3d", pos[i*r+j])
			}
			fmt.Println()
		}

		return
	} else {
		for j := down; j <= up; j++ {
			//放入
			pos[i] = subnum[j] //模拟放入
			if subnum[j] != -1 && check(i, n) {
				subnum[j] = -1
				fillbox(i+1, n, r, count) //回溯递归
				subnum[j] = pos[i]
			}
		}
	}

}

func main() {
	count := 0
	fillbox(0, 9, 3, &count) //调用
	fmt.Println(count)
}
