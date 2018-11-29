package main
import (
	"fmt"
	"strconv"
)

func main() {
	// var a int
	// var b float32
	// var c float64
	// var d bool
	// var e string
	
	// fmt.Printf("a=%d, b=%f, c=%f, d=%v, e=%v\n", a, b, c, d,e)

	// c = 36955411.123589774
	// fmt.Printf("c=%9.3f, c的原始值=%v",c, c )
	var b,b2,b3 bool
	str1 := "true"
	str2 := "t"
	str3 := "1"

	b, _= strconv.ParseBool(str1)
	fmt.Printf("b=%v,	b的数据类型：%T\n",b,b)

	b2, _= strconv.ParseBool(str2)
	fmt.Printf("b2=%v,	b2的数据类型：%T\n",b2,b2)
	
	b3, _= strconv.ParseBool(str3)
	fmt.Printf("b3=%v,	b3的数据类型：%T\n",b3,b3)
}