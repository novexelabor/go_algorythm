package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const N = 84331445

//结构体定义数据
type QQ struct {
	QQuser   int
	password string
}

func SortForMergeQQ(arr []QQ, left int, right int) {
	for i := left; i <= right; i++ {
		temp := arr[i] //备份数据
		var j int
		for j = i; j > left && arr[j-1].QQuser > temp.QQuser; j-- { //定位
			arr[j] = arr[j-1] //数据往后移动
		}
		arr[j] = temp //插入
	}
}

func swapQQ(arr []QQ, i int, j int) { //数据交换
	arr[i], arr[j] = arr[j], arr[i]
}

//3       1 8 2 3  9 7
//2       1 3   3  897

//递归快速排序
func QuickSortXQQ(arr []QQ, left int, right int) {
	if right-left < 15 { //数组剩下3个数，直接插入排序
		SortForMergeQQ(arr, left, right)
	} else {
		//随机找一个数字，放在第一个位置
		swapQQ(arr, left, rand.Int()%(right-left+1)+left)
		vdata := arr[left] //坐标数组，比我小，左边，比我大右边
		lt := left         // arr [left+1,  lt] <vata
		gt := right + 1    //arr[gt...  right] >vata
		i := left + 1      //arr[lt+1,...i] ==vdata
		for i < gt {
			if arr[i].QQuser < vdata.QQuser {
				swapQQ(arr, i, lt+1) //移动到小于的地方
				lt++                 //前进循环
				i++

			} else if arr[i].QQuser > vdata.QQuser {
				swapQQ(arr, i, gt-1) //移动到大于的地方
				gt--

			} else {
				i++
			}
		}
		swapQQ(arr, left, lt)         //交换头部位置
		QuickSortXQQ(arr, left, lt-1) //递归处理小于那一段
		QuickSortXQQ(arr, gt, right)  //递归处理大于那一段

	}

}

//快速排序核心程序
func QuicksortPlusQQ(arr []QQ) {
	QuickSortXQQ(arr, 0, len(arr)-1)
}

func bin_searchQQX(arr []QQ, data int) int {
	low := 0             //最下方
	high := len(arr) - 1 //最上方

	for low <= high { //循环的终止条件
		//fmt.Println(arr[low:high])
		mid := (low + high) / 2
		//fmt.Println("mid",mid)
		if arr[mid].QQuser > data {
			high = mid - 1
		} else if arr[mid].QQuser < data {
			low = mid + 1
		} else {
			return mid //找到
		}
	}
	return -1

}
func main() {
	alldata := make([]QQ, N, N) //开辟数组，1800万

	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day3\\QQ.txt" //路径
	sqlfile, _ := os.Open(path)                                   //打开文件
	defer sqlfile.Close()

	i := 0                         //统计行数
	br := bufio.NewReader(sqlfile) //读取文件对象
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF { //文件结束跳出循环
			break
		}
		linestr := string(line)                 //转化为字符串
		lines := strings.Split(linestr, "----") //切割数据
		if len(lines) == 2 {                    //判断切割成功
			alldata[i].QQuser, _ = strconv.Atoi(lines[0]) //字符串转化为整数
			alldata[i].password = lines[1]

		}

		i++

	}
	fmt.Println(i, "内存载入完成")
	starttimex := time.Now()
	QuicksortPlusQQ(alldata)
	fmt.Println("排序完成", time.Since(starttimex))
	for {
		fmt.Println("请输入要查询的用户名")
		var inputstr int
		fmt.Scanf("%d", &inputstr) //用户输入

		starttime := time.Now()
		index := bin_searchQQX(alldata, inputstr)
		fmt.Println("index", index)
		if index == -1 {
			fmt.Println("找不到")
		} else {
			fmt.Println("找到", alldata[index])
		}

		fmt.Println("本次查询用了", time.Since(starttime))

	}

}

func main1x() {
	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day3\\QQ.txt" //路径
	sqlfile, _ := os.Open(path)                                   //打开文件
	defer sqlfile.Close()

	i := 0                         //统计行数
	br := bufio.NewReader(sqlfile) //读取文件对象
	for {
		_, _, end := br.ReadLine()
		if end == io.EOF { //文件结束跳出循环
			break
		}

		i++

	}
	fmt.Println(i, "内存载入完成")
}
