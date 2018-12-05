package main
import (
	"fmt"
	"go_code/project15/utils"
)

 var a = test()

func test() int {
	fmt.Println("test() ~~~~~~~~")
	return 80
}
//init函数
func init() {
	fmt.Println("init().........", a - 2)
	// a:=2
	// b:=3
	//return 5
}
func main() {
	fmt.Println("main()..........", a)
	

	fmt.Println(utils.Age, utils.Name)

}