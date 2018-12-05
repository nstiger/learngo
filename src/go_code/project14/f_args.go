package main
import (
	"fmt"
)

//函数可变形参
func sum(n1 int, args... int)(res int) {
	res = n1
	for i :=0; i < len(args); i++ {
		res += args[i]
	}
	return res
}

func main() {
	// var a int
	// a = sum(10, 20, 30, -45)
	a := sum(10, 20, 30, -45)
	fmt.Println(a)
}