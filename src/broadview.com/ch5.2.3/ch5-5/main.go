package main
import (
	"fmt"
)

//IsDivid is a export function
func IsDivid(a, b int) bool {

	if a % b != 0 {
		return false
	}
	return true
}

type boolFunc func(int, int) bool //声明一个boolFunc函数类型

func filter(slice []int, f boolFunc) ([]int, []int) {
	var result3, result5 []int
	for _, value := range slice {
		if f(value, 3) {
			result3 = append(result3, value)
		} 
		if f(value, 5) {
			result5 = append(result5, value)
		}
	}
	return result3, result5
}

func main() {
	slice := []int {15, 16, 20, 21, 36, 77, 88, 104, 120, 127}
	fmt.Println("slice = ", slice)
	divid3, divid5 := filter(slice, IsDivid) //函数当作值来传递 IsDivid函数符合filter的f参数的形态
	fmt.Printf("divid3 = %v, divid5 = %v \n", divid3, divid5)

}

