package main
import (
	"fmt"
)
/*  1、定义一个Circle结构体，字段为radius
	2、声明一个方法area和Circle绑定，可以返回面积
*/
type Circle struct{
	Radius float64
}

func (c Circle) area() float64{
	return 3.14 * c.Radius * c.Radius 
}

func (cp *Circle) pArea() float64{
	fmt.Printf("pArea cp的地址： %p\n", cp)
	return 3.14 * cp.Radius * cp.Radius   //标准写法 (*cp).Radius
	//return 3.14 * (*cp).Radius * (*cp).Radius

}

func main() {
	var c1 Circle
	c1.Radius = 4.0
	area1 := c1.area()
	fmt.Println("c1的面积=", area1)

	//一般情况下，会在定义方法的时候，把结构体变量定义为指针，指向调用该方法的变量（实例）
	//这样做的目的是因为结构体是值类型，为了提高调用效率，调用指针可以直接改变结构体变量（实例）的字段值！
	//由于Golang底层已经做了优化，所以当在结构体方法定义时，结构体变量为指针时，标准的(*c).Radius可以简写为c.Radius
	//同样，调用该方法的变量写法(&c1).area()也可以简写为c1.area()!!

	var cp Circle
	cp.Radius = 9.0
	area2 := cp.pArea()  //标准写法(&p.pArea())
	//area2 := (&cp).pArea()
	fmt.Println("cp的面积=", area2)

	fmt.Printf("main cp的地址： %p\n", &cp)
}