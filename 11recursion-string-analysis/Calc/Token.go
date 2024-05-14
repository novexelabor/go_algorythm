package Calc

//处理运算符
type Token struct {
	Type    string //类型
	Literal string //意义
}

func NewToken(tokentype string, c byte) Token {
	return Token{tokentype, string(c)}
}
