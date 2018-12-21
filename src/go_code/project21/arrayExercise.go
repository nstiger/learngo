package main
import (
	"fmt"
)
func main() {
	//创建一个byte类型的26个元素的数组，分别放置‘A'～’Z'，使用for循环访问并打印出来。
	var myChars [26]byte
	for i:=0; i< 26; i++ {
		myChars[i] = 'A' + byte(i) 
	} 
	for i:=0; i< 26; i++ {
	fmt.Printf("%c\n",myChars[i])
	}
}