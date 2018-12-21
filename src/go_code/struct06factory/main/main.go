package main
import (
	"fmt"
	"go_code/struct06factory/model"
)

func main() {
	// var student model.NewStudent
	// student.Name = "tom"
	// student.Score = 88.8

	//model包定义的结构体类型的首字母是小写，所以main包无法直接调用
	//为了解决该问题，引入工厂模式
	//func NewStudent(n string, s float64) *student
	//在model包中，定义一个方法，通过*student返回值，绑定student结构体类型，方法的传入参数绑定结构体字段
	stu := model.NewStudent("tom", 77.80)
	fmt.Println(stu) //stu是一个student结构体类型的指针，因此，输出时表现为：&{..., ...}

	fmt.Println(*stu) //用取值符*，就可以将stu的值输出，结果表现为：{..., ...}

	//fmt.Println("name = ", stu.Name, "score = ", stu.Score) //(*stu).Name是标准写法，编译器优化。

	//如果现在model包中定义的NewStudent结构体类型的字段score的首字母是小写
	//这意味着在model包之外无法访问
	//为了访问score字段，在model包中定义了一个绑定student结构体（实际上是结构体指针）的方法，通过该方法，
	//我们可以在main包中，访问score字段的值

	fmt.Println("name=", stu.Name, "score=", stu.GetScore())



}