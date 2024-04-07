package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//文件的读写

func main() {
	fi, err := os.Open("/Volumes/WILLING/Go语言数据结构与算法/2.数组排序/CSDN-中文IT社区-600万.sql")
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fi.Close() //延迟关闭文件

	path := "/Users/renshanwan/Documents/wanlin/go_algorythm/Sort/csdn/csdnmail.txt"
	savefile, _ := os.Create(path)
	defer savefile.Close()
	save := bufio.NewWriter(savefile) //对象用于写入

	br := bufio.NewReader(fi)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break //t跳出循环
		}
		//fmt.Println(string(line))
		linestr := string(line)                 //读取，转化为字符串
		mystrs := strings.Split(linestr, " # ") //字符串切割
		//fmt.Println(mystrs[1])
		fmt.Fprintln(save, mystrs[2])

	}
	save.Flush() //刷新

}
