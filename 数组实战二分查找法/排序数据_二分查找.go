package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type QQ struct {
	QQuser int
	QQpass string
}

const N = 84331445

//1297925096----xuanjiuaiao.

func QuickSortStruct(arr []QQ) []QQ {
	length := len(arr)
	if length <= 1 {
		return arr //只有一元素无需排序
	} else {
		splitdata := arr[0].QQuser //第一个数字
		low := make([]QQ, 0, 0)
		high := make([]QQ, 0, 0)
		mid := make([]QQ, 0, 0)
		mid = append(mid, arr[0]) //保存分离的数据
		//数据分为三段处理，分别是大于，等于，小于
		for i := 1; i < length; i++ {
			if arr[i].QQuser < splitdata {
				low = append(low, arr[i])
			} else if arr[i].QQuser > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSortStruct(low), QuickSortStruct(high) //递归循环
		myarr := append(append(low, mid...), high...)           //数据归并
		return myarr

	}
}
func bin_searchstruct(arr []QQ, data int) int {
	left := 0
	right := len(arr) - 1 //最下最上面
	for left < right {
		mid := (left + right) / 2
		if arr[mid].QQuser > data {
			right = mid - 1
		} else if arr[mid].QQuser < data {
			left = mid + 1 //移动
		} else {
			return mid //找到
		}
	}
	return -1
}

func main() {
	const nx = 15000002
	allstrs := make([]QQ, nx, nx) //初始化数组

	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day1数据结构\\QQ.txt"
	QQfile, _ := os.Open(path)    //打开文件
	defer QQfile.Close()          //最后关闭文件
	i := 0                        //统计一共多少行
	br := bufio.NewReader(QQfile) //读取数据
	for {
		line, _, end := br.ReadLine() //读取一行数据
		if end == io.EOF {            //文件关闭。跳出循环
			break
		}
		linestr := string(line)
		lines := strings.Split(linestr, "----")
		if len(lines) == 2 {
			allstrs[i].QQuser, _ = strconv.Atoi(lines[0]) //字符串转整数
			allstrs[i].QQpass = lines[1]
		}
		if i == 15000000 {
			break
		}
		//allstrs[i]=string(line)
		i++
	}
	fmt.Println("数据载入内存", i)
	time.Sleep(time.Second * 1)
	//排序
	starttimesort := time.Now() //时间开始
	fmt.Println("开始排序", len(allstrs))
	allstrs = QuickSortStruct(allstrs)
	fmt.Println("结束排序一共用了", time.Since(starttimesort))

	for {
		fmt.Println("请输入要查询的数据")
		var QQ int
		fmt.Scanf("%d", &QQ) //查询QQ

		starttime := time.Now() //时间开始
		//顺序查找，修改二分查找
		index := bin_searchstruct(allstrs, QQ)
		if index == -1 {
			fmt.Println("数据查找不到")
		} else {
			fmt.Println("数据找到", index, allstrs[index].QQuser, allstrs[index].QQpass)
		}

		fmt.Println("一共用了", time.Since(starttime))
		//break
	}

}

/*
数据载入内存
请输入要查询的数据
81829191
55657869 81829191 kekexili2008
一共用了 404.3606ms
请输入要查询的数据
645070801
68278495 645070801 1346883819900
一共用了 330.0641ms
请输入要查询的数据
645070807
82453294 645070807 kekexili3307
一共用了 309.1778ms
*/
