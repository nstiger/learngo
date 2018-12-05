package main
import (
	"fmt"
)

func main() {
	//break - lable 案例
	lable2:
	//设置lable2标签，从这里跳出下一个for循环
	//打印结果为0,1
	for i := 0; i < 4; i++ {
		//lable1:   
		//设置lable1标签,从这里跳出下一个for循环
		//打印结果是0，1，0，1，0，1，0，1
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break lable1
				break lable2
			}
			fmt.Println("j=", j)
		}
	}
}