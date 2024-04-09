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

//结构体定义数据
type uuu9 struct {
	user     string
	md5      string
	email    string
	password string
}

func main() {
	alldata := make([]uuu9, N, N) //开辟数组，1800万

	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day3\\uuu9.com.sql" //路径
	sqlfile, _ := os.Open(path)                                         //打开文件
	defer sqlfile.Close()                                               //延迟关闭文件
	i := 0                                                              //统计行数
	br := bufio.NewReader(sqlfile)                                      //读取文件对象
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

		i++

	}
	fmt.Println(i, "内存载入完成")

	for {
		fmt.Println("请输入要查询的用户名")
		var inputstr string
		fmt.Scanln(&inputstr) //用户输入

		starttime := time.Now()
		for i := 0; i < N; i++ {
			if alldata[i].user == inputstr {
				fmt.Println(alldata[i])
			}
		}
		fmt.Println("本次查询用了", time.Since(starttime))

	}

}

func main1() {
	//统计行数
	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day3\\uuu9.com.sql" //路径
	sqlfile, _ := os.Open(path)                                         //打开文件
	defer sqlfile.Close()                                               //延迟关闭文件

	i := 0                         //统计行数
	br := bufio.NewReader(sqlfile) //读取文件对象
	for {
		_, _, end := br.ReadLine()
		if end == io.EOF { //文件结束跳出循环
			break
		}
		i++

	}
	fmt.Println(i)

}
