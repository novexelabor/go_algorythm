package main

import (
	"fmt"
	"strings"
)

//a<b<c  首先比较第一个字母，
//  左边小于右边 -1 .左边大于右边+1,
//  第一个字母比较不成功比较第二个
func main() {
	fmt.Println(3 > 2)
	//字符串比较调用strings.Compare()方法
	fmt.Println(strings.Compare("ba", "bc"))
	fmt.Println(strings.Compare("bs", "be"))

	//这是字符串地址的比较
	fmt.Println("ab" < "fd")
}
