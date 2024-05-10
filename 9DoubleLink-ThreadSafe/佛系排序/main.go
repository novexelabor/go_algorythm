package main

import (
	"fmt"
	"math/rand"
	"time"
)

//判断数组是否从小到大
func isOrder(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			return false
		}
	}
	return true
}

//洗牌算法
func randList(list []int) {
	data := make([]int, len(list))   //新建一个数组
	copy(data, list)                 //拷贝数组
	rand.Seed(time.Now().UnixNano()) //定义随机数种子，用时间戳表示随机种子数值不一样，随机数也就不一样的
	index := rand.Perm(len(list))    //随机选择一个切片,参数n,index的随机分布
	fmt.Println(index)
	for i, k := range index {
		list[i] = data[k]
	}
	fmt.Println("list", list) //打印list

}

func main1() {
	list := []int{1, 9, 2, 8, 3, 7, 4, 5}
	randList(list)
}

func main() {
	list := []int{1, 9, 2, 8, 3, 7, 4, 5}
	fmt.Println(list)
	count := 0
	for {
		if isOrder(list) {
			fmt.Println("排序完成", list)
			break
		} else {
			randList(list)
			count++
		}
	}
	fmt.Println(count)
}
