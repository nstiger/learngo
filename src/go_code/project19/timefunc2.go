package main
import (
	"fmt"
	"time"
	"strconv"
)
//需求，编写一段代码，统计一个函数执行的时间

func test03() {
	str := ""
	i := 0
	// for i := 0; i < 50000; i++ {
	// 	str += "hello" + strconv.Itoa(i)		
	// }    //跟以下for循环的写法效率一样
	for {
		str += "hello" + strconv.Itoa(i)
		if i == 50000 {
			break
		}
		i++
	}
}

func main() {
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Println("test03运行的时间是", end - start)
}