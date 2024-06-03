
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

import "sort"

type Student struct {
	Name  string
	Score int
}

func (p *Student) Compare(c2 Comparable) int {
	return p.Score - c2.(*Student).Score
}

type Integer int

func (i Integer) Merge(i2 MergerAble) MergerAble {
	return Integer(int(i) + int(i2.(Integer)))
}

func (i Integer) Compare(c2 Comparable) int {
	return int(i) - int(c2.(Integer))
}

type Stringer string

func (s Stringer) Compare(s2 Comparable) int {
	if s == s2 {
		return 0
	}
	if sort.StringsAreSorted([]string{string(s), string(s2.(Stringer))}) {
		return -1
	}
	return 1
}
