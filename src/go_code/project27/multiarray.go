package main
import (
	"fmt"
)

func main() {
	var multiArr [2][3]int32 = [...][3]int32 {{0,0,0},{0,1,0}}
	fmt.Println(multiArr)
	fmt.Printf("二维数组的下标是%p\n", &multiArr)
	fmt.Printf("multiArr[0]数组地址：=%p\n", &multiArr[0])
	fmt.Printf("multiArr[1]数组地址：=%p\n", &multiArr[1])

	fmt.Printf("multiArr[0][0]:=%p\n", &multiArr[0][0])
	fmt.Printf("multiArr[1][0]:=%p\n", &multiArr[1][0])

	fmt.Printf("multiArr[1][1]:=%p\n", &multiArr[1][1])
}