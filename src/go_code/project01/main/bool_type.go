// --bool_type.go
//golang中bool类型的使用

package main
import "fmt"
import "unsafe"

func main()  {
	var b = false
	fmt.Println("b=", b)
	//bool类型占用1个字节存储空间
	fmt.Println("b 的占用空间 =", unsafe.Sizeof(b))
	//bool类型智能取true或false
}