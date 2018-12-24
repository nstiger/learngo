package main
import (
	"fmt"
)

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)

	fmt.Println()
	//intChan <-300

	//遍历管道
	intChan2 := make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	//1.遍历管道时不能使用普通的for循环，len在发生变化
	//2.如果channel没有关闭，则会出现deadlock错误

	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v = ", v)
	}

}