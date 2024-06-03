
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

import "fmt"

type RedBlackTreeMap struct {
	root *rbTreeMapNode
	size int
}

const RED = true
const BLACK = false

type rbTreeMapNode struct {
	key         Comparable
	value       interface{}
	left, right *rbTreeMapNode
	color       bool
}

func CreateRedBlackTreeMap() *RedBlackTreeMap {
	return &RedBlackTreeMap{
		root: nil,
		size: 0,
	}
}

func (rbt *RedBlackTreeMap) isRed(node *rbTreeMapNode) bool {
	if node == nil {
		return BLACK
	}
	return node.color
}

func (rbt *RedBlackTreeMap) Add(k interface{}, v interface{}) {
	rbt.root = rbt.add(rbt.root, k, v)
	rbt.root.color = BLACK
}

func (rbt *RedBlackTreeMap) add(node *rbTreeMapNode, k interface{}, v interface{}) *rbTreeMapNode {
	if node == nil {
		rbt.size++
		return &rbTreeMapNode{k.(Comparable), v, nil, nil, RED}
	}

	if k.(Comparable).Compare(node.key) < 0 {
		node.left = rbt.add(node.left, k, v)
	} else if k.(Comparable).Compare(node.key) > 0 {
		node.right = rbt.add(node.right, k, v)
	} else {
		node.value = v
	}

if rbt.isRed(node.right) && !rbt.isRed(node.left) { //保持红黑质的三次操作
		node = rbt.leftRotate(node)
	}
	if rbt.isRed(node.right) && rbt.isRed(node.left.left) {
		node = rbt.rightRotate(node)
	}
	if rbt.isRed(node.right) && rbt.isRed(node.left) {
		rbt.flipColor(node)
	}
	return node
}

/*
  节点x node.right=x.left
        /\/\x.left=节点
       T1 x节点t3 x.color=节点.color
            /\/\node.color=red
           t2 t3 t1 t2


**/

func (rbt *RedBlackTreeMap) leftRotate(node *rbTreeMapNode) *rbTreeMapNode {
	x := node.right
	node.right = x.left
	x.left = node

	x.color = node.color
	node.color = RED
	return x
}

func (rbt *RedBlackTreeMap) rightRotate(node *rbTreeMapNode) *rbTreeMapNode {
	x := node.left
	node.left = x.right
	x.right = node

	x.color = node.color
	node.color = RED
	return x
}
func (rbt *RedBlackTreeMap) flipColor(node *rbTreeMapNode) {
	node.color = RED
	node.left.color = BLACK
	node.right.color = BLACK
}

func (rbt *RedBlackTreeMap) Remove(k interface{}) interface{} {
	ret := rbt.Get(k)
	rbt.root = rbt.remove(rbt.root, k)
	return ret
}
func (rbt *RedBlackTreeMap) remove(node *rbTreeMapNode, k interface{}) *rbTreeMapNode {
	if node == nil {
		return nil
	}
	if k.(Comparable).Compare(node.key) < 0 {
		node.left = rbt.remove(node.left, k)
		return node
	} else if k.(Comparable).Compare(node.key) > 0 {
		node.right = rbt.remove(node.right, k)
		return node
	} else {
		if node.left == nil {
			right := node.right
			node.right = nil
			rbt.size--
			return right
		} else if node.right == nil {
			left := node.left
			node.left = nil
			rbt.size--
			return left
		} else {
			ret := rbt.min(node.right)
			ret.right = rbt.remove(node.right, ret.key)
			ret.left = node.left
			node.left, node.right = nil, nil
			return ret
		}
} //todo完成删除后的平衡
}

func (rbt *RedBlackTreeMap) min(node *rbTreeMapNode) *rbTreeMapNode {
	if node.left == nil {
		return node
	}
	return rbt.min(node.left)
}
func (rbt *RedBlackTreeMap) max(node *rbTreeMapNode) *rbTreeMapNode {
	if node.right == nil {
		return node
	}
	return rbt.max(node.right)
}
func (rbt *RedBlackTreeMap) Contains(key interface{}) bool {
	return rbt.contains(rbt.root, key)
}
func (rbt *RedBlackTreeMap) contains(node *rbTreeMapNode, key interface{}) bool {
	if node == nil {
		return false
	} else if node.key == key {
		return true
	}
	if node.key.Compare(key.(Comparable)) < 0 {
		return rbt.contains(node.left, key)
	} else {
		return rbt.contains(node.right, key)
	}
}
func (rbt *RedBlackTreeMap) Get(key interface{}) interface{} {
	return rbt.get(rbt.root, key)
}
func (rbt *RedBlackTreeMap) get(node *rbTreeMapNode, key interface{}) interface{} {
	if rbt.size == 0 {
		return nil
	}
	if node == nil {
		return nil
	} else if node.key == key {
		return node.value
	}

	if key.(Comparable).Compare(node.key) < 0 {
		return rbt.get(node.left, key)
	} else {
		return rbt.get(node.right, key)
	}
}
func (rbt *RedBlackTreeMap) Set(key interface{}, newValue interface{}) {
	rbt.set(rbt.root, key, newValue)
}
func (rbt *RedBlackTreeMap) set(node *rbTreeMapNode, key interface{}, newValue interface{}) {
	if rbt.size == 0 {
		panic("bu cun zai zhe yang de key")
	}
	if node == nil {
		panic("bu cun zai zhe yang de key1")
	}
	if node.key == key {
		node.value = newValue
		return
	}

	if key.(Comparable).Compare(node.key) < 0 {
		rbt.set(node.left, key, newValue)
	} else {
		rbt.set(node.right, key, newValue)
	}
}
func (rbt *RedBlackTreeMap) GetSize() int {
	return rbt.size
}

func (rbt *RedBlackTreeMap) IsEmpty() bool {
	return rbt.size == 0
}

func (rbt *RedBlackTreeMap) String() string {
	str := ""
	return rbt.createString(rbt.root, 0, str)
}

func (rbt *RedBlackTreeMap) createString(node *rbTreeMapNode, depth int, str string) string {
	if node == nil {
		str += rbt.depthString(depth) + "nil \n"
		return str
	}

	str += rbt.depthString(depth) + "key:" + fmt.Sprint(node.key) + " value:" + fmt.Sprint(node.value) + "\n"
	str = rbt.createString(node.left, depth+1, str)
	str = rbt.createString(node.right, depth+1, str)
	return str
}
func (rbt *RedBlackTreeMap) depthString(depth int) string {
	str := fmt.Sprint("")
	for i := 0; i < depth; i++ {
		str += fmt.Sprint("--")
	}
	return str
}
