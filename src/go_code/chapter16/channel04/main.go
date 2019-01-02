/*
应用实例3
需求：要求统计1-8000的数字中，哪些是素数
思路：
1、三个管道,intChan、primeChan、exitChan分别存放读入的数字、经计算为素数的数、标识4个协程完成处理的管道
2、1个读入1-80000数字的协程，负责将数字加入数字管道;
   4个协程负责从数字管道读取数字进行验证，如果为素数，将该数加入素数队列，如果本协程处理完毕，将true写入退出管道
3、主程序协程负责从exitChan读取4个T标识，如果读取完成则退出

*/



package main
import (
	"fmt"
)

func putNum(intChan chan int) {
	for i := 1; i <= 5000; i++ {
		intChan <- i
	}
	//向管道写入数据完成后关闭管道，关闭管道不影响读取
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	count := 0
	for {
		flag := true


		num, ok := <- intChan
		count++
		if !ok {
			break //
		}
		
		//判断取出的数字是否素数
		for i := 2; i < num; i++ {
			if (num % i == 0) {
				//primeChan <- num
				flag = false
				break
			} 
		} 
				
		//走出以上for循环的数是素数，写入primeChan
		if flag {
			primeChan <- num
		}
	}
	//退出for循环表示已经取不到数据了
	exitChan <- true
	fmt.Println("primeNum协程处理了", count)
}


func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	//1.开启putNum协程，向intChan管道写入8000个数据
	go putNum(intChan)

	//2.开启4个（或多个）primeChan协程，从intChan管道读取数据，并判断是否素数
	for i:=0; i<4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//3.主线程从exitChan取数据（4个），形成阻塞，如果4个数取完，说明4个primeNum协程都已经执行完毕
	//4.把这个阻塞也作为一个线程，并用匿名函数形式调用
	go func(){
		for i := 0; i < 4; i++ {
			<-exitChan  //取出的数据是什么不重要，可以直接丢弃
		}
		//关闭primeChan管道，打开阻塞
		close(primeChan)
	}()

	//5.遍历primeChan， 打印所有的素数
	for {
		v, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%v\n", v)
	}

	fmt.Println("主线程退出！")


}