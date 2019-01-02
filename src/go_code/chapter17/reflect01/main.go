/*
演示对（基本数据类型，interface{}，reflect.value）进行反射的基本操作
*/
package main
import (
	"fmt"
	"reflect"
)

//定义一个结构体类型
type Student struct{
	Name string
	Age int
}

//专门演示基本数据类型反射
func reflectTest01(b interface{}) {
	//通过反射，获取传入的变量的type、kind和值
	//1.先获取reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp =", rTyp)

	//2.获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal =", rVal)

	//以下写法类型不匹配
	// n2 := 2 + rVal       
	//要通过reflect.Value类型的方法
	n2 := 2 + rVal.Int()
	fmt.Println("n2 =", n2)
	fmt.Printf("rVal的值 = %v rVal的类型：%T \n", rVal, rVal)

	//3.下面将rVal转成interface{}, 用reflect.Value类型的Interface方法
	iV := rVal.Interface()
	//4.用断言将iV转成int类型
	num2 := iV.(int) 
	fmt.Println("num2 = ", num2)
}

//演示结构体数据类型的反射
func reflectTest02(b interface{}) {
	//通过反射，获取传入的变量的type、kind和值
	//1.先获取reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp =", rTyp)

	//2.获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal =", rVal)



	//3.下面将rVal转成interface{}, 用reflect.Value类型的Interface方法
	iV := rVal.Interface()
	//4.使用简单带检测的类型断言将interface类型转为结构体类型，可以用switch断言的形式来实现(11章)
	stu, ok := iV.(Student)
	if ok {
		fmt.Println("stu.Name=", stu.Name)
	}

	//5.获取传入变量的kind,两种方法
	kind1 := rTyp.Kind()
	kind2 := rVal.Kind()
	fmt.Printf("kind1= %v kind2= %v\n", kind1, kind2)

}

func main() {
	var num int = 100
	reflectTest01(num)

	stu := Student{
		Name : "tom",
		Age : 20,
	}

	reflectTest02(stu)

	//常量定义的一点补充
	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
}