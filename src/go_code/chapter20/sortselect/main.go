package main
import (
	"fmt"
)

func SelectSort(arr *[6]int) {

	//数组元素个数为j, 遍历j-1次，完成全部的遍历
	for j := 0; j < len(arr)-1; j++ {
		//第j次排序，用数组首元素跟其后元素比较，遍历所有元素，选出最小值
		//记录最小值和其下标
		Max := arr[j]
		MaxIndex := j

		for i := j + 1; i < len(arr); i++ {
			if Max < arr[i] {
				Max = arr[i]
				MaxIndex = i
			}
		}
		//如果MaxIndex下标发生改变，则说明需要交换
		if MaxIndex != j {
			arr[MaxIndex] = arr[j]
			arr[j] = Max
		}
		fmt.Printf("第%d次遍历：%v\n", j+1, arr)
	}

		// 	//第一次排序，用数组首元素跟其后元素比较，遍历所有元素，选出最小值
		// //记录最小值和其下标
		// Max := arr[0]
		// MaxIndex := 0

		// for i := 0 + 1; i < len(arr); i++ {
		// 	if Max < arr[i] {
		// 		Max = arr[i]
		// 		MaxIndex = i
		// 	}
		// }
		// //如果MaxIndex下标发生改变，则说明需要交换
		// if MaxIndex != 0 {
		// 	arr[MaxIndex] = arr[0]
		// 	arr[0] = Max
		// }
		// fmt.Printf("第1次遍历：%v\n", arr)

}

func main() {
	//var arr [6]int
	arr := [6]int{10, 8, 14, 23, 95, 101}

	fmt.Println("main() arr=", arr)

	SelectSort(&arr)
}