package main
import (
	"fmt"
)
func main() {
	//slice的第一种定义方式
	var intArr [9]int = [...]int {1, 22, 33, 66, 99, 98, 7 , 8 ,9 }
	slice := intArr[1:7]

	fmt.Println(intArr)
	fmt.Println("slice的元素：", slice)
	fmt.Println("slice的长度：", len(slice))
	fmt.Println("slice的容量：", cap(slice))

	fmt.Printf("intArr下标1的元素的内存地址=%p\n", &intArr[1])
	fmt.Printf("slice的值=%v\n", slice)
	fmt.Printf("slcie第一个元素的值=%v\n", slice[0])
	fmt.Printf("slice的内存地址=%p\n", &slice)

	fmt.Println("***********************************************")
	//slice的第二种定义方式
	var mySlice []float64 = make([]float64, 4, 10)
	mySlice[1] = 12.60
	mySlice[3] = 10
	fmt.Println("mySlice=", mySlice)
	fmt.Println("***********************************************")
	//slice的第三种定义方式
	var strslice []string = []string {"tom", "jerry", "micky"}
	fmt.Println(strslice)
	fmt.Println("strslice size", len(strslice))
	fmt.Println("strslice cap", cap(strslice))
 }