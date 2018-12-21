package main
import (
	"fmt"
)

func main() {
	var str string = "Hello, world.您吃了吗？"  //字符串中如果有中文会有显示问题。用[]rune切片方式解决
    st001 := []rune(str)
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	for i :=0; i < len(st001); i++ {
		fmt.Printf("%c \n", st001[i])
	}


		

	var str1 string = "hello, andrewong.早上好"    //该方式没有中文字符显示问题。
	for index, val := range str1 {
		fmt.Printf("index = %v, val = %c\n", index, val)
	}

	//打印99乘法表
	for i := 1; i <= 9; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print(j, " * ", i, " = ", i * j, " ")
		}
		fmt.Println()
	}



}

