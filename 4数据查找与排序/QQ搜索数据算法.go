package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"strconv"
)

const N = 14331445

//结构体定义数据
type QQ struct {
	QQuser   int
	password string
}

func QuickSortQQ(arr []QQ) []QQ {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		splitdata := arr[0]          //以第一个为基准
		low := make([]QQ, 0, 0)      //存储比我小的
		high := make([]QQ, 0, 0)     //存储比我大的
		mid := make([]QQ, 0, 0)      //存储与我相等
		mid = append(mid, splitdata) //加入第一个相等

		for i := 1; i < length; i++ {
			if arr[i].QQuser < splitdata.QQuser {
				low = append(low, arr[i])
			} else if arr[i].QQuser > splitdata.QQuser {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSortQQ(low), QuickSortQQ(high) //切割递归处理
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}

func bin_searchQQ(arr []QQ, data int) int {
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
	alldata := make([]QQ, N+2, N+2) //开辟数组，1800万

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
		if i == N {
			break
		}
		i++

	}
	fmt.Println(i, "内存载入完成")
	alldata = QuickSortQQ(alldata)
	fmt.Println("排序完成")
	for {
		fmt.Println("请输入要查询的用户名")
		var inputstr int
		fmt.Scanf("%d", &inputstr) //用户输入

		starttime := time.Now()
		index := bin_searchQQ(alldata, inputstr)
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
