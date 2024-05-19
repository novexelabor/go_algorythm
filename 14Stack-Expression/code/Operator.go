package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	//"fmt"
	//"fmt"
)

//（a+1）*b
//a=1,b=2
//4
const (
	//四则运算符号
	AVAIABLE_CODE = "+-*/()^√"
	//小数点符号
	AVAIABLE_DECIMAL_CODE = "1234567890.E"
	//参数符号
	AVAILABLE_PARAMETER_CODE = "abcdefghijklmnopqrstuvwxyz"
)

//判断计算符
func isAvaiablecode(c string) bool {
	return strings.IndexAny(AVAIABLE_CODE, c) != -1
}

//判断小数
func isBelongToDecimal(c string) bool {
	return strings.IndexAny(AVAIABLE_DECIMAL_CODE, c) != -1
}

//判断参数
func isParmertercode(c string) bool {
	return strings.IndexAny(AVAILABLE_PARAMETER_CODE, c) != -1
}

//计算
func exesingleExpression(left float64, right float64, exp string) float64 {
	if exp == "+" {
		return left + right
	} else if exp == "-" {
		return left - right
	} else if exp == "*" {
		return left * right
	} else if exp == "/" {
		return left / right
	} else if exp == "^" {
		return math.Pow(left, right)
	} else if exp == "√" {
		return math.Pow(left, 1/right)
	}
	return 0.0
}

//(1+a)*b   a=1  b=2
//参数替换
//[]string{"I","12","a","18"}  ---str是替换参数的切片
func changeParameter(Parameter string, str []string) string {
	for i := 0; i < len(str); i += 2 {
		if Parameter == str[i] {
			return str[i+1]
		}
	}
	return Parameter
}

//四则运算类
type Operator struct {
	sentence         string   //1+(2*3)//文字表达式
	opers            []string //表达式存储，  2*3  1+6
	suffixExpression []string //后缀表达式

}

//返回数据
func (ths *Operator) GetOpers() []string {
	return ths.opers
}
func (ths *Operator) GesuffixExpression() []string {
	return ths.suffixExpression
}

//新建一个四则运算类
func NewOperator(sentence string) (*Operator, error) {
	o := &Operator{sentence,
		make([]string, 0, len(sentence)),
		make([]string, 0, len(sentence))} //初始化分配内存
	o.init()                       //初始化
	err := o.setSuffixExpression() //处理后缀表达式
	return o, err

}

//初始化,字符串切割，放入opers   1  +    2
func (ths *Operator) init() {
	//type rune = int32\
	//123
	value := make([]rune, 0, len(ths.sentence)) //定义数组开辟内存,处理数字
	flag := false                               //标识

	for _, c := range ths.sentence { //循环每一个字符
		if isBelongToDecimal(string(c)) { //处理数字
			value = append(value, c)
			flag = true
			if c == rune('E') {
				flag = false
			}

		} else if (c == rune('-') || c == rune('+')) && !flag {
			value = append(value, c)
			flag = true
		} else if isAvaiablecode(string(c)) {
			if flag && len(value) > 0 {
				ths.opers = append(ths.opers, string(value)) //追加计算符号
				flag = false
			}
			value = value[:0]                        //清空
			ths.opers = append(ths.opers, string(c)) //加入字符

		} else if isParmertercode(string(c)) {
			if flag && len(value) > 0 {
				ths.opers = append(ths.opers, string(value)) //追加计算符号

			}
			value = value[:0]                        //清空
			ths.opers = append(ths.opers, string(c)) //加入字符
			flag = true

		} else if c == rune('I') {
			if flag && len(value) > 0 {
				ths.opers = append(ths.opers, string(value)) //追加计算符号

			}
			value = value[:0]                        //清空
			ths.opers = append(ths.opers, string(c)) //加入字符
			flag = false
		}
	}
	//处理最后一段
	if flag && len(value) > 0 {
		ths.opers = append(ths.opers, string(value))
	}

}

//中缀表达式转化为后缀表达式，自动具备了顺序
//1+（2*3）中缀表达式
//后缀表达式，解决优先级问题
//(1+2)  1+2
//将opers中存放的表达式转化成后缀形式,存放至suffixExpression中

func (ths *Operator) setSuffixExpression() error {
	op := NewStack()    //新建一个栈
	var top interface{} //任何类型，数字，运算符
	for i := 0; i < len(ths.opers); i++ {
		cur := ths.opers[i] //获得当前字符
		if cur == "(" {
			op.Push(cur) //压栈
		} else if cur == ")" {
			for {
				top = op.Pop() //弹出数据
				if top == "(" {
					break
				}
				ths.suffixExpression = append(ths.suffixExpression, M2string(top))
				//括号之间的数据截取
			}

		} else if cur == "+" || cur == "-" || cur == "*" || cur == "/" || cur == "^" || cur == "√" {
			for {
				if op.Empty() || op.Peak() == "(" || ((cur == "*" || cur == "/") && (op.Peak() == "+" || op.Peak() == "-")) || ((cur == "^" || cur == "√") && (op.Peak() == "*" || op.Peak() == "/" || op.Peak() == "+" || op.Peak() == "-")) {
					op.Push(cur)
					break
				} else {
					top = op.Pop()
					ths.suffixExpression = append(ths.suffixExpression, M2string(top))
				}
			}
		} else {
			ths.suffixExpression = append(ths.suffixExpression, M2string(cur))
		}

	}
	for {
		if op.Empty() { //如果栈为空，跳出循环
			break
		}
		top := op.Pop()
		if top != "(" {
			ths.suffixExpression = append(ths.suffixExpression, M2string(top))
		} else {
			return errors.New("错误表达式")
		}
	}

	return nil
}

//计算结果
func (ths *Operator) Execute(str []string) (value float64, err error) {
	temp := NewStack()
	for i := 0; i < len(ths.suffixExpression); i++ {
		st := changeParameter(ths.suffixExpression[i], str) //替换数据
		if val, err := strconv.ParseFloat(strings.TrimSpace(st), 64); err == nil {
			temp.Push(val) //压入数据
		} else {
			exp := ths.suffixExpression[i] //取出数据
			if exp == "I" {
				v1 := temp.Pop()       //取得弹出数据
				temp.Push(M2int64(v1)) //压入数据.计算好的结果
			} else {
				rights := temp.Pop()       //取出数据  123
				right := M2float64(rights) //数据转化
				lefts := temp.Pop()
				left := M2float64(lefts)
				temp.Push(exesingleExpression(left, right, exp)) //递归调用
			}
		}
	}
	value = M2float64(temp.Pop())
	return
}
