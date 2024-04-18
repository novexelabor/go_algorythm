package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int //两个数据
}

type Query struct {
	X, Y int //两个数据
}
type Last int

func (t *Last) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B //乘法
	fmt.Println(reply, "乘法执行了")
	return nil
}
func (t *Last) Divide(args *Args, query *Query) error {
	if args.B == 0 {
		return errors.New("不能除以0")
	}
	query.X = args.A / args.B
	query.Y = args.A % args.B
	fmt.Println(query, "除法执行了")
	return nil
}

func main() {
	la := new(Last)
	fmt.Println(la, "=la")
	rpc.Register(la) //注册类型
	rpc.HandleHTTP() //设定http类型
	//err:=http.ListenAndServe(":1234",nil)
	list, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	http.Serve(list, nil) //4个步骤

}
