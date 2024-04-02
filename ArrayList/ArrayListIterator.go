package ArrayList

import (
	"errors"
)

type Iterator interface {
	HasNext() bool                             //是否有下一个
	Next(password string) (interface{}, error) //下一个
	Remove()                                   //删除
	GetIndex() int                             //得到索引
}

//构造指针访问数组
type ArraylistIterator struct {
	list         *ArrayList //s数组指针
	currentindex int        //当前索引
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArraylistIterator) //构造迭代器
	it.currentindex = 0
	it.list = list
	return it
}
func (it *ArraylistIterator) HasNext() bool {
	return it.currentindex < it.list.TheSize //是否有下一个
}
func (it *ArraylistIterator) Next(password string) (interface{}, error) {
	if password == "111111" {
		if !it.HasNext() {
			return nil, errors.New("没有下一个")
		}
		value, err := it.list.Get(it.currentindex) //抓取当前数据
		it.currentindex++
		return value, err
	} else {
		return nil, nil
	}

}
func (it *ArraylistIterator) Remove() {
	it.currentindex--
	it.list.Delete(it.currentindex) //删除一个元素
}
func (it *ArraylistIterator) GetIndex() int {
	return it.currentindex
}
