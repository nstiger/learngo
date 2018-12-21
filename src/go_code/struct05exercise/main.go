package main
import (
	"fmt"
)

type integer int
func (i integer) jious() bool {
	if i % 2 == 0 {
		return true
	} else {
		return false
	}
} 

func main() {
	//写一个方法判断一个数是奇数还是偶数
	fmt.Println()

	var i integer = 12
	fmt.Println(i, i.jious())	
}