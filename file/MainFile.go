package main

import (
	"errors"
	"fmt"
	"go_algorythm/Queue"
	"go_algorythm/StackArray"
	"io/ioutil"
)

//递归遍历文件夹
func GetAll(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _, fi := range read { //range遍历
		if fi.IsDir() { //判断是否是文件夹
			fulldir := path + "/" + fi.Name() // 构造新路径
			files = append(files, fulldir)    //追加路径
			files, _ = GetAll(fulldir, files)
		} else {
			fulldir := path + "/" + fi.Name() // 构造新路径
			files = append(files, fulldir)    //追加路径
		}
	}

	return files, nil
}

func main1() {
	path := "/Users/renshanwan/Documents/wanlin" //mac上直接粘贴路径
	files := []string{}                          //字符串数组，初始化为空
	files, _ = GetAll(path, files)               //抓取文件夹

	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

//使用栈实现非递归遍历文件夹
func main2() {
	path := "/Users/renshanwan/Documents/wanlin" //mac上直接粘贴路径
	files := []string{}

	mystack := StackArray.NewStack()
	mystack.Push(path)

	for !mystack.IsEmpty() {
		path := mystack.Pop().(string) // 强制类型转换成string
		files = append(files, path)
		read, _ := ioutil.ReadDir(path) //抓取路径下所有的路径

		//read ,err := ioutil.ReadDir(path)  //抓取路径下所有的路径
		// if err != nil {  //添加err循环内该语句会出错
		// 	return
		// }

		for _, fi := range read { //遍历
			if fi.IsDir() {
				fulldir := path + "/" + fi.Name()
				mystack.Push(fulldir) //文件夹压入栈
			} else {
				fulldir := path + "/" + fi.Name()
				files = append(files, fulldir) //追加路径
			}
		}
	}
	for i := 0; i < len(files); i++ { //打印
		fmt.Println(files[i])
	}
}

//使用队列实现非递归遍历文件夹
func main3() {
	path := "/Users/renshanwan/Documents/wanlin" //mac上直接粘贴路径
	files := []string{}

	myqueue := Queue.NewQueue()
	myqueue.EnQueue(path)

	for !myqueue.IsEmpty() {
		path := myqueue.DeQueue()
		files = append(files, path.(string)) //空接口强制类型转换成字符串

		read, _ := ioutil.ReadDir(path.(string)) //抓取路径下所有的路径
		for _, fi := range read {                //遍历
			if fi.IsDir() {
				fulldir := path.(string) + "/" + fi.Name()
				myqueue.EnQueue(fulldir) //文件夹压入栈
			} else {
				fulldir := path.(string) + "/" + fi.Name()
				files = append(files, fulldir) //追加路径
			}
		}
	}
	for i := 0; i < len(files); i++ { //打印
		fmt.Println(files[i])
	}
	fmt.Println("完成任务")
}
