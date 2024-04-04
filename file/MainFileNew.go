package main

import (
	"fmt"
	"io/ioutil"
)

//level求解文件夹深度
var dep int = 0 //var声明全局变量

func GetALLX(path string, level int) {
	if level > dep {
		dep = level
	}
	fmt.Println("level", level)
	levelstr := ""
	if level == 1 {
		levelstr = "+"
	} else {
		for ; level > 1; level-- {
			levelstr += "|--"
		}
		levelstr += "+"
	}

	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return
	}
	for _, fi := range read { //循环每个文件或者文件夹
		if fi.IsDir() { //判断是否文件夹
			fulldir := path + "/" + fi.Name() //构造新的路径
			fmt.Println(levelstr + fulldir)
			//newlevel:=level+1
			//fmt.Println("call")
			GetALLX(fulldir, level+1) //文件夹递归处理

		} else {
			fulldir := path + "/" + fi.Name() //构造新的路径
			fmt.Println(levelstr + fulldir)
		}

	}

}

func main4() {
	path := "/Users/renshanwan/Documents/wanlin" //mac上直接粘贴路径
	GetALLX(path, 1)
}

//求解最大深度
func main() {
	path := "/Users/renshanwan/Documents/wanlin" //mac上直接粘贴路径
	GetALLX(path, 1)

	fmt.Println(dep) //打印最大深度数值
}
