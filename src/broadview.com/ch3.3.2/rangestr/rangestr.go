/*遍历字符串中的字符的方式*/
package main
import (
	"fmt"
)

func main() {
	s := "我是中国人"
	for i := 0; i < len(s); i++ {
	fmt.Printf("%c", s[i])
	//fmt.Printf("%d = %v \n", i, s[i]) //输出单字节值
	}
	fmt.Printf("\n")
	for _, v := range s {
		fmt.Printf("%c", v)
	}
	fmt.Print("\n")
}