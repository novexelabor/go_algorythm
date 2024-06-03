//说明 ： 该代码是不对的，是为了书写测试用例

package main

import (
	"testing"
)

func Test_Split(t *testing.T) { //测试函数的写法
	// 初始化BtreeNode
	parent := NewBtreeNode(3, 3, false)
	btreetest := &Btree{parent, 3}
	for i := 0; i < 6; i++ {
		btreetest.Insert(i)
	}

	// 执行Split操作
	idx := 2
	parent.Split(3, idx)

	// 验证结果
	if parent.N != 4 {
		t.Errorf("parent.N = %d; want 4", parent.N)
	}
	if parent.keys[idx] != 2 {
		t.Errorf("parent.keys[%d] = %d; want 2", idx, parent.keys[idx])
	}
	if parent.Children[idx+1].N != 3 {
		t.Errorf("parent.Children[%d+1].N = %d; want 3", idx, parent.Children[idx+1].N)
	}
	if parent.Children[idx].N != 2 {
		t.Errorf("parent.Children[%d].N = %d; want 2", idx, parent.Children[idx].N)
	}
}
