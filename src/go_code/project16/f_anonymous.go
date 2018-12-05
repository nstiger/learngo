package main
import (
	"fmt"
)
//匿名函数
func main() {

	var n1 int
	var n2 int
	n1 = 40
	n2 = 60
	res := func (num1 int, num2 int)(int) {
		return num1 + num2
	}(n1, n2)

	res1 := func (num1 int, num2 int)(int, int) {
		var t int
		t = num1
		num1 = num2
		num2 = t
		return num1, num2
	}(n1, n2)
	fmt.Println("调用匿名函数", res)
	fmt.Println("调用匿名函数res1", res1)
}