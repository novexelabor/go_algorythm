package main

import (
	"fmt"
	"go_algorythm/11recursion-string-analysis/Calc"
)

//sin  cos,  log
func main() {

	//fmt.Println(Calc.Calc("1+2*(1+(1+1))+2*(3-1)+-1"))
	fmt.Println(Calc.Calc("1+!1+!2+!3"))
	fmt.Println(Calc.Calc("!1"))
	fmt.Println(Calc.Calc("!0"))
	fmt.Println(Calc.Calc("1+2*!0+3*!1"))
	fmt.Println(Calc.Calc("3>2"))
	fmt.Println(Calc.Calc("3>12"))
	fmt.Println(Calc.Calc("3+!0>5"))
}
