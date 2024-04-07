package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	starttime := time.Now() //时间开始

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
		//fmt.Println(string(line))//显示数据
		linestr := string(line)
		if strings.Contains(linestr, "yincheng") { //字符串搜索
			fmt.Println(linestr)
		}
		i++
	}

	fmt.Println("一共用了", time.Since(starttime))
	fmt.Println("数据一共这么多行", i)

}

/*
.一共用了 41.0401598s
数据一共这么多行 84331445
*/
