package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
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
func MasterHandler(conn net.Conn, ch chan int) {
	<-ch
	msg := time.Now().String() //消息
	fmt.Println(msg, len(ch))
	time.Sleep(time.Second * 150)
	sendArray([]int{2, 9, 7, 6, 4}, conn)
	fmt.Println("send over")
}

func main() {

	tcpaddr, err := net.ResolveTCPAddr("tcp4", "localhost:7000")
	CheckError(err)
	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	CheckError(err)
	sendArray([]int{2, 9, 7, 6, 4, 12, 39}, conn)
	fmt.Println("send over")

	ch := make(chan int, 100) //交换消息
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop() //关闭计时器
	for {

		select {
		case <-ticker.C:
			ch <- 1
			go MasterHandler(conn, ch)
		case <-time.After(time.Second * 10):
			fmt.Println("time out")
			conn.Close()
		}

	}

}
