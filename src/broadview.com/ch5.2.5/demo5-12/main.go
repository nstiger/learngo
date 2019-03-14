package main

import (
	"fmt"
)

func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	//nextNumber为一个函数，函数 i 为 0
	nextNumber := getSequence()

	//调用nextNumber 函数， i变量自增1并返回
	fmt.Printf("%d ", nextNumber())
	fmt.Printf("%d ", nextNumber())
	fmt.Printf("%d ", nextNumber())

	//创建新的函数 nextNumber1, 并查看结果
	nextNumber1 := getSequence()
	fmt.Printf("%d ", nextNumber1())
	fmt.Printf("%d ", nextNumber1())
}