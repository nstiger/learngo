package main
import (
	"fmt"
)

func main() {
	var name string
	var age int8
	var salary float32
	var pass bool
/*******  第一种方式，使用fmt.Scanln函数，用回车等待接收函数
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水：")
	fmt.Scanln(&salary)

	fmt.Println("请输入是否通过考试：")
	fmt.Scanln(&pass)

	fmt.Printf("姓名：%v\t年龄：%v\t薪水：%v\t是否通过考试：%v\n",name, age, salary, pass)
********/
	//第二种方式，用fmt.Scanf函数一次性接收输入内容
	fmt.Println("请输入您的姓名、年龄、薪水、是否通过考试，中间用空格隔开")
	fmt.Scanf("%s %d %f %b", &name, &age, &salary, &pass)
	fmt.Printf("姓名：%v\n年龄：%v\n薪水：%v\n是否通过考试：%v\n",name, age, salary, pass)
}