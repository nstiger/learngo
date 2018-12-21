package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	//定义一个五个随机元素的数组并反转打印

	var intarr [5]int
	rand.Seed(time.Now().UnixNano())
	for i :=0; i< len(intarr); i++ {
		intarr[i] = rand.Intn(100)
	}
	fmt.Println(intarr)

	for i:=len(intarr); i >0; i-- {
		fmt.Printf(" %v", intarr[i-1])
		}
	fmt.Println()
}