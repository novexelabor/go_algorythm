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

const N = 84331445

//1297925096----xuanjiuaiao.

func main() {

	allstrs := make(map[int]string, N) //初始化map映射

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
			QQuser, _ := strconv.Atoi(lines[0]) //字符串转整数
			QQpass := lines[1]
			allstrs[QQuser] = QQpass //映射到map中
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
		QQpass, err := allstrs[QQ]
		if err {
			fmt.Println(QQ, QQpass, "存在")
		} else {
			fmt.Println("找不到")
		}

		fmt.Println("一共用了", time.Since(starttime))
		//break
	}

}

/*
数据载入内存
请输入要查询的数据
81829191
81829191 kekexili2008 存在
一共用了 0s
请输入要查询的数据
645070801
645070801 1346883819900 存在
一共用了 0s
请输入要查询的数据
645070807
645070807 kekexili3307 存在
一共用了 0s
请输入要查询的数据
645012345
找不到
一共用了 0s
请输入要查询的数据
78025078
找不到
一共用了 0s
请输入要查询的数据
77025077
找不到
一共用了 0s
请输入要查询的数据
0 520520520520 存在
一共用了 0s
请输入要查询的数据
*/
