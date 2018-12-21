package main
import (
	"fmt"
)
//用结构类型作为map的value
type stu struct {
	Name string
	Age int
	Address string
}

func main()  {
	students := make(map[string]stu, 10)
	//创建两个学生
	stu1 := stu{"tom", 18, "北京"}
	stu2 := stu{"mary",20, "上海"}

	students["no1"] =stu1
	students["no2"] =stu2

	fmt.Println(students)

	//遍历各个学生的信息
	for k, v := range students{
		fmt.Printf("学生的编号是：%v\n", k)
		fmt.Printf("学生的姓名：%v\n", v.Name)
		fmt.Printf("学生的年龄：%v\n", v.Age)
		fmt.Printf("学生的地址：%v\n", v.Address)
	}
}