package main

import (
	"fmt"
)

var i = 0

func print() {
	fmt.Println(i)
}

func main() {
	for ; i < 5; i++ {
		defer print()
	}
}

/*
输出结果：
5
5
5
5
5
*/