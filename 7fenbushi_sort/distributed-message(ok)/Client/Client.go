package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func IntTobytes(n int) []byte {
	data := int64(n)
	bytebuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytebuffer, binary.BigEndian, data)
	return bytebuffer.Bytes()
}
func BytesToInt(bts []byte) int { //int64
	bytebuffer := bytes.NewBuffer(bts)
	var data int64
	binary.Read(bytebuffer, binary.BigEndian, &data)
	return int(data)
}

func ServerMsgHandler(conn net.Conn) {
	buf := make([]byte, 16)
	defer conn.Close() //延迟关闭
	arr := []int{}     //数组保存数据
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Sever  close")
			return
		}
		if n == 16 {
			data1 := BytesToInt(buf[:len(buf)/2]) //取出的第一个数据
			data2 := BytesToInt(buf[len(buf)/2:]) //取出第二个数据
			if data1 == 0 && data2 == 0 {
				arr = make([]int, 0) //开辟数据
			}
			if data1 == 1 {
				arr = append(arr, data2)
			}
			if data1 == 0 && data2 == 1 {
				fmt.Println("数组接收完成", arr)
				arr = make([]int, 0) //开辟数据
			}

		}

	}

}

func main() {
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8848")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpaddr) //链接
	if err != nil {
		panic(err)
	}
	go ServerMsgHandler(conn) //协程接受数据

	arr := []int{1, 9, 2, 8, 7, 3, 5, 6, 10, 4, 23, 24}
	length := len(arr)
	// -1  "1"
	// 8 abcdefgh
	//4  abcd
	//-1  "0"

	//0 0  开始传输
	//1 1
	//1 9
	//1 2
	//1 8
	//1  7
	//1  3
	//1  5
	//1 6
	//1 10
	//1 4
	//0 1 //结束传输
	mybstart := IntTobytes(0)
	mybstart = append(mybstart, IntTobytes(0)...)
	conn.Write(mybstart)

	for i := 0; i < length; i++ {
		mybdata := IntTobytes(1)
		mybdata = append(mybdata, IntTobytes(arr[i])...)
		conn.Write(mybdata)
	}

	mybend := IntTobytes(0)
	mybend = append(mybend, IntTobytes(1)...)
	conn.Write(mybend)

	time.Sleep(time.Second * 30)

	/*
		for{




			var inpustr string
			fmt.Scanln(&inpustr)
			conn.Write([]byte(inpustr))
			buf:=make([]byte,1024)
			n,_:=conn.Read(buf)//读取数据
			fmt.Println(string(buf[:n]))
		}*/

}
