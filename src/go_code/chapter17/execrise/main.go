/*
给定变量var v float64 = 1.2, 请使用反射来得到它的reflect.Value，然后获得对应的Type，kind和值，并将
reflect.Value转换为interface{}， 再将interface{}转换成float64
*/
package main
import (
	"reflect"
	"fmt"
)

func main() {
	var v float64 = 1.2
	rVal := reflect.ValueOf(v)
	rTyp := reflect.TypeOf(v)

	fmt.Printf("rVal.type=%v, rVal.Kind =%v, rVal.value=%v ", rTyp, rTyp.Kind(), rVal)

	iV := rVal.Interface()
	num := iV.(float64)
	fmt.Println("num = ", num)

	fmt.Println("*******************************")
	//传地址，用Elem()得到变量反射类型的指针，用SetFloat方法修改变量的值
	rValf := reflect.ValueOf(&v)

	rValf.Elem().SetFloat(11.22)
	fmt.Println("v = ", v)

	fmt.Println("***************改string变量的值****************")
	//修改string变量的值
	var str string = "tom"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jerry")
	fmt.Println("str = ", str)
	
}