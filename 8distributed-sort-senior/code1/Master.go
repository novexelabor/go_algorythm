package main

import (
	"fmt"
	"net"
	"os"
)

//处理错误
func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	tcpaddr, err := net.ResolveTCPAddr("tcp4", "localhost:7000")
	CheckError(err)
	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	CheckError(err)
	for {

		var inputstr string
		fmt.Scanln(&inputstr)        //输入
		conn.Write([]byte(inputstr)) //传输
		buf := make([]byte, 1024)    //接收数据
		n, _ := conn.Read(buf)
		fmt.Println(string(buf[:n]))
	}

}
