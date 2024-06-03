
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

type avlTreeSet struct {
	atm *AvlTreeMap
}

func CreateAvlTreeSet() *avlTreeSet {
	return &avlTreeSet{
		atm: CreateAvlTreeMap(),
	}
}
func (ats *avlTreeSet) GetSize() int {
	return ats.atm.GetSize()
}

func (ats *avlTreeSet) IsEmpty() bool {
	return ats.atm.IsEmpty()
}

func (ats *avlTreeSet) Add(value Comparable) {
	ats.atm.Add(value, nil)
}

func (ats *avlTreeSet) Remove(value Comparable) {
	ats.atm.Remove(value)
}

func (ats *avlTreeSet) Contains(value Comparable) bool {
	return ats.atm.Contains(value)
}

func (ats *avlTreeSet) String() string {
	return ats.atm.String()
}
