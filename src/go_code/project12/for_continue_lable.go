package main
import (
	"fmt"
)
//break与continue的比较
//break跳出当前for循环，从上一级for循环的循环变量迭代处执行
//continue跳出当前循环，从本级for循环的循环变量迭代处执行
//lable指明跳出哪一层for循环

func main() {
	lable2:
	//此时continue跳转到下一个for循环执行
	//应该打印出4轮0，1
	for i:= 0; i < 4; i++ {
		//lable1:
		//此时continue跳转到下一个for循环处执行
		//此时应该打印4轮0,1,3,4,5,6,7,8,9
		for j := 0; j < 10; j++ {
			if j == 2 {
				//continue lable1
				continue lable2
			}
			fmt.Println("j=",j)
		}
	}
}