package main

import "fmt"

type set struct {
	mylist *List //内部实现基于链表
}

//根据索引返回数据
func (myset *set) GetAt(i uint64) Object {
	return (*myset).mylist.GetAt(i)
}

//返回集合大小
func (myset *set) GetSize() uint64 {
	return (*myset).mylist.GetSize()
}

//初始化
func (myset *set) Init(match ...MatchFun) {
	mylist := new(List)      //新的链表
	(*myset).mylist = mylist //初始化
	if len(match) == 0 {
		mylist.Init() //初始化
	} else {
		mylist.Init(match[0]) //初始化序列
	}
}

//判断数据是否存在
func (myset *set) Isin(data Object) bool {
	return (*myset).mylist.IsMember(data)
}

//插入
func (myset *set) Insert(data Object) bool {
	if !myset.Isin(data) {
		return (*myset).mylist.Append(data)
	}
	return false
}

//判断是否为空
func (myset *set) IsEmpty() bool {
	return (*myset).mylist.IsEmpty()
}

//删除
func (myset *set) Remove(data Object) bool {
	return (*myset).mylist.Remove(data)
}

//1,2,3
//1,4,5
func (myset *set) Union(set1 *set) *set {
	if set1 == nil {
		return myset
	}
	if myset == nil {
		return set1
	}
	//集合新建的集合存储新的结构
	nset := new(set)
	nset.Init((*((*myset).mylist)).myMatch)
	if myset.IsEmpty() && set1.IsEmpty() {
		return nset //空集合
	}
	for i := uint64(0); i < myset.GetSize(); i++ {
		nset.Insert(myset.GetAt(i)) //插入数据
	}
	var data Object //判断set1的数据在myset是否存在，存在就不管，不存在插入
	for i := uint64(0); i < set1.GetSize(); i++ {
		data = set1.GetAt(i)
		if !nset.Isin(data) {
			nset.Insert(data)
		}
	}
	return nset

}

//123
//235   //23
func (myset *set) Share(set1 *set) *set {
	if set1 == nil {
		return nil
	}
	if myset == nil {
		return nil
	}
	nset := new(set)
	nset.Init((*((*myset).mylist)).myMatch)
	if myset.IsEmpty() && set1.IsEmpty() {
		return nset //空集合
	}
	largeset := myset //保存最多元素
	smallset := set1  //保存较小元素
	if set1.GetSize() > myset.GetSize() {
		largeset = set1
		smallset = myset
	}
	var data Object
	for i := uint64(0); i < largeset.GetSize(); i++ {
		data = largeset.GetAt(i) //保存两者都有的元素

		if smallset.Isin(data) {
			nset.Insert(data)
		}
	}
	return nset

}
func (myset *set) Different(set1 *set) *set {
	if set1 == nil {
		return nil
	}
	if myset == nil {
		return nil
	}
	nset := new(set)
	nset.Init((*((*myset).mylist)).myMatch)
	if myset.IsEmpty() && set1.IsEmpty() {
		return nset //空集合
	}

	var data Object
	for i := uint64(0); i < myset.GetSize(); i++ {
		data = myset.GetAt(i) //保存myset有，而set1没有

		if !set1.Isin(data) {
			nset.Insert(data)
		}
	}
	return nset

}

func (myset *set) IsSub(subset *set) bool {
	if myset == nil {
		return false
	}
	if subset == nil {
		return true
	}
	for i := uint64(0); i < subset.GetSize(); i++ {
		if !myset.Isin(subset.GetAt(i)) { //有一个不存在就不是子集
			return false
		}
	}
	return true
}

func (myset *set) IsEquals(subset *set) bool {
	if myset == nil || subset == nil {
		return false
	}
	if myset == nil && subset == nil {
		return true
	}
	nset := myset.Share(subset)              //两方都有
	return nset.GetSize() == myset.GetSize() //集合相等
}

type SetIterator struct {
	index uint64 //索引
	myset *set   //集合
}

//新建一个迭代器
func (myset *set) GetIterator() *SetIterator {
	it := new(SetIterator)
	(*it).index = 0
	(*it).myset = myset
	return it
}
func (it *SetIterator) HashNext() bool {
	set := (*it).myset
	index := (*it).index
	return index < set.GetSize() //判断是否有下一个
}
func (it *SetIterator) Next() Object {
	set := (*it).myset
	index := (*it).index //根据集合与索引
	if index < set.GetSize() {
		data := set.GetAt(index) //取出数据
		(*it).index++
		return data
	}
	return nil
}

func match(data1 Object, data2 Object) int {
	if data1 == data2 {

		return 0
	} else {
		return 1
	}
}

func main() {

	myset1 := new(set)
	myset1.Init(match)
	myset2 := new(set)
	myset2.Init(match)
	myset1.Insert(1)
	myset1.Insert(2)
	myset1.Insert(3)

	for it := myset1.GetIterator(); it.HashNext(); {
		fmt.Println(it.Next())
	}
	fmt.Println("-------------------------")

	myset2.Insert(4)
	myset2.Insert(2)
	myset2.Insert(3)
	for it := myset2.GetIterator(); it.HashNext(); {
		fmt.Println(it.Next())
	}
	fmt.Println("-------------------------")
	myset3 := myset1.Union(myset2)
	for it := myset3.GetIterator(); it.HashNext(); {
		fmt.Println(it.Next())
	}
	fmt.Println("-------------------------")
	myset4 := myset1.Share(myset2)
	for it := myset4.GetIterator(); it.HashNext(); {
		fmt.Println(it.Next())
	}
	fmt.Println("-------------------------")
	myset5 := myset1.Different(myset2)
	for it := myset5.GetIterator(); it.HashNext(); {
		fmt.Println(it.Next())
	}
	fmt.Println(myset1.IsEquals(myset2))
	fmt.Println(myset1.IsSub(myset2))
}
