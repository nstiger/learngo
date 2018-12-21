package main
import (
	"fmt"
)
// 有序数组的二分查找，用递归调用
//本例用升序数组

func binarySearch(arr [9]int, leftIndex int, rightIndex int, middle int, findval int){
	middle = (leftIndex + rightIndex) / 2
	//找到递归调用的退出机制
	if leftIndex > rightIndex {
			fmt.Println("找不到......")
			return
		}

	if arr[middle] > findval {
		binarySearch(arr, leftIndex,middle - 1, middle, findval)
	} else if arr[middle] < findval {
		binarySearch(arr, middle + 1, rightIndex, middle, findval) 
	} else {
		fmt.Println("Goooooot it!", middle)
	}
}


func main() {
	var arr [9]int = [...]int {1, 7, 11, 23, 35, 40, 66,71, 87}
	fmt.Println(arr)
	binarySearch(arr, 0, 8, 4, 35)
}