package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const N = 63812122

//数组性能，用二分查找2ms.不用200ms

func main() {
	allmap := make(map[string]string, N) //开辟一个map映射

	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day3\\163.txt" //路径
	sqlfile, _ := os.Open(path)                                    //打开文件
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
			allmap[lines[0]] = lines[1] //映射到map

		}

		i++

	}
	fmt.Println(i, "内存载入完成")

	for {
		fmt.Println("请输入要查询的用户名")
		var inputstr string
		fmt.Scanln(&inputstr) //用户输入

		starttime := time.Now()
		getdata, err := allmap[inputstr] //
		if err {                         //找到了，返回true
			fmt.Println(getdata, inputstr, "存在")
		} else {
			fmt.Println("不存在")
		}

		fmt.Println("本次查询用了", time.Since(starttime))

	}

}

func main1() {
	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day3\\163.txt" //路径
	sqlfile, _ := os.Open(path)                                    //打开文件
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
