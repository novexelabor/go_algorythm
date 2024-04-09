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

const N = 84331445

//结构体定义数据
type QQ struct {
	QQuser   int
	password string
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

	for {
		fmt.Println("请输入要查询的用户名")
		var inputstr int
		fmt.Scanf("%d", &inputstr) //用户输入

		starttime := time.Now()
		for i := 0; i < N; i++ {
			if alldata[i].QQuser == inputstr {
				fmt.Println(alldata[i])
			}
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
