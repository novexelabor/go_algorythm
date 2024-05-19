package main

import "fmt"

//1+(2-1*3)+(1+2)*(1+(1+(2-1)))
//栈，二叉树
func main() {
	op, err := NewOperator("1+2+(1+2*(1+(1+1)))")
	fmt.Println(op.opers)
	fmt.Println(op.suffixExpression)
	if err != nil {
		fmt.Println(err)
	}
	value, err := op.Execute([]string{})
	fmt.Println(value)
}

func main3() {
	op, err := NewOperator("I+a")
	if err != nil {
		fmt.Println(err)
	}
	value, err := op.Execute([]string{"I", "12", "a", "18"})
	fmt.Println(value)
}

func main2() {
	//op,err:=NewOperator("1*4-2")
	op, err := NewOperator("1+3+2")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(op.GetOpers())
	for i := 0; i < len(op.opers); i++ {
		fmt.Println(op.opers[i])
	}
	fmt.Println("----------------------------")
	fmt.Println(op.suffixExpression)

}
func main1() {
	mystack := NewStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	for mystack.Len() != 0 {
		fmt.Println(mystack.Pop())
	}
}
