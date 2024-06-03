
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

import "fmt"

type QuickUnionUnionFind struct {
	parent []int
	rank   []int
}

func CreateQuickUnionUnionFind(size int) *QuickUnionUnionFind {
	parent := make([]int, size)
	high := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		high[i] = 1
	}
	return &QuickUnionUnionFind{
		parent: parent,
		rank:   high,
	}
}

func (qf *QuickUnionUnionFind) find(x int) int {
	if x < 0 || x >= qf.GetSize() {
		panic("x is out of bound")
	}
	for x != qf.parent[x] {
		qf.parent[x] = qf.parent[qf.parent[x]]
		x = qf.parent[x]
	}
	return x
}

func (qf *QuickUnionUnionFind) Union(p, q int) {
	pId := qf.find(p)
	qId := qf.find(q)

	if pId == qId {
		return
	}
	if qf.rank[pId] < qf.rank[qId] {
		qf.parent[pId] = qId
	} else if qf.rank[pId] > qf.rank[qId] {
		qf.parent[qId] = pId
} else { //qf.等级[pid]==qf.等级[qid]
		qf.parent[qId] = pId
		qf.rank[pId]++
	}

}

func (qf *QuickUnionUnionFind) IsConnected(p, q int) bool {
	return qf.find(p) == qf.find(q)
}

func (qf *QuickUnionUnionFind) GetSize() int {
	return len(qf.parent)
}

func (qf *QuickUnionUnionFind) String() string {
	return fmt.Sprintln(qf.parent)
}
