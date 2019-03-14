/*
匿名函数
func(x, y int) int {return x + y}

通过变量名对函数进行调用
fplus := func(x, y int) int {return x + y }
fplus(3, 4)

也可以后面紧跟运行参数
func(x, y int) int {return x + y } (3, 4)

*/
package main

import (
	"fmt"
)

func main() {
	func (num int) int {
		sum := 0
		for i := 1; i <= num; i++ {
			sum += i
		}
		fmt.Println(sum)
		return sum
	} (100)
}
