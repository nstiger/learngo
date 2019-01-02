/*
goroutine中使用recover捕获panic，解决协程中出现错误，导致程序崩溃的问题
通过启用协程recover捕获这个错误，进行处理，即使这个协程出现问题，主程序不受影响，仍然可以继续运行
*/

package main 
import (
	"fmt"
	"time"
)

func SayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello, world!", i)
	}
}

func Test() {
	//调用一个匿名函数处理当Test函数出现错误的情况，用recover捕获错误。
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("Test()出现错误！")
		} 
	}()

	//定义一个map，不make直接使用（赋值），这样会导致一个错误
	var myMap map[int]string
	myMap[0] = "golang" //error

}

func main() {
	go SayHello()
	go Test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok = ", i)
		time.Sleep(time.Second * 2)
	}


}