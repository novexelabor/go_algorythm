
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

import "fmt"

type QuickFindUnionFind struct {
	id []int
}

func CreateQuickFindUnionFind(size int) *QuickFindUnionFind {
	arr := make([]int, size)
	for key := range arr {
		arr[key] = key
	}
	return &QuickFindUnionFind{
		id: arr,
	}
}

func (qf *QuickFindUnionFind) find(x int) int {
	if x < 0 || x >= qf.GetSize() {
		panic("x is out of bound")
	}
	return qf.id[x]
}

func (qf *QuickFindUnionFind) Union(p, q int) {
	pId := qf.find(p)
	qId := qf.find(q)

	if pId == qId {
		return
	}
//fmt.println（PID、QID）
	for key, value := range qf.id {
		if value == pId {
			qf.id[key] = qId
		}
	}
}

func (qf *QuickFindUnionFind) IsConnected(p, q int) bool {
	return qf.find(p) == qf.find(q)
}

func (qf *QuickFindUnionFind) GetSize() int {
	return len(qf.id)
}

func (qf *QuickFindUnionFind) String() string {
	return fmt.Sprintln(qf.id)
}
