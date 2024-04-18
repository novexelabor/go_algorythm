package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//整数，字符串与字节的转换

func IntTobytes(n int) []byte {
	data := int64(n)
	bytebuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytebuffer, binary.BigEndian, data)
	return bytebuffer.Bytes()
}

func BytesToInt(bts []byte) int {
	bytebuffer := bytes.NewBuffer(bts)
	var data int64
	binary.Read(bytebuffer, binary.BigEndian, &data)
	return int(data)
}

func main() {
	fmt.Println(IntTobytes(1))
	fmt.Println(BytesToInt(IntTobytes(1)))
	fmt.Println([]byte("123"))
	fmt.Println(string([]byte("123")))
	myb := IntTobytes(1)
	myb = append(myb, IntTobytes(1)...)
	fmt.Println(myb)
	fmt.Println(myb[:len(myb)/2])

}

//验证一下没有问题
