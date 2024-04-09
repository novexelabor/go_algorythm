package main

import "fmt"

type Set struct {
	buf  []interface{}        //存储数据
	num  int                  //数量
	hash map[interface{}]bool //借助map实现映射
}

//新建一个可以变长的set
func NewSet() *Set {
	return &Set{make([]interface{}, 0), 0, make(map[interface{}]bool)}
}

func (this *Set) Add(value interface{}) bool {
	if this.isExist(value) {
		return false
	} else {
		this.buf = append(this.buf, value) //追加
		this.hash[value] = true
		this.num++
		return true
	}
}
func (this *Set) isExist(value interface{}) bool {
	return this.hash[value]
}
func (this *Set) Strings() []interface{} {
	return this.buf
}
func main() {
	set := NewSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	fmt.Println(set)
}
