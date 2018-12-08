package main
import (
	"fmt"
)

func main() {
	var heroes [3]string = [3]string {"micky", "jerry", "tom"}
	for index, value:=range heroes {
		fmt.Printf("index=%v, value=%v\n", index, value)
	}
	//[micky jerry tom]
	fmt.Println(heroes)
	heroes[2] = "duck"
	fmt.Println(heroes)

	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	var numarr [5]int = [5]int {2,4,5,6,8}
	for i := 0; i < len(numarr); i++ {
		fmt.Printf("下标=%v, num=%v, 数组元素的地址=%v\n", i, numarr[i], &numarr[i])
	}

	fmt.Println("**********************************************************")
	var score [5]float32 = [...]float32 {2:78.5, 0:66, 1:88, 4:56, 3:96}
	for i, v :=range score {
		fmt.Printf("index:=%v\t, value=%v\t, 数组元素的地址=%p\n", i, v, &score[i])
	}

	fmt.Println("-----------------------------------------------------------")
	//数组的缺省值
	var arr01 [3]int
	var arr02 [3]float64
	var arr03 [3]string
	var arr04 [3]bool
	fmt.Println(arr01,arr02,arr03,arr04)

}