package main

import (
	"log"
	"net/rpc"
	"fmt"
	"os"

)

type Args struct{
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("用法：", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddress+ ":1234")
	if err != nil {
		log.Fatal("Dial错误：", err)
	}
	//Synchronous call
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith错误：", err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith错误：", err)
	}
	fmt.Printf("Arith: %d / %d = %d 余 %d\n", args.A, args.B, quot.Quo, quot.Rem)

}