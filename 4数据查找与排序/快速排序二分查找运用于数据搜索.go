package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const N = 18333811

//const  N=10333811
//结构体定义数据
type uuu9 struct {
	user     string
	md5      string
	email    string
	password string
}

func QuickSortU(arr []uuu9) []uuu9 {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		splitdata := arr[0]          //以第一个为基准
		low := make([]uuu9, 0, 0)    //存储比我小的
		high := make([]uuu9, 0, 0)   //存储比我大的
		mid := make([]uuu9, 0, 0)    //存储与我相等
		mid = append(mid, splitdata) //加入第一个相等

		for i := 1; i < length; i++ {
			if arr[i].user < splitdata.user {
				low = append(low, arr[i])
			} else if arr[i].user > splitdata.user {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSortU(low), QuickSortU(high) //切割递归处理
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}
func bin_searchU(arr []uuu9, data string) int {
	low := 0             //最下方
	high := len(arr) - 1 //最上方

	for low <= high { //循环的终止条件
		//fmt.Println(arr[low:high])
		mid := (low + high) / 2
		//fmt.Println("mid",mid)
		if arr[mid].user > data {
			high = mid - 1
		} else if arr[mid].user < data {
			low = mid + 1
		} else {
			return mid //找到
		}
	}
	return -1

}

func main() {
	alldata := make([]uuu9, N+2, N+2) //开辟数组，1800万

	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day3\\uuu9.com.sql" //路径
	sqlfile, _ := os.Open(path)                                         //打开文件
	defer sqlfile.Close()
	i := 0                         //统计行数
	br := bufio.NewReader(sqlfile) //读取文件对象
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF { //文件结束跳出循环
			break
		}
		linestr := string(line)                //转化为字符串
		lines := strings.Split(linestr, " | ") //切割数据
		if len(lines) == 4 {                   //判断切割成功
			alldata[i].user = lines[0]
			alldata[i].md5 = lines[1]
			alldata[i].email = lines[2]
			alldata[i].password = lines[3]
		}
		if i == N {
			break
		}
		i++

	}
	fmt.Println(i, "内存载入完成")
	alldata = QuickSortU(alldata)
	fmt.Println("排序完成")

	for {
		fmt.Println("请输入要查询的用户名")
		var inputstr string
		fmt.Scanln(&inputstr) //用户输入

		starttime := time.Now()
		index := bin_searchU(alldata, inputstr)
		fmt.Println("index", index)
		if index == -1 {
			fmt.Println("找不到")
		} else {
			fmt.Println("找到", alldata[index])
		}
		fmt.Println("本次查询用了", time.Since(starttime))

	}

}
