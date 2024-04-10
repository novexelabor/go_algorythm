package main

import (
	"bufio"
	"fmt"
	"go_algorythm/5Link/code/Double_Link"
	"go_algorythm/5Link/code/Single_Link"
	"io"
	"os"
	"time"
	// "go_algorythm/5Link/code/Double_Link"
)

func main1() {
	list := Single_Link.NewSingleLinkList()
	node1 := Single_Link.NewSingleLinkNode(1)
	node2 := Single_Link.NewSingleLinkNode(2)
	node3 := Single_Link.NewSingleLinkNode(3)
	list.InsertNodeFront(node1)
	fmt.Println(list)
	list.InsertNodeFront(node2)
	fmt.Println(list)
	list.InsertNodeFront(node3)
	fmt.Println(list)
}

func main2() {
	list := Single_Link.NewSingleLinkList()
	node1 := Single_Link.NewSingleLinkNode(1)
	node2 := Single_Link.NewSingleLinkNode(2)
	node3 := Single_Link.NewSingleLinkNode(3)
	list.InsertNodeBack(node1)
	fmt.Println(list)
	list.InsertNodeBack(node2)
	fmt.Println(list)
	list.InsertNodeBack(node3)
	fmt.Println(list)
	node4 := Single_Link.NewSingleLinkNode(4)
	//list.InsertNodeValueBack(2,node4)
	list.InsertNodeValueFront(2, node4)
	fmt.Println(list)
	fmt.Println(list.GetNodeAtIndex(2))

	list.DeleteNode(node4)
	fmt.Println(list)
	list.Deleteatindex(1)
	fmt.Println(list)
}

func main3() {
	list := Single_Link.NewSingleLinkList()

	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day4\\猴岛游戏社区\\houdao\\1_1.txt" //路径
	sqlfile, _ := os.Open(path)                                                    //打开文件
	i := 0                                                                         //统计行数
	br := bufio.NewReader(sqlfile)                                                 //读取文件对象
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF { //文件结束跳出循环
			break
		}
		linestr := string(line) //转化为字符串
		nodestr := Single_Link.NewSingleLinkNode(linestr)
		list.InsertNodeFront(nodestr) //数据插入链表
		//fmt.Println(linestr)
		i++

	}
	fmt.Println(i, "内存载入完成")

	for {
		fmt.Println("请输入要查询的用户名")
		var inputstr string
		fmt.Scanln(&inputstr) //用户输入
		starttime := time.Now()
		list.FindString(inputstr)

		fmt.Println("本次查询用了", time.Since(starttime))

	}

}

func main4() {
	list := Single_Link.NewSingleLinkList()
	node1 := Single_Link.NewSingleLinkNode(1)
	node2 := Single_Link.NewSingleLinkNode(2)
	node3 := Single_Link.NewSingleLinkNode(3)
	node4 := Single_Link.NewSingleLinkNode(4)
	list.InsertNodeBack(node1)
	fmt.Println(list)
	list.InsertNodeBack(node2)
	fmt.Println(list)
	list.InsertNodeBack(node3)
	fmt.Println(list)
	list.InsertNodeBack(node4)
	fmt.Println(list)
	//fmt.Println(list)
	//fmt.Println(list.GetMid().Value())
	list.ReverseList()
	fmt.Println(list)

}

func main5() {
	dlist := Double_Link.NewDoubleLinkList()
	node1 := Double_Link.NewDoubleLinkNode(1)
	node2 := Double_Link.NewDoubleLinkNode(2)
	node3 := Double_Link.NewDoubleLinkNode(3)
	node4 := Double_Link.NewDoubleLinkNode(4)
	node5 := Double_Link.NewDoubleLinkNode(5)
	dlist.InsertHead(node1)
	dlist.InsertHead(node2)
	dlist.InsertHead(node3)
	dlist.InsertHead(node4)
	dlist.InsertHead(node5)
	//node6:=Double_Link.NewDoubleLinkNode(6)
	//node7:=Double_Link.NewDoubleLinkNode(6)
	//dlist.InsertValueHead(node3,node6)
	fmt.Println(dlist.String())
	//dlist.InsertValueBack(node3,node6)
	//dlist.InsertValueHead(node3,node7)
	//dlist.InsertValueBackByValue(3,node7)
	//dlist.DeleteNode(node2)
	dlist.DeleteNodeAtindex(3)
	fmt.Println(dlist.String())
}

func main() {
	pathlist := []string{"C:\\Users\\Tsinghua-yincheng\\Desktop\\day4\\猴岛游戏社区\\houdao\\1_1.txt",
		"C:\\Users\\Tsinghua-yincheng\\Desktop\\day4\\猴岛游戏社区\\houdao\\1_2.txt",
		"C:\\Users\\Tsinghua-yincheng\\Desktop\\day4\\猴岛游戏社区\\houdao\\1_3.txt"}
	dlist := Double_Link.NewDoubleLinkList()
	for i := 0; i < len(pathlist); i++ {
		path := pathlist[i]            //路径
		sqlfile, _ := os.Open(path)    //打开文件
		br := bufio.NewReader(sqlfile) //读取文件对象
		for {
			line, _, end := br.ReadLine()
			if end == io.EOF { //文件结束跳出循环
				break
			}
			linestr := string(line)                        //转化为字符串
			node := Double_Link.NewDoubleLinkNode(linestr) //新建一个节点
			dlist.InsertHead(node)                         //插入节点

		}
	}
	fmt.Println("内存载入完成", dlist.Getlength())

	for {
		fmt.Println("请输入要查询的用户名")
		var inputstr string
		fmt.Scanln(&inputstr) //用户输入
		starttime := time.Now()
		dlist.FindString(inputstr)

		fmt.Println("本次查询用了", time.Since(starttime))

	}

}
