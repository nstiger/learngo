/*
函数作为参数
只要被调用的函数的返回值个数、返回值类型和返回值的顺序与调用函数所需求的实参是一致的，就可以
把这个被调用函数当作其他函数的参数
*/

package main

import (
	"fmt"
)

func pipe(ff func() int) int {
	return ff()
}

//formatFunc 定义函数类型
type FormatFunc func(s string, x, y int) string

func format(ff FormatFunc, s string, x, y int) string {
	return ff(s, x, y)
}

func main() {
	s1 := pipe(func() int { return 100})	//直接将匿名函数当参数
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)
	fmt.Println(s1, s2)
}