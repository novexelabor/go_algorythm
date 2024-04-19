package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
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
func BytesToInt(b []byte) int {
	bytesBuf := bytes.NewBuffer(b) //空的字节数组
	var data int64
	binary.Read(bytesBuf, binary.BigEndian, &data) //解码
	return int(data)
}
func sendArray(arr []int, conn net.Conn) { //参数链接conn
	length := len(arr) //数组长度
	mybstart := IntToBytes(0)
	mybstart = append(mybstart, IntToBytes(0)...)
	conn.Write(mybstart) //2个字节00

	for i := 0; i < length; i++ {
		mybdata := IntToBytes(1)
		mybdata = append(mybdata, IntToBytes(arr[i])...)
		conn.Write(mybdata) //写入数组数据
	}

	mybend := IntToBytes(0)
	mybend = append(mybend, IntToBytes(1)...)
	conn.Write(mybend) //01

}

func ServerMsgHandler(conn net.Conn) <-chan int { //网络通信+channel
	out := make(chan int, 1024)
	buf := make([]byte, 16) //每次读取16byte，2字节的数据
	defer conn.Close()      //处理文件关闭
	arr := []int{}          //数据接收数据
	for {                   //for循环来接受数据
		n, err := conn.Read(buf) //读取数据
		if err != nil {
			//fmt.Println("client关闭",conn.RemoteAddr())
			return nil
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
				for i := 0; i < len(arr); i++ {
					out <- arr[i] //数组压入管道
				}
				close(out)
				return out

				//arr = make([]int, 0, 0)
			}

		}

	}

}
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024) //先开辟
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 { //循环取出最小值，归并之后有序
			if !ok2 || (ok1 && v1 <= v2) { //v2没有值或者有值且大于等于v1
				out <- v1 //取出v1
				v1, ok1 = <-in1
			} else {
				out <- v2 //取出v2
				v2, ok2 = <-in2
			}

		}

		close(out) //for range遍历，chan关闭时，遍历结束
		fmt.Printf("归并排序完成")
	}()
	return out
}

func main() {
	arrlist := [][]int{{1, 109, 2, 107}, {100, 103, 101, 102}}
	sortResult := []<-chan int{}
	for i := 0; i < 2; i++ {
		tcpaddr, err := net.ResolveTCPAddr("tcp4", "localhost:700"+strconv.Itoa(i+1))
		CheckError(err)
		conn, err := net.DialTCP("tcp", nil, tcpaddr)
		CheckError(err)
		sendArray(arrlist[i], conn)
		sortResult = append(sortResult, ServerMsgHandler(conn)) //处理数据接收

	}
	fmt.Println(sortResult)
	fmt.Println(sortResult[0])
	fmt.Println(sortResult[1])
	last := Merge(sortResult[0], sortResult[1]) //归并排序

	for v := range last {
		fmt.Printf("%d  ", v)
	}
	time.Sleep(30 * time.Second)

}
