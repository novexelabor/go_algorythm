package main

import "fmt"

const N = 5
const W = 6

var weight []int = []int{3, 1, 1, 1, 1} //3,2   //2,1,2
var val []int = []int{9, 3, 2, 6, 7}    //7    //7
var dp [N + 1][W + 1]int                //保存中间结果
var record [N][W + 1]int                //保存中间结果
// [0 0 0 0 0 0]
// [0 0 0 4 4 4]
// [0 3 3 4 7 7]
// [0 3 5 5 7 9]
// [0 4 7 9 9 11]
//3 4
//3 4
func solve2() int {
	length := len(val)
	for i := 0; i < length; i++ {
		for j := weight[i]; j <= W; j++ {
			dp[i+1][j] = max(dp[i][j], dp[i][j-weight[i]]+val[i])
		}
	}
	fmt.Println(dp)
	return dp[N][W]
}

func init() { //初始化背包结果
	for i := 0; i < N; i++ {
		for j := 0; j <= W; j++ {
			record[i][j] = -1
		}
	}
}

//两个数之间最大值
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//i是物品的下标，total总量
func solve(i int, total int) int { //record[][]
	result := 0 //结果
	if i >= N {
		return result
	}
	if record[i][total] != -1 { //如果数据已经记录，直接返回
		return record[i][total]
	}
	if weight[i] > total {
		record[i][total] = solve(i+1, total) //当前物品大于总量，跳出，计算下一个
	} else {
		//递归求最大值,退出一个再加一个
		result = max(solve(i+1, total), solve(i+1, total-weight[i])+val[i])
	}
	record[i][total] = result //记录中间结果
	return record[i][total]
}

func main() {
	fmt.Println(solve(0, W))
	fmt.Println(solve2())
}
