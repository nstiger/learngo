/*
range 右边表达式可以是string	array/slice		map		channel
返回值的输出的第一个值是index		index		key		element            
        第二个值 str[index]		str[index]		m[key]	
*/


package main

import (
	"fmt"
)

func main() {
	str := "abcz"
	for i, v := range str {
		fmt.Printf("str的第%d个元素的值是%d\n", i, v)
	}

	m := map[string]int{"a": 1, "b" :2, "c" :5}
	for k, v := range m { //返回key, value
		fmt.Println(k, v)
	}

	numbers := []int{1, 2, 3, 4}
	for i, x := range numbers {
		fmt.Printf("第%d次， x的值是%d \n", i, x)
	}
}