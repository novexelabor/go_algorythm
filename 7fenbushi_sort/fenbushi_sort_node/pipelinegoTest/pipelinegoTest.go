package main

import (
	"bufio"
	"fmt"
	"time"

	//有时候无法通过是因为使用了汉字，变成因为就好了
	"go_algorythm/7fenbushi_sort/fenbushi_sort_node/pipelineMiddleWare"
	"os"
)

func main1() {
	var filename = "7fenbushi_sort/fenbushi_sort_node/data1.in" //二进制文件写入
	var count = 100000

	file, err := os.Create(filename) //创建打开了一个文件
	if err != nil {
		panic(err)
	}
	defer file.Close()

	mypipe := pipelineMiddleWare.RandomSource(count) //
	writer := bufio.NewWriter(file)                  //写入
	pipelineMiddleWare.WriterSlink(writer, mypipe)   //写入
	writer.Flush()                                   //刷新,写入数据之后要刷新Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	mypipeX := pipelineMiddleWare.ReaderSource(bufio.NewReader(file), -1)
	counter := 0
	for v := range mypipeX {
		fmt.Println(v)
		counter++
		if counter > 1000 {
			break
		}
	}
}

func main() {
	go func() {
		myp := pipelineMiddleWare.Merge(
			pipelineMiddleWare.InMemorySort(pipelineMiddleWare.ArraySource(3, 9, 2, 1, 10)),
			pipelineMiddleWare.InMemorySort(pipelineMiddleWare.ArraySource(13, 19, 12, 11, 110)),
		)

		for v := range myp {
			fmt.Println(v)
		}
	}()
	time.Sleep(time.Second * 10)
}
