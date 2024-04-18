package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"
)

//判断30秒内有没有产生通信
//超过30秒退出
func HeartBeat(conn net.Conn, heartchan chan byte, timeout int) {
	fmt.Println(" HeartBeat")
	i := 0
	for i < 15 { //打印多少个数据
		select {
		case hc := <-heartchan:
			fmt.Println(string(hc))
			//log.Println("heartchan", string(hc))
			conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		case <-time.After(time.Second * 30):
			fmt.Println("time out", conn.RemoteAddr())
			log.Println("time out", conn.RemoteAddr()) //客户端超时
			conn.Close()
		}

	}
}

//处理心跳的channel
func HeartChanHandler(n []byte, beatch chan byte) {
	fmt.Println(" HeartChanHandler", len(n))
	for _, v := range n {
		beatch <- v
	}
	close(beatch) //关闭管道
}

func MsgHandler(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn close")
			return
		}
		//clientip:=conn.RemoteAddr()//远程地址
		msg := buf[1:n]
		if n != 0 {

			if string(buf[0:1]) == "0" {
				fmt.Println("client  data", string(buf[1:n]))

				conn.Write([]byte("收到数据:" + string(buf[1:n]) + "\n"))
			} else {
				fmt.Println("client  cmd", string(buf[1:n]))
				cmd := exec.Command(string(buf[1:n])) //执行命令
				cmd.Run()
				conn.Write([]byte("收到命令:" + string(buf[1:n]) + "\n"))
			}

		}
		fmt.Println("-----------------")
		beatch := make(chan byte)
		go HeartBeat(conn, beatch, 30)
		go HeartChanHandler(msg, beatch)

	}

}
func main() {
	server_listener, err := net.Listen("tcp", "127.0.0.1:8848")
	if err != nil {
		panic(err) //处理错误
	}
	defer server_listener.Close() //延迟关闭
	for {
		new_conn, err := server_listener.Accept() //接收消息
		if err != nil {
			panic(err) //处理错误
		}
		go MsgHandler(new_conn) //处理客户端消息

	}

}
