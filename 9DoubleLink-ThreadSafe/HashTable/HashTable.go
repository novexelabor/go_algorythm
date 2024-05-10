package main

import (
	"fmt"
	"math"

	"errors"
)

//哈希表
type HashTable struct {
	Table map[int]*List //映射
	size  int           //大小
	cap   int           //容量
}

//元素
type Item struct {
	key   string
	value interface{}
}

//返回哈希表节点
func NewHashTable(cap int) *HashTable {
	table := make(map[int]*List, cap) //初始化
	return &HashTable{table, 0, cap}
}
func (ht *HashTable) Get(key string) (interface{}, error) {
	index := ht.pos(key)             //索引,map的key
	item, err := ht.find(index, key) //查找数据
	if item == nil {
		return "", errors.New("没有找到")
	}
	return item.value, err
}
func (ht *HashTable) Put(key, value string) {
	index := ht.pos(key) //索引
	if ht.Table[index] == nil {
		ht.Table[index] = NewList() //新建一个节点

	}
	item := &Item{key, value} //新建插入的对象
	data, err := ht.find(index, key)
	if err != nil { //没有找到
		ht.Table[index].Append(item) //数据插入
		ht.size++
	} else {
		data.value = value //替换
	}

}

func (ht *HashTable) Del(key string) error {
	index := ht.pos(key)
	mylist := ht.Table[index] //定位要删除的链表
	var val *Item
	mylist.Each(func(node Node) {
		if node.Value.(*Item).key == key {
			val = node.Value.(*Item) //取出数据
		}
	})
	if val == nil {
		return nil //返回数据
	}
	ht.size--
	return mylist.Remove(val) //s删除

}

//循环哈希表的多个链表，循环每个链表的元素
func (ht *HashTable) Foreach(f func(item *Item)) {
	for k := range ht.Table {
		if ht.Table[k] != nil {
			ht.Table[k].Each(func(node Node) {
				f(node.Value.(*Item))
			})
		}
	}
}
func (ht *HashTable) pos(s string) int {
	return hashCode(s) % ht.cap //根据哈希值计算,使用哈希值取余计算
}
func (ht *HashTable) find(i int, key string) (*Item, error) {
	mylist := ht.Table[i] //每一个哈希值对应一个链表
	var val *Item
	mylist.Each(func(node Node) {
		if node.Value.(*Item).key == key {
			val = node.Value.(*Item) //取出数据
		}
	})
	if val == nil {
		return nil, errors.New("not find")
	}
	return val, nil
}

//根据字符串计算哈希
func hashCode(str string) int {
	hash := int32(0)
	for i := 0; i < len(str); i++ {
		hash = int32(hash<<5-hash) + int32(str[i]) //计算形式,可以调用库函数
		hash &= hash                               //哈希计算
	}
	return int(math.Abs(float64(hash)))
}

func main() {
	ht := NewHashTable(1000)
	ht.Put("yincheng1", "123456")
	ht.Put("yincheng2", "1234567")
	ht.Put("yincheng3", "1234568")
	ht.Put("yincheng4", "1234569")
	ht.Put("yincheng5", "1234560")
	fmt.Println(ht.Table)
	ht.Put("yincheng3", "abc")
	ht.Del("yincheng3")
	fmt.Println(ht.Table)
	val, err := ht.Get("yincheng3")
	fmt.Println(val, err)

}

//-------------------------------------
type List struct { //链表
	Length int
	Head   *Node
	Tail   *Node
}

func NewList() *List {
	l := new(List)
	l.Length = 0
	return l
}

type Node struct {
	Value interface{}
	Prev  *Node
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{Value: value}
}

func (l *List) Len() int {
	return l.Length
}

func (l *List) IsEmpty() bool {
	return l.Length == 0
}

//插入在head之前
func (l *List) Prepend(value interface{}) {
	node := NewNode(value)
	if l.Len() == 0 {
		l.Head = node
		l.Tail = l.Head
	} else {
		formerHead := l.Head
		formerHead.Prev = node

		node.Next = formerHead
		l.Head = node
	}

	l.Length++
}

func (l *List) Append(value interface{}) {
	node := NewNode(value)

	if l.Len() == 0 { //只有一个的时候
		l.Head = node
		l.Tail = l.Head
	} else {
		formerTail := l.Tail
		formerTail.Next = node

		node.Prev = formerTail
		l.Tail = node
	}

	l.Length++
}

func (l *List) Add(value interface{}, index int) error {
	if index > l.Len() {
		return errors.New("index out of range")
	}

	node := NewNode(value)

	if l.Len() == 0 || index == 0 {
		l.Prepend(value)
		return nil
	}

	if l.Len()-1 == index {
		l.Append(value)
		return nil
	}

	nextNode, _ := l.Get(index)
	prevNode := nextNode.Prev

	prevNode.Next = node
	node.Prev = prevNode

	nextNode.Prev = node
	node.Next = nextNode

	l.Length++

	return nil
}

func (l *List) Remove(value interface{}) error {
	if l.Len() == 0 {
		return errors.New("empty list")
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Length--
		return nil
	}

	found := 0
	for n := l.Head; n != nil; n = n.Next {

		if n.Value == value && found == 0 {
			n.Next.Prev, n.Prev.Next = n.Prev, n.Next
			l.Length--
			found++
		}
	}

	if found == 0 {
		return errors.New("Node not found")
	}

	return nil
}

func (l *List) Get(index int) (*Node, error) {
	if index > l.Len() {
		return nil, errors.New("index out of range")
	}

	node := l.Head
	for i := 1; i < index; i++ {
		node = node.Next
	}

	return node, nil
}

func (l *List) Find(node *Node) (int, error) { //找到的是下标
	if l.Len() == 0 {
		return 0, errors.New("empty list")
	}

	index := 0
	found := -1
	l.Map(func(n *Node) {
		index++
		if n.Value == node.Value && found == -1 {
			found = index
		}
	})

	if found == -1 {
		return 0, errors.New("Item not found")
	}

	return found, nil
}

func (l *List) Clear() {
	l.Length = 0
	l.Head = nil
	l.Tail = nil
}

func (l *List) Concat(k *List) {
	l.Tail.Next, k.Head.Prev = k.Head, l.Tail
	l.Tail = k.Tail
	l.Length += k.Length
}

func (list *List) Map(f func(node *Node)) {
	for node := list.Head; node != nil; node = node.Next {
		//n := node.Value.(*Node)
		f(node)
	}
}

//函数作为参数
func (list *List) Each(f func(node Node)) {
	for node := list.Head; node != nil; node = node.Next {
		f(*node)
	}
}
