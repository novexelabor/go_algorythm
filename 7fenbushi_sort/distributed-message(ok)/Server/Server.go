package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sort"
)

func BytesToInt(bts []byte) int {
	bytebuffer := bytes.NewBuffer(bts)
	var data int64
	binary.Read(bytebuffer, binary.BigEndian, &data)
	return int(data)
}
func IntTobytes(n int) []byte {
	data := int64(n)
	bytebuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytebuffer, binary.BigEndian, data)
	return bytebuffer.Bytes()
}
func MsgHandler(conn net.Conn) {
	buf := make([]byte, 16)
	defer conn.Close()
	arr := []int{} //数组保存数据
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client close")
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
				sort.Ints(arr) //排序
				fmt.Println("数组排序完成", arr)
				//写入
				mybstart := IntTobytes(0)
				mybstart = append(mybstart, IntTobytes(0)...)
				conn.Write(mybstart)

				for i := 0; i < len(arr); i++ {
					mybdata := IntTobytes(1)
					mybdata = append(mybdata, IntTobytes(arr[i])...)
					conn.Write(mybdata)
				}

				mybend := IntTobytes(0)
				mybend = append(mybend, IntTobytes(1)...)
				conn.Write(mybend)

				arr = make([]int, 0) //开辟数据
			}

		}

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
