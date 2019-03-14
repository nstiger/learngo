package main
import (
	"fmt"
)

func main() {
	var f = adder()
	//相当于fmt.Println(adder()(1))
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}

func adder() func(int) int {
	var x int //闭包中的变量可以在闭包函数体内声明，也可以在外部函数声明
	return func(d int) int {
		fmt.Println("x = ", x,  "d = ", d)
		x += d
		return x
	}
}