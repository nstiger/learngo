/*
channel使用细节和注意事项
使用select可以解决从管道取数据的阻塞问题
因为有时候不确定何时关闭管道，此时可以使用select解决问题
*/

package main
import (
	"fmt"
)

func main() {
	intChan :=make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i< 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)		
	}

	//传统的方法遍历管道时，如果未关闭会阻塞导致deadlock
	label:
	for {
		select {
			case v := <-intChan :
				fmt.Printf("从intChan里面取出数据%d\n", v)
			case v := <-stringChan :
				fmt.Printf("从stringChan里面取出数据%s\n", v)
			default :
				fmt.Printf("取不到数据，结束！程序员可以自定其他逻辑！\n")
				//return
				break label
		}
	}




}