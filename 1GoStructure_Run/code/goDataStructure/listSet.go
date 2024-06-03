
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

type ListSet struct {
	il *ImproveList
}

func CreateListSet() *ListSet {
	return &ListSet{CreateImproveList()}
}
func (ls *ListSet) GetSize() int {
	return ls.il.GetSize()
}

func (ls *ListSet) IsEmpty() bool {
	return ls.il.IsEmpty()
}

func (ls *ListSet) Add(value Comparable) {
	if !ls.il.Contains(value) {
		ls.il.AddHead(value)
	}
}

func (ls *ListSet) Remove(value Comparable) {
	ls.il.RemoveFirst()
}

func (ls *ListSet) Contains(value Comparable) bool {
	return ls.il.Contains(value)
}
func (ls *ListSet) String() string {
	return ls.il.String()
}
