package main

import "fmt"

var pos, b [80]int //10*8// pos[i]=j,第i行皇后在j位置   b[i],第i行y有没有位置
var c, d [150]int  // c[i+j]//i -> 0--7   j --07  i+j 0-14
// d[j-i+7]反对角线
//填充数组0，1
func putn(i, j, n int) {

	//棋盘正对角线和斜对角线下标计算所的数值是一样的j+i;j-i+(n-1)
	//  pos[i] 记录第几行，b[]表示列，c[]表示 \ 位置，d[]表示 / 位置
	pos[i], b[j], c[j-i+7], d[i+j] = j, n, n, n
}

//检查皇后
func checkPos(i, j int) int {
	if b[j] == 1 || c[j-i+7] == 1 || d[i+j] == 1 {
		return 0
	}
	return 1
}
func showQ(n int) {
	fmt.Println("----------------------------")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if pos[i] == j { //第一行就存储第几列，记录位置
				fmt.Printf("%s", "■")
			} else {
				fmt.Printf("%s", "□")
			}
		}
		fmt.Println()
	}
}
func Queen(i, n int, count *int) {
	if i > 7 {
		*count++
		showQ(n)
		return //完成
	} else {
		for j := 0; j < n; j++ {
			if checkPos(i, j) == 1 {
				putn(i, j, 1)
				Queen(i+1, n, count)
				putn(i, j, 0)
			}
		}
	}

}

func main() {
	n, count := 8, 0
	Queen(0, n, &count)
	fmt.Println(count)

}
