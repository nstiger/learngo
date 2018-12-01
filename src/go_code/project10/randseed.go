package main
import (
	"fmt"
	"math/rand"
	"time"
	
)

//随机生成1-100之间的数，计算得到”99“所用的次数
//n = rand.Intn(100) + 1 //[0,100)
//为了得到随机数，还需要为rand函数生成一个种子
//time.Now().Unix():返回一个unix起始时间至当前的秒数
//rand.Seed(time.Now().Unix())

func main() {

	var count int = 0  //计数器
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println("n=", n)
		count++
		if (n == 99) {
			break //表示跳出for循环
		}
	}
	fmt.Println("生成99,一共使用了", count)

}