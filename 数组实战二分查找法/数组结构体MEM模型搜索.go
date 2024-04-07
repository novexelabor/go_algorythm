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

func main() {

	allstrs := make([]QQ, N, N) //初始化数组

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

		//allstrs[i]=string(line)
		i++
	}
	fmt.Println("数据载入内存")
	time.Sleep(time.Second * 10)

	for {
		fmt.Println("请输入要查询的数据")
		var QQ int
		fmt.Scanf("%d", &QQ) //查询QQ

		starttime := time.Now() //时间开始
		for j := 0; j < N; j++ {
			if allstrs[j].QQuser == QQ {
				fmt.Println(j, allstrs[j].QQuser, allstrs[j].QQpass) //根据数据查询QQ，密码
			}

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
