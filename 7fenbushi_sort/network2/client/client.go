package main

import (
	"fmt"
	"net"
)

func main() {
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8848")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpaddr) //链接
	if err != nil {
		panic(err)
	}

	for {
		var inpustr string
		fmt.Scanln(&inpustr)
		conn.Write([]byte(inpustr))
		buf := make([]byte, 1024)

		n, _ := conn.Read(buf) //读取数据
		fmt.Println(string(buf[:n]))
	}

}

//0helloworld //数据
//1calc  //命令
