
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

import "fmt"

type ArraySegmentTree struct {
	data []MergerAble
	tree []MergerAble
//兼并合并
}

//类型合并功能（接口，接口）接口
func TransInts(original []int) []MergerAble {
	ret := make([]MergerAble, len(original))
	for key, value := range original {
		ret[key] = Integer(value)
	}
	return ret
}
func CreateSegmentTree(arr []MergerAble) *ArraySegmentTree {
	tree := make([]MergerAble, len(arr)*4)
	st := &ArraySegmentTree{
		data: arr,
		tree: tree,
//合并：
	}
	st.buildSegmentTree(0, 0, len(arr)-1)
	return st
}
func (st *ArraySegmentTree) buildSegmentTree(index, l, r int) {
	if l == r {
		st.tree[index] = st.data[l]
		return
	}

	left := st.leftChild(index)
	right := st.rightChild(index)
	mid := l + (r-l)/2

	st.buildSegmentTree(left, l, mid)
	st.buildSegmentTree(right, mid+1, r)

	st.tree[index] = st.tree[left].Merge(st.tree[right])
}

func (st *ArraySegmentTree) Query(qleft, qright int) MergerAble {
	if qleft < 0 || qleft >= len(st.data) || qright < 0 || qright >= len(st.data) || qleft > qright {
		panic("query condition is illegal")
	}

	return st.query(0, 0, len(st.data)-1, qleft, qright)
}

func (st *ArraySegmentTree) query(index, l, r, qleft, qright int) MergerAble {
	if l == r {
		return st.tree[index]
	}

	left := st.leftChild(index)
	right := st.rightChild(index)
	mid := l + (r-l)/2

	if qleft >= mid+1 {
		return st.query(right, mid+1, r, qleft, qright)
	} else if qright <= mid {
		return st.query(left, l, mid, qleft, qright)
	}
	leftRes := st.query(left, l, mid, qleft, qright)
	rightRes := st.query(right, mid+1, r, qleft, qright)
	return leftRes.Merge(rightRes)
}

func (st *ArraySegmentTree) Set(index int, e MergerAble) {
	if index < 0 || index >= len(st.data) {
		panic("index of tree is illegal")
	}
	st.data[index] = e
	st.set(0, 0, len(st.data)-1, index, e)
}

func (st *ArraySegmentTree) set(tree, l, r, index int, e MergerAble) {
	if l == r {
		st.tree[tree] = e
		return
	}

	left := st.leftChild(tree)
	right := st.rightChild(tree)
	mid := l + (r-l)/2

	if index >= mid+1 {
		st.set(right, mid+1, r, index, e)
	} else {
		st.set(left, l, mid, index, e)
	}

	st.tree[tree] = st.tree[left].Merge(st.tree[right])
}

func (st *ArraySegmentTree) GetSize() int {
	return len(st.data)
}

func (st *ArraySegmentTree) leftChild(index int) int {
	return index*2 + 1
}

func (st *ArraySegmentTree) rightChild(index int) int {
	return index*2 + 2
}

func (st *ArraySegmentTree) String() string {
	return fmt.Sprintln(st.tree)
}
