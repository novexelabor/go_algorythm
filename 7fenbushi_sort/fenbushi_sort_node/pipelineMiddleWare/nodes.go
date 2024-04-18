package pipelineMiddleWare

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var starttime time.Time //构造时间

func Init() {
	starttime = time.Now() //初始化
}

func UseTime() {
	fmt.Println(time.Since(starttime)) //统计消耗时间
}

//内存排序   只读：<-chan   ， 只写：chan<-
func InMemorySort(in <-chan int) <-chan int {
	out := make(chan int, 1024) //新的管道
	go func() {
		data := []int{} //创建一个数组，储存数据并且排序
		for v := range in {
			data = append(data, v) //数据压入数组
		}
		fmt.Println("数据读取完成", time.Since(starttime))
		//先获取数据，获取完数据之后，然后排序
		sort.Ints(data) //排序
		for _, v := range data {
			out <- v //压入数据
		}
		fmt.Println("排序完成")
		close(out) //关闭管道
	}()
	return out
}

//合并,两个管道的数据有序，归并有序的数据压入到另外一个管道
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024) //新的管道
	go func() {
		fmt.Println("归并开始")
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		//归并排序
		for ok1 || ok2 { //通道关闭，都是false
			if !ok2 || (ok1 && v1 <= v2) { //写入v1，所以OK1一定是true，且v1小于等于v2
				out <- v1 //取出V1，压入，再次读取v1 ； v2不存在，OK2为false，直接写入即可
				v1, ok1 = <-in1

			} else {
				out <- v2 //取出V2，压入，再次读取v2
				v2, ok2 = <-in2
			}
		}
		close(out) //没有数值读取了，则直接关闭归并的通道
		fmt.Println("归并结束")
	}()

	return out //归并通道读取的数据，然后把归并好的通道返回，方便协程读取数据
}

//reader 只要实现了read()方法都实现了该接口
//io.Open()打开文件;strings和bytes数组都可以转化为reader接口
//读取数据
func ReaderSource(reader io.Reader, chunksize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buf := make([]byte, 8) //64
		readsize := 0
		for {
			n, err := reader.Read(buf)
			readsize += n
			if n > 0 {
				//每次通道传送的数据是Uint64的整数
				out <- int(binary.BigEndian.Uint64(buf)) //数据压入
			} //超过chunksize的范围了
			if err != nil || (chunksize != -1 && readsize >= chunksize) {
				break //跳出循环
			}
		}

		close(out)
	}()
	return out

}

//写入
func WriterSlink(writer io.Writer, in <-chan int) {
	for v := range in { //for range在通道close()时停止
		buf := make([]byte, 8)                     //64位 8字节
		binary.BigEndian.PutUint64(buf, uint64(v)) //字节转换
		writer.Write(buf)                          //写入
	}

}

//随机数数组
func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int() //压入随机数
		}
		close(out) //关闭管道
	}()

	return out
}

//多路合并5，多个参数，inputs相当于是切片数组
func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	} else {
		//分为两个部分即可
		m := len(inputs) / 2
		return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...)) //递归
	}
}

//num是切片数组
func ArraySource(num ...int) <-chan int {
	var out = make(chan int) //该通道是没有缓冲的
	go func() {
		for _, v := range num {
			out <- v //数组的数据压入进去
		}
		close(out)
	}()
	return out
}
