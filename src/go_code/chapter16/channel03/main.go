
/*
goroutine和channel的结合
应用实例1
完成goroutine和channel协同工作的案例
1、开启一个writeData的协程，向管道intChan中写入50个整数
2、开启一个readData协程，从管道intChan中读取writeData写入的数据
3、注意：writeData和readData操作的是同一个管道
4、主线程需要等待writeData和readData协程都完成工作后才能退出（管道）
*/

package main
import (
	"fmt"
)

func writeData(intChan chan int) {
	for i:=1; i<=50; i++ {
		intChan<-i
		fmt.Printf("write data %v\n", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <- intChan
		if !ok {
			break
		}
		fmt.Printf("从intChan读取read data = %v\n", v)
	}

	//读取数据结束后，写入exitChan,并关闭exitChan
	exitChan <- true
	close(exitChan)

}

func main() {

	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <- exitChan
		if !ok {
			break
		}
	}
}