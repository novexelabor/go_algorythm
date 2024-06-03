
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

import "fmt"

type ListStack struct {
	nl *NodeList
}

func CreateListStack() *ListStack {
	return &ListStack{CreateNodeList()}
}
func (ls *ListStack) GetSize() int {
	return ls.nl.GetSize()
}

func (ls *ListStack) IsEmpty() bool {
	return ls.nl.IsEmpty()
}

func (ls *ListStack) Push(e interface{}) {
	ls.nl.AddHead(e)
}

func (ls *ListStack) Pop() interface{} {
	return ls.nl.RemoveFirst()
}

func (ls *ListStack) Peek() interface{} {
	return ls.nl.GetFirst()
}

func (ls *ListStack) String() string {
	str := fmt.Sprint("Stack:")
	str += fmt.Sprint("top [")

	str += fmt.Sprint(ls.nl.String())

	str += fmt.Sprint("]")
//fmt.println（字符串）
	return str
}
