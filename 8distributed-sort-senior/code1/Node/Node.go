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

func MsgHandler(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close() //处理文件关闭
	for {
		n, err := conn.Read(buf) //读取数据
		if err != nil {
			fmt.Println("client关闭", conn.RemoteAddr())
			return
		}
		fmt.Println("client send", string(buf))
		conn.Write([]byte("收到" + string(buf[:n])))
	}
}

func main() {
	server, err := net.Listen("tcp", "localhost:7000") //创建服务器
	CheckError(err)                                    //处理错误
	defer server.Close()                               //延迟关闭服务器
	for {
		new_conn, err := server.Accept() //接收数据信息
		CheckError(err)                  //处理错误
		go MsgHandler(new_conn)          //并发处理链接信息
	}

}
