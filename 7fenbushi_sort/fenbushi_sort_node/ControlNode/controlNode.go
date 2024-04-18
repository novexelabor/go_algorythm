package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"go_algorythm/7fenbushi_sort/fenbushi_sort_node/pipelineMiddleWare"
	"net"
	"strconv"
	"time"
)

func IntTobytes(n int) []byte { //Write()方法
	data := int64(n)
	bytebuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytebuffer, binary.BigEndian, data)
	return bytebuffer.Bytes()
}
func BytesToInt(bts []byte) int { //Read()方法
	bytebuffer := bytes.NewBuffer(bts)
	var data int64
	binary.Read(bytebuffer, binary.BigEndian, &data)
	return int(data)
}

//处理器函数返回channel，来传输数据，从而实现多线程处理
func ServerMsgHandler(conn net.Conn) <-chan int {
	out := make(chan int, 1024) //新的管道
	defer conn.Close()          //延迟关闭链接
	buf := make([]byte, 16)     //两个int64的数
	arr := []int{}              //数组保存数据
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Sever  close")
			return nil
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
				for i := 0; i < len(arr); i++ {
					out <- arr[i] //数组压入管道,channel用来线程之间传输，
					//这个channel比重新写入到conn中高效简洁的多
					//消息网络传递，分布式高并发(多线程)的处理
				}
				close(out) //关闭管道
				return out

			}

		}

	}

}

//把用链接发送数据，写成func
func SendArray(arr []int, conn net.Conn) {
	length := len(arr)
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
}

func main() {
	arrlist := [][]int{{1, 9, 2, 8, 7, 3, 5, 6, 10, 4, 23, 24}, {11, 19, 12, 18, 17, 13, 15, 16, 101, 14, 123, 124}}
	sortResults := []<-chan int{} //为空

	for i := 0; i < 2; i++ {
		tcpaddr, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(5000+i))
		if err != nil {
			panic(err)
		}
		//addr := "127.0.0.1:" + strconv.Itoa(5000+i)
		//conn, err := net.Dial("tcp", addr)
		conn, err := net.DialTCP("tcp", nil, tcpaddr) //链接
		if err != nil {
			panic(err)
		}

		SendArray(arrlist[i], conn) //写入数据
		//这里的服务器函数参数就是链接conn，如果是string IP地址，需要Listen()和Accept()方法返回conn
		sortResults = append(sortResults, ServerMsgHandler(conn)) //conn

	}
	fmt.Println(len(sortResults))
	last := pipelineMiddleWare.Merge(sortResults[0], sortResults[1])
	for v := range last {
		fmt.Printf("%d ", v)
	}
	time.Sleep(time.Second * 30)

}

//监听
// func listen1(addr string) net.Conn {
// 	listen, err := net.Listen("tcp", addr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer listen.Close()

// 	conn, err := listen.Accept()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return conn
// }
