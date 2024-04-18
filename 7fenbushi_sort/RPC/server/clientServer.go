package main

import (
	"fmt"
	"net/rpc"
)

type ArgsX struct {
	A, B int //两个数据
}

type Query struct {
	X, Y int //两个数据
}

func main() {
	//severip:=":1234"
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println(err)
	}
	i1 := 13
	i2 := 5
	args := ArgsX{i1, i2}
	var reply int
	err = client.Call("Last.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(args.A, args.B, reply) ///乘法

	var qu Query
	err = client.Call("Last.Divide", args, &qu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(args.A, args.B, qu.X, qu.Y)

}
