package ArrayList

type StackArray interface {
	Clear()                //清空
	Size() int             //大小
	Pop() interface{}      //弹出
	Push(data interface{}) //压入
	IsFull() bool          //是否满了
	IsEmpty() bool         //是否为空
}
type Stack struct {
	myarray *ArrayList
	capsize int //最大范围

}

func NewArrayListStack() *Stack {
	mystack := new(Stack)
	mystack.myarray = NewArrayList() //数组
	mystack.capsize = 10             //空间
	return mystack

}
func (mystack *Stack) Clear() {
	mystack.myarray.Clear()
	mystack.capsize = 10 //空间
}
func (mystack *Stack) Size() int {
	return mystack.myarray.TheSize
}
func (mystack *Stack) Pop() interface{} {
	if !mystack.IsEmpty() {
		last := mystack.myarray.dataStore[mystack.myarray.TheSize-1]
		mystack.myarray.Delete(mystack.myarray.TheSize - 1)
		return last

	}
	return nil
}
func (mystack *Stack) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.myarray.Append(data)
	}
}
func (mystack *Stack) IsFull() bool { //判断满了
	if mystack.myarray.TheSize >= mystack.capsize {
		return true
	} else {
		return false
	}
}
func (mystack *Stack) IsEmpty() bool { //判断为空
	if mystack.myarray.TheSize == 0 {
		return true
	} else {
		return false
	}
}
