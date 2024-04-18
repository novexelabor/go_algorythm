package main

import (
	"fmt"
	"sync"
	"time"
)

//map映射
//map管理上亿的数据，瞬间查找
//多线程需要维护线程安全,多线程访问同一个临界资源时，需要维护线程安全

type SyncMap struct {
	mymap         map[string]string
	*sync.RWMutex //读写锁
}

var smap SyncMap   //公有的访问map,全局变量
var done chan bool //通道，是否完成，用来同步消息

func write1() {
	keys := []string{"1", "2", "3"}
	for _, k := range keys {
		smap.Lock()
		smap.mymap[k] = k //赋值
		smap.Unlock()
		time.Sleep(1 * time.Second)
	}
	done <- true //通道写入我们干完了
}
func write2() {
	keys := []string{"a1", "b2", "c3"}
	for _, k := range keys {
		smap.Lock()
		smap.mymap[k] = k //赋值
		smap.Unlock()
		time.Sleep(1 * time.Second)
	}
	done <- true //通道写入我们干完了
}

func read() {
	smap.RLock() //读的时候枷锁,读锁，读者可以继续申请
	fmt.Println("readLock")
	for k, v := range smap.mymap {
		fmt.Println(k, v)
	}
	smap.RUnlock() //Rlock()和RUnlock()是成对使用的
}

func main() {
	smap = SyncMap{make(map[string]string), new(sync.RWMutex)}
	done = make(chan bool, 2) //管道处理写入
	go write1()
	go write2()

	for {
		read()
		if len(done) == 2 {
			fmt.Println(smap.mymap)
			for k, v := range smap.mymap {
				fmt.Println(k, v)
			}
			break
		} else {
			time.Sleep(1 * time.Second) //没有完成继续等待
		}
	}

}
