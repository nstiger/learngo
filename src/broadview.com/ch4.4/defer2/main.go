/*
defer是一个栈，遵循先入后出，后进先出，当一个函数内部有多个defer语句时，最后面的defer语句最先执行
*/


package main

import (
	"fmt"

)

var i = 0

func print(i int) {
	fmt.Println(i)
}

func main() {
	for ; i < 5; i++ {
		defer print(i)
	}
}

/*
输出结果：
4
3
2
1
0
*/