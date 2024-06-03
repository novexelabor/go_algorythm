
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

type TrieMap struct {
	root *trieMapNode
	size int
}
type trieMapNode struct {
	isWord bool
	value  interface{}
	next   map[rune]*trieMapNode
}

func (tm *TrieMap) Add(k interface{}, v interface{}) { //完成Triemap实现词频统计
	panic("implement me")
}

func (tm *TrieMap) Remove(k interface{}) interface{} {
	panic("implement me")
}

func (tm *TrieMap) Contains(key interface{}) bool {
	panic("implement me")
}

func (tm *TrieMap) Get(key interface{}) interface{} {
	panic("implement me")
}

func (tm *TrieMap) Set(key interface{}, newValue interface{}) {
	panic("implement me")
}

func (tm *TrieMap) GetSize() int {
	panic("implement me")
}

func (tm *TrieMap) IsEmpty() bool {
	panic("implement me")
}

func (tm *TrieMap) String() string {
	panic("implement me")
}
