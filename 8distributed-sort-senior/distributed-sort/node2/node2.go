package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"sort"
	"time"
)

//处理错误
func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//字节转化整数
func BytesToInt(b []byte) int {
	bytesBuf := bytes.NewBuffer(b) //空的字节数组
	var data int64
	binary.Read(bytesBuf, binary.BigEndian, &data) //解码
	return int(data)
}
func IntToBytes(n int) []byte {
	data := int64(n)                               //64位电脑8字节
	bytesBuf := bytes.NewBuffer([]byte{})          //空的字节数组
	binary.Write(bytesBuf, binary.BigEndian, data) //写入数据
	return bytesBuf.Bytes()                        //返回字节
}
func MsgHandler(conn net.Conn) {
	buf := make([]byte, 16)
	defer conn.Close() //处理文件关闭
	arr := []int{}     //数据接收数据
	for {
		n, err := conn.Read(buf) //读取数据
		if err != nil {
			fmt.Println("client关闭", conn.RemoteAddr())
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
				sort.Ints(arr)
				fmt.Println("数组排序完成", arr)

				sendArray(arr, conn) //排序完成的数组，返回给服务器

				arr = make([]int, 0)
			}

		}
		msg := buf[1:3] //备份buf
		beatch := make(chan byte)
		go HearBeat(conn, beatch, 5)
		go HeartChanHander(msg, beatch)

	}
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

//心跳机制
func HearBeat(conn net.Conn, heartchan chan byte, timeout int) {
	select {
	case hc := <-heartchan:
		fmt.Println("<-heartchan", string(hc))
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break
	case <-time.After(time.Second * 5):
		fmt.Println("time is out")
		conn.Close()
	}
}

//channel处理心跳
func HeartChanHander(n []byte, beatch chan byte) {
	for _, v := range n {
		beatch <- v //管道压入数据
	}
	close(beatch)
}

func main() {
	server, err := net.Listen("tcp", "localhost:7002") //创建服务器
	CheckError(err)                                    //处理错误
	defer server.Close()                               //延迟关闭服务器
	for {
		new_conn, err := server.Accept() //接收数据信息
		CheckError(err)                  //处理错误
		go MsgHandler(new_conn)          //并发处理链接信息

	}

}
