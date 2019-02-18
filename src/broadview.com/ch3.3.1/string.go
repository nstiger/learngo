package main
import (
	"fmt"
)

func main() {
	s := "aA 你 2"
	fmt.Println("字符串长度：", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	//编译错误，字符串的值不能修改
	//s := "你好，"
	t := "你好，"
	//字符串可以连接，但原字符串不会改变
	t += "世界。"
	fmt.Println(s)
	//s[0] = 'L'
	fmt.Println(t)

}