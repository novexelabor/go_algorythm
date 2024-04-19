package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"

	"errors"
)

//处理错误
func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func IntToBytes(n int) []byte {
	data := int64(n)                               //64位电脑8字节
	bytesBuf := bytes.NewBuffer([]byte{})          //空的字节数组
	binary.Write(bytesBuf, binary.BigEndian, data) //写入数据
	return bytesBuf.Bytes()                        //返回字节
}
func BytesToInt(b []byte) int {
	bytesBuf := bytes.NewBuffer(b) //空的字节数组
	var data int64
	binary.Read(bytesBuf, binary.BigEndian, &data) //解码
	return int(data)
}
func sendArray(arr []int, conn net.Conn) {
	length := len(arr) //数组长度
	mybstart := IntToBytes(0)
	mybstart = append(mybstart, IntToBytes(0)...)
	conn.Write(mybstart)

	for i := 0; i < length; i++ {
		mybdata := IntToBytes(1)
		mybdata = append(mybdata, IntToBytes(arr[i])...)
		conn.Write(mybdata)
	}

	mybend := IntToBytes(0)
	mybend = append(mybend, IntToBytes(1)...)
	conn.Write(mybend)

}

//
func MasterHandler(conn net.Conn, ch chan int) { //客户端函数,chan变成参数了
	<-ch                       //收到消息，才能继续进行
	msg := time.Now().String() //消息
	fmt.Println(msg)
	sendArray([]int{2, 9, 7, 6, 4}, conn) //发送数据
	fmt.Println("send over")

}

func doWork(conn net.Conn) error { //客户端函数
	ch := make(chan int, 100)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop() //延迟关闭
	for {
		select {
		case stat := <-ch:
			if stat == 2 {
				return errors.New("服务器没有消息")
			}
		case <-ticker.C: //自动时间控制
			ch <- 1
			go ServerMsgHandler(conn) //服务器消息处理函数
		case <-time.After(time.Second * 10):
			defer conn.Close()
			fmt.Println("关闭超时链接")
		}
	}
}

//处理双工通信

func ServerMsgHandler(conn net.Conn) {
	buf := make([]byte, 16)
	defer conn.Close() //处理文件关闭
	arr := []int{}     //数据接收数据
	for {
		n, err := conn.Read(buf) //读取数据
		if err != nil {
			//fmt.Println("client关闭",conn.RemoteAddr())
			return
		}
		if n == 16 {
			data1 := BytesToInt(buf[:len(buf)/2])
			data2 := BytesToInt(buf[len(buf)/2:])
			fmt.Println("data1,2", data1, data2)
			if data1 == 0 && data2 == 0 {
				arr = make([]int, 0)
			}
			if data1 == 1 {
				arr = append(arr, data2)
			}
			if data1 == 0 && data2 == 1 {
				fmt.Println("数组接收完成", arr)
				arr = make([]int, 0)
			}

		}

	}

}

func main() {

	tcpaddr, err := net.ResolveTCPAddr("tcp4", "localhost:7000")
	CheckError(err)
	for {
		conn, err := net.DialTCP("tcp", nil, tcpaddr)
		if err != nil {
			fmt.Println(err)
		} else {
			sendArray([]int{2, 9, 7, 6, 4, 12, 39}, conn)
			fmt.Println("send over")

			doWork(conn)
		}

		time.Sleep(time.Second * 10)
	}

	/*
		ch:=make(chan int ,100)//交换消息
		ticker:=time.NewTicker(time.Second)
		defer  ticker.Stop()//关闭计时器
		for {

			select {
				case<-ticker.C:
					ch<-1
					//go  MasterHandler(conn ,ch)
					go  ServerMsgHandler(conn )
				case <-time.After(time.Second*10)	:
					defer conn.Close()
					fmt.Println("time out")
			}

		}
	*/

}
