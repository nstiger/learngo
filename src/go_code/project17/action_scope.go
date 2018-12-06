package main
import (
	"fmt"
)
//变量的作用域
var name string = "tom"

/*下面语法是错误的，该语句等价于 
   var Name string
   Name = string 这是一条赋值语句，不允许在函数体之外使用！
   
Name := "mary" */

	func test01() {
		fmt.Println("test01 print", name)
	}

	func test02() {
		//下一句相当于给全局变量重新赋值
		name ="jack"
		//下一句相当于重新定义一个局部变量name，该变量的引用不影响全局变量的值
		//name := "jack"
		//var name string = "jack"
		fmt.Println("test02 print", name)
	}

func main() {
	fmt.Println("main first print", name)
	test01()
	test02()
	test01()
}