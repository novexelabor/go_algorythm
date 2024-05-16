package Calc

type Lexer struct {
	input        string //输入字符串
	position     int    //位置
	readposition int    //读取位置
	ch           byte   //读取一个字节
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input} //初始化
	l.ReadChar()              //前进一个字符
	return l
}

//分离数据，操作
func (l *Lexer) NextToken() Token {
	var tok Token
	l.SkipWhiteSpace() //跳过垃圾字符
	switch l.ch {
	case '(':
		tok = NewToken(LPAREN, l.ch)
	case ')':
		tok = NewToken(RPAREN, l.ch)
	case '+':
		tok = NewToken(PLUS, l.ch)
	case '-':
		tok = NewToken(MINIUS, l.ch)
	case '*':
		tok = NewToken(ASTERISK, l.ch)
	case '/':
		tok = NewToken(SLASH, l.ch)
	case '!':
		tok = NewToken(G0, l.ch)
	case '>':
		tok = NewToken(BIG, l.ch)
	//case '0':
	//tok.Literal=""
	//tok.Type=EOF //关闭
	default:
		if IsDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.ReadNumber()
			return tok
		} else {
			tok = NewToken(ILLEGAL, l.ch) //非法字符
		}

	}
	l.ReadChar()
	return tok

}

//提取一个字符
func (l *Lexer) ReadChar() {
	if l.readposition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readposition]
	}
	l.position = l.readposition
	l.readposition += 1
}

//123+2,切割数字出来
func (l *Lexer) ReadNumber() string {
	position := l.position //记录第一个数字的位置
	for IsDigit(l.ch) {
		l.ReadChar() //连续提取数字
	}
	return l.input[position:l.position] //返回连续的数字字符串
}

//跳过空格
func (l *Lexer) SkipWhiteSpace() {
	if l.ch == '\t' || l.ch == ' ' || l.ch == '\r' || l.ch == '\n' {
		l.ReadChar()
	}
}
