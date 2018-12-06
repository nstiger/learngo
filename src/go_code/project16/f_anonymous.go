package main
import (
	"fmt"
)
//匿名函数

//全局匿名函数，可以在整个程序中使用,写法区别注意！！！
	var (
		Fun1 = func (no1, no2 int) int {
		return no1 * no2
	}
	)

func main() {

	var n1 int
	var n2 int
	n1 = 40
	n2 = 60
	// res := func (num1 int, num2 int)(int) {
	// 	return num1 + num2
	// }(n1, n2)

	//以下方式还是属于把匿名函数赋给一个变量，只是有两个返回值需要接收。
	res1 := func (num1 int, num2 int)(int, int) {
		var t int
		t = num1
		num1 = num2
		num2 = t
		return num1, num2
	}

	//如果把匿名函数赋给一个变量，如a，可以多次调用。注意！！！赋值的时候函数定义后面不带传入参数。
	a := func (num1, num2 int) int {
		return num1 - num2
	}

	fmt.Println("调用匿名函数的第二种方式", a(n1, n2))
	fmt.Printf("a的数据类型%T\n", a)

	// fmt.Println("调用匿名函数", res)
	t3, t4 := res1(n1, n2)
	fmt.Println("调用匿名函数res1,t3,t4", t3, t4)

	var t1 int =100
	var t2 int =200
	fmt.Println("试试看a能不能用多次", a(t1, t2)) 


	res5 := Fun1(3,5)
	fmt.Println("全局匿名函数的用法～～", res5)
	
}