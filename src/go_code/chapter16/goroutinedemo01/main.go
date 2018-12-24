package main
import (
	"fmt"
	"time"
	"strconv"
)
/*
编写一个程序完成如下功能
1、在主线程中开启一个goroutine,该协程每隔一秒输出"hello, world"
2、在主线程中，也每隔一秒输出"hello, golang"，输出10次后，退出程序
3、要求主线程和协程goroutine同时执行
4、画出主线程和goroutine执行流程图
*/

func test() {

	for i:=1; i<=10; i++ {
		fmt.Println("hello, world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test()

	for i:=1; i<=10; i++ {
		fmt.Println("hello, golang!" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}