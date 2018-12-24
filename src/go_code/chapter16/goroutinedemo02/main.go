package main
import (
	"fmt"
	"time"
	"sync"
)

/*
思路：
1、编写一个函数，计算n个数的阶乘，并放入到map中
2、启动多个协程，将统计结果放入map
3、定义一个全局的map
*/

var ( 
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	lock sync.Mutex
)	
func test(n int) {
	

	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//在这里加锁，写入myMap
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()  
	//问题：concurrent map writes,200个线程同时对map做写入操作
}

func main() {

	//启动200个协程
	for i:=1;i<=20;i++ {
		go test(i)		
	}

	//加入主线程等待时间,休眠10秒钟,目的是等待协程执行完毕
	time.Sleep(time.Second * 10)

	//此处也要加锁，
	lock.Lock()
	//输出结果，遍历这个结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}