package main
import (
	"fmt"
)
type Monster struct{
	Name string
	Age int
}
func main() {
	//创建结构体变量的第一种方式
	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1
	fmt.Println(monster1)
	fmt.Println(monster2)

	monster2.Name = "青牛精"
	fmt.Println(monster1)
	fmt.Println(monster2)

	//创建结构体变量的第二种方式
	monster3 := Monster {"蜘蛛精",300}
	fmt.Println(monster3)

}