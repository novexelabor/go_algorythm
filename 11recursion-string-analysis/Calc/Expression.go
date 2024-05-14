package Calc

import (
	"bytes"
)

//sin x  cosy
//接口
type Expression interface {
	String() string
}

//整数求值
type IntergerLiteralExpression struct {
	Token Token
	Value int64
}

func (il *IntergerLiteralExpression) String() string {
	return il.Token.Literal
}

//前缀 1+ -1
type PrefixExpression struct {
	Token    Token
	Operator string
	Right    Expression
}

//括号内部计算(+1)
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    Token
	Left     Expression
	operator string
	Right    Expression
}

//1+2  (1+2)
func (in *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(in.Left.String())
	out.WriteString(" ")
	out.WriteString(in.operator)
	out.WriteString(" ")
	out.WriteString(in.Right.String())
	out.WriteString(")")
	return out.String()
}
