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
func MsgHandler(conn net.Conn) {
	//tempbuf := make([]byte, 0)
	buf := make([]byte, 16)
	defer conn.Close() //处理文件关闭
	arr := []int{}     //数据接收数据
	for {
		//time.Sleep(time.Second*1)
		n, err := conn.Read(buf) //读取数据

		if err != nil {
			fmt.Println("client没有读取到", conn.RemoteAddr())
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
				arr = make([]int, 0)
			}

		} else {
			//continue
			fmt.Printf("not 16")
		}
		//tempbuf = make([]byte, 0)
		beatch := make(chan byte)
		go HearBeat(conn, beatch, 15)
		go HeartChanHander(buf[:1], beatch)

	}
}

//心跳机制
func HearBeat(conn net.Conn, heartchan chan byte, timeout int) {

	select {
	case t := <-time.After(time.Second * 2):
		fmt.Println("time is out close", t)
		conn.Close()
	case hc := <-heartchan:
		fmt.Println("<-heartchan", string(hc))
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break

	}
}

//channel处理心跳
func HeartChanHander(n []byte, beatch chan byte) {
	fmt.Println("HeartChanHander", len(n))
	for _, v := range n {
		beatch <- v //管道压入数据
	}
	close(beatch)
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
