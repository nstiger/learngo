package main
import "fmt"

//冒泡排序算法bubble sort，把数组元素按从小到大的值排列

func fab(arr *[9]int) {

	for j:=0; j< len(arr)-1; j++ {
		for i:=0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp 
			}
		}
	}

}

func main() {
	var arr [9]int = [9]int {24, 69, 80, 57, 13, 1,133, 55, 77}
	fmt.Println(arr)

	fab(&arr)
	fmt.Println(arr)
}