package ArrayList

type StackArrayX interface {
	Clear()                //清空
	Size() int             //大小
	Pop() interface{}      //弹出
	Push(data interface{}) //压入
	IsFull() bool          //是否满了
	IsEmpty() bool         //是否为空
}
type StackX struct {
	myarray *ArrayList
	Myit    Iterator
}

func NewArrayListStackX() *StackX {
	mystack := new(StackX)
	mystack.myarray = NewArrayList()          //数组
	mystack.Myit = mystack.myarray.Iterator() //迭代
	return mystack

}
func (mystack *StackX) Clear() {
	mystack.myarray.Clear()
	mystack.myarray.TheSize = 0
}
func (mystack *StackX) Size() int {
	return mystack.myarray.TheSize
}
func (mystack *StackX) Pop() interface{} {
	if !mystack.IsEmpty() { //从TheSize这端出
		last := mystack.myarray.dataStore[mystack.myarray.TheSize-1]
		mystack.myarray.Delete(mystack.myarray.TheSize - 1)
		return last

	}
	return nil
}
func (mystack *StackX) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.myarray.Append(data)
	}
}
func (mystack *StackX) IsFull() bool { //判断满了
	if mystack.myarray.TheSize >= 10 {
		return true
	} else {
		return false
	}
}
func (mystack *StackX) IsEmpty() bool { //判断为空
	if mystack.myarray.TheSize == 0 {
		return true
	} else {
		return false
	}
}
