
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

type BSTreeSet struct {
	bst *BSTreeMap
}

func CreateBSTreeSet() *BSTreeSet {
	return &BSTreeSet{CreateBSTreeMap()}
}

func (ts *BSTreeSet) GetSize() int {
	return ts.bst.GetSize()
}

func (ts *BSTreeSet) IsEmpty() bool {
	return ts.bst.IsEmpty()
}

func (ts *BSTreeSet) Add(value Comparable) {
	ts.bst.Add(value, nil)
}

func (ts *BSTreeSet) Remove(value Comparable) {
	ts.bst.Remove(value)
}

func (ts *BSTreeSet) Contains(value Comparable) bool {
	return ts.bst.Contains(value)
}
func (ts *BSTreeSet) String() string {
	return ts.bst.String()
}
