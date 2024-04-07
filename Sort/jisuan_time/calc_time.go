package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	//"strings"
	"time"
)

func QuickSortString(arr []string) []string {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		splitdata := arr[0]          //以第一个为基准
		low := make([]string, 0, 0)  //存储比我小的
		high := make([]string, 0, 0) //存储比我大的
		mid := make([]string, 0, 0)  //存储与我相等
		mid = append(mid, splitdata) //加入第一个相等

		for i := 1; i < length; i++ {
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSortString(low), QuickSortString(high) //切割递归处理
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}
func InsertSortString(arr []string) []string {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		for i := 1; i < length; i++ { //跳过第一个
			backup := arr[i] //备份插入的数据
			j := i - 1       //上一个位置循环找到位置插入
			for j >= 0 && backup < arr[j] {
				arr[j+1] = arr[j] //从前往后移动
				j--
			}
			arr[j+1] = backup //插入
			//fmt.Println(arr)
		}

		return arr

	}
}

func main() {
	t1 := time.Now()
	const N = 6428632 //需要开辟的内存
	allstrs := make([]string, N)
	fmt.Println(len(allstrs)) //开辟数组存储数据

	fi, err := os.Open("/Volumes/WILLING/Go语言数据结构与算法/2.数组排序/CSDNpass.txt")
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fi.Close() //延迟关闭文件

	br := bufio.NewReader(fi)
	i := 0
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break //t跳出循环
		}
		//fmt.Println(string(line))
		linestr := string(line) //读取，转化为字符串
		allstrs[i] = linestr
		i++

	}
	fmt.Println(i, "ok,读取完成")
	//time.Sleep(1*time.Second)
	//allstrs=QuickSortString(allstrs)
	allstrs = InsertSortString(allstrs)
	fmt.Println("ok,排序完成")

	used := time.Since(t1)
	fmt.Println("使用的时间")
	fmt.Println(used)

	path := "/Users/renshanwan/Documents/wanlin/go_algorythm/Sort/jisuan_time/CSDNSortmail_insert.txt"
	savefile, _ := os.Create(path)
	defer savefile.Close()
	save := bufio.NewWriter(savefile) //对象用于写入
	for i := 0; i < len(allstrs); i++ {
		fmt.Fprintln(save, allstrs[i])
	}
	save.Flush()

}
