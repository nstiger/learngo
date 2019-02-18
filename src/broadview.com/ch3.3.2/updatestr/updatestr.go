/*用[]byte和[]rune切片修改字符串*/
package main
import (
	"fmt"
)

func main() {

	//[]byte
	s := "Hello 世界！"
	b := []byte(s) //转换为[]byte，自动复制数据
	b[5] = ',' //修改[]byte,只能修改为英文的“，”，中文的“，”占三个字节。
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", b)

	fmt.Println("============rune方式=============")
	//[]rune
	r := []rune(s)	//转换为[]rune,自动复制数据
	r[6] = '中' 	//修改[]rune,只能用单引号
	r[7] = '国'
	fmt.Println(s)	
	fmt.Println(string(r)) //[]rune转换为字符串


}