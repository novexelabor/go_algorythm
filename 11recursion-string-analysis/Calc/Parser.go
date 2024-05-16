package Calc

import (
	"fmt"
	"strconv"
)

type (
	prefixParseFn func() Expression
	infixParseFn  func(Expression) Expression
)
type Parser struct {
	l *Lexer

	curToken  Token //当前的
	peekToken Token //提取

	//解析
	prefixParseFns map[string]func() Expression
	infixParseFns  map[string]func(Expression) Expression

	//处理错误
	errors []string
}

//map插入数据
func (p *Parser) RigisterPrefix(tokentype string, fn func() Expression) {
	p.prefixParseFns[tokentype] = fn
}
func (p *Parser) RigisterInfix(tokentype string, fn func(Expression) Expression) {
	p.infixParseFns[tokentype] = fn
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.prefixParseFns = make(map[string]func() Expression)
	p.RigisterPrefix(INT, p.ParseIntergerLiteral)
	p.RigisterPrefix(MINIUS, p.ParsePrefixExpression)
	p.RigisterPrefix(PLUS, p.ParsePrefixExpression)
	p.RigisterPrefix(LPAREN, p.ParseGroupExpression)
	p.RigisterPrefix(G0, p.ParsePrefixExpression)

	p.infixParseFns = make(map[string]func(Expression) Expression)
	//p.RigisterInfix(PLUS,p.pa)
	p.RigisterInfix(PLUS, p.ParseInfixExpression)
	p.RigisterInfix(MINIUS, p.ParseInfixExpression)
	p.RigisterInfix(SLASH, p.ParseInfixExpression)
	p.RigisterInfix(ASTERISK, p.ParseInfixExpression)
	p.RigisterInfix(MOD, p.ParseInfixExpression)
	p.RigisterInfix(BIG, p.ParseInfixExpression)

	p.nextToken()
	p.nextToken()
	return p
}
func (p *Parser) Errors() []string {
	return p.errors //返回错误
}
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken() //循环截取数据处理
}
func (p *Parser) peekError(t string) { //处理解析错误
	msg := fmt.Sprintf("bug %s %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

//取出优先级
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	} else {
		return LOWEST
	}
}

//当前优先级
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	} else {
		return LOWEST
	}
}

//提取判断
func (p *Parser) peekTokenis(t string) bool {
	return p.peekToken.Type == t
}

//处理运算符
func (p *Parser) expectPeek(t string) bool {
	if p.peekTokenis(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
func (p *Parser) ParseExpression(precedence int) Expression {
	prefix := p.prefixParseFns[p.curToken.Type] //计算函数
	returnExp := prefix()                       //相当于一个数了，前缀的数

	for precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type] //前缀
		if infix == nil {
			return returnExp
		}
		p.nextToken()
		returnExp = infix(returnExp) //循环取出数据与函数
	}

	return returnExp
}

//1+  -1
func (p *Parser) ParsePrefixExpression() Expression {
	expression := &PrefixExpression{Token: p.curToken, Operator: p.curToken.Literal}
	p.nextToken()
	expression.Right = p.ParseExpression(PREFIX) //处理右边
	return expression
}

//解析数据
func (p *Parser) ParseIntergerLiteral() Expression {
	lit := &IntergerLiteralExpression{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64) //解析数据
	if err != nil {
		msg := fmt.Sprintf("bug %s int %s", "", p.peekToken.Type)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}
func (p *Parser) ParseGroupExpression() Expression {
	p.nextToken()
	exp := p.ParseExpression(LOWEST) //设定权限
	if !p.expectPeek(RPAREN) {
		return nil
	}
	return exp
}

func (p *Parser) ParseInfixExpression(left Expression) Expression {
	expresssion := &InfixExpression{Token: p.curToken, operator: p.curToken.Literal, Left: left}
	precedence := p.curPrecedence()
	p.nextToken() //循环继续
	expresssion.Right = p.ParseExpression(precedence)
	return expresssion

}
