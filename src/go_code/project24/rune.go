package main
import (
	"fmt"
)

func main() {
	str := "hello@atguigu"
	arr :=str[6:]
	fmt.Println("arr=", arr)

	//把str中的第一个字母h改为z

	var arr1 []byte = []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)

	//把str中的第一个字母z改为中文”北“
	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println("str=",str) 
}	