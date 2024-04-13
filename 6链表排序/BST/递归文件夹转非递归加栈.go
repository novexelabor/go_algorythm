package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

//递归函数
func GetAll(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return files, errors.New("文件夹不能读取")
	}
	for _, fi := range read {
		if fi.IsDir() {
			fulldir := path + "\\" + fi.Name() //文件夹路径
			files = append(files, fulldir)     //插入列表
			files, err = GetAll(fulldir, files)
			if err != nil {
				return files, errors.New("文件夹不能读取")
			}

		} else {
			fullname := path + "\\" + fi.Name() //处理文件
			files = append(files, fullname)     //插入列表
		}
	}
	return files, nil

}

//非递归遍历文件夹
type Node struct {
	data string
	next *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(value string)
	Pop() (string, error)
	Length() int
}

func NewStack() *Node {
	return &Node{}
}

func (n *Node) IsEmpty() bool { //判断是否为空
	return n.next == nil
}
func (n *Node) Push(value string) {
	newnode := &Node{data: value} //初始化
	newnode.next = n.next
	n.next = newnode
}
func (n *Node) Pop() (string, error) {
	if n.IsEmpty() == true {
		return "", errors.New("bug")
	}
	value := n.next.data
	n.next = n.next.next
	return value, nil
}
func (n *Node) Length() int {
	pnext := n
	length := 0
	for pnext.next != nil { //返回长度
		pnext = pnext.next
		length++
	}
	return length
}
func main() {
	files := []string{}
	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day4"

	mystack := NewStack() //生成栈
	mystack.Push(path)    //入栈

	for !mystack.IsEmpty() { //判断栈非空
		path, err := mystack.Pop() //弹出栈
		if err != nil {
			break
		}
		files = append(files, path)       //处理
		read, err := ioutil.ReadDir(path) //读取文件夹
		if err != nil {
			break
		}
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path + "\\" + fi.Name() //文件夹路径
				//files=append(files,fulldir) //插入列表
				mystack.Push(fulldir)

			} else {
				fullname := path + "\\" + fi.Name() //处理文件
				files = append(files, fullname)     //插入列表
			}
		}

	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}

}
func main1() {
	files := []string{}
	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day3"
	files, _ = GetAll(path, files)
	//fmt.Printf("%v",files)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
