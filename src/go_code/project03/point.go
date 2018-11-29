package main
import (
	"fmt"
)

func main()  {
	var i int =10
	var ptr *int = &i
	
	fmt.Println("i的初值=", i)
	*ptr = 25
	fmt.Println("i的地址:,i的值", &i,i)
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr指向的值：%v\n", *ptr)

	fmt.Printf("i%v,ptr%v,ptrAddress%v,ptrValue%v",i, ptr, &ptr, *ptr)
}

