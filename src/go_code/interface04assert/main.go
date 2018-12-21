package main
import (
	"fmt"
)

//类型断言的最佳实践2

//写一个函数，可以判断输入的参数是什么类型
type Student struct {
	Name string
}

func TypeJudge(items ...interface{}){
	for index, x := range items {
		switch x.(type) {
		case float32:
			fmt.Printf("第%v个参数的类型是 float32, 值是=%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数的类型是 float64, 值是=%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数的类型是 整型, 值是=%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数的类型是 string, 值是=%v\n", index, x)
		case Student:
			fmt.Printf("第%v个参数的类型是 Student结构体类型, 值是=%v\n", index, x)
		case *Student:
			fmt.Printf("第%v个参数的类型是 Student结构体指针类型, 值是=%v\n", index, x)

		default:
			fmt.Printf("第%v个参数的类型未知, 值是=%v\n", index, x)
		}
	}
}



func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string ="tom"
	address := "北京"
	n4 := 300
//	var n5 *int
	stu := Student{
		Name : "tom",
	}

	stu001 := &Student{
		Name : "mary",
	}

	TypeJudge(n1, n2, n3, name, address, n4, stu, stu001)
}