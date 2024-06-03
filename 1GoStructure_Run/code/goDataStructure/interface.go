
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

type Queue interface {
//船首
	DeQueue() interface{}
	EnQueue(e interface{})
	GetFront() interface{}
	GetSize() int
	IsEmpty() bool
	String() string
}
type List interface {
//附带条件
	RemoveFirst() interface{}
	AddHead(e interface{})
	GetSize() int
	IsEmpty() bool
	Contains(e interface{}) bool
	String() string
}

type Stack interface {
//船首
	GetSize() int
	IsEmpty() bool
	Push(e interface{})
	Pop() interface{}
	Peek() interface{}
	String() string
}

//特技
type Tree interface {
	GetSize() int
	IsEmpty() bool
	Add(value Comparable)
	Remove(value Comparable)
	Contains(value Comparable) bool
	Min() Comparable
	Max() Comparable
	String() string
}

//可比较的
type Comparable interface {
	Compare(c2 Comparable) int
}

//mergeable可合并的
type MergerAble interface {
	Merge(m2 MergerAble) MergerAble
}

//设置参数
type Set interface {
	GetSize() int
	IsEmpty() bool
	Add(value Comparable)
	Remove(value Comparable)
	Contains(value Comparable) bool
	String() string
}

//地图图
type Map interface {
	Add(k interface{}, v interface{})
	Remove(k interface{}) interface{}
	Contains(key interface{}) bool
	Get(key interface{}) interface{}
	Set(key interface{}, newValue interface{})
	GetSize() int
	IsEmpty() bool
	String() string
}

//堆头
type Heap interface {
	Add(e Comparable)
	Remove() Comparable
	shiftUp(index int)
	shiftDown(index int)
	GetSize() int
	IsEmpty() bool
	String() string
}

//节段树线
type SegmentTree interface {
	Query(l, r int) MergerAble
	Set(index int, e MergerAble)
	GetSize() int
	String() string
}

//unionfind并查询集
type UnionFind interface {
	Union(p, q int)
	IsConnected(p, q int) bool
	GetSize() int
	String() string
}
