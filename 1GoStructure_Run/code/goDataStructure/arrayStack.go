
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

import "fmt"

type ArrayStack struct {
	a *arr
}

func CreateArrayStack(cap int) *ArrayStack {
	return &ArrayStack{CreateArray(cap)}
}

func (as *ArrayStack) GetSize() int {
	return as.a.GetSize()
}

func (as *ArrayStack) IsEmpty() bool {
	return as.a.IsEmpty()
}

func (as *ArrayStack) Push(e interface{}) {
	as.a.AddLast(e)
}

func (as *ArrayStack) Pop() interface{} {
	return as.a.RemoveLast()
}

func (as *ArrayStack) Peek() interface{} {
	return as.a.GetLast()
}
func (as *ArrayStack) String() string {
	str := fmt.Sprint("Stack:")
	str += fmt.Sprint(("["))
	for i := 0; i < as.a.size; i++ {
		str += fmt.Sprint((*as).a.data[i])
		if i != as.a.size-1 {
			str += fmt.Sprint(",")
		}
	}
	str += fmt.Sprint("] top")
//fmt.println（字符串）
	return str
}
