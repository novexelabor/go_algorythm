package Calc

const (
	ILLEGAL = "ILLEGAL" //非法字符串
	EOF     = "EOF"     //终止
	INT     = "INT"     //整数

	PLUS     = "+" //加减乘除
	MINIUS   = "-"
	ASTERISK = "*"
	SLASH    = "/"
	MOD      = "%"

	BIG = ">" //5>1 return 1  1>5 0
	G0  = "!" //!3344=0

	LPAREN = "(" //括号
	RPAREN = ")"
)

const (
	_      int = iota //优先级
	LOWEST            //级别 （）
	BIGGER
	SUM     //+ -
	PRODUCT //* /
	PREFIX  // -1
	CALL    //1+(3)

)

//构造集合实现优先级
var precedences = map[string]int{
	PLUS:     SUM,
	MINIUS:   SUM,
	SLASH:    PRODUCT,
	ASTERISK: PRODUCT,
	MOD:      PRODUCT,
	LPAREN:   CALL,
	G0:       PREFIX,
	BIG:      BIGGER,
}
