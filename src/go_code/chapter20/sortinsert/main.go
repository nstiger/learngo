/*
插入排序法
把n个待排序元素看作一个有序表和无序表，开始时有序表只包含一个元素，无序表包含n-1个元素，排序过程中
每次从无序表取出第一个元素，将它的排序码与有序表元素的排序码逐个比较，将它插入到有序表中的适当位置，使之成为新的有序表
*/

package main
import (
	"fmt"
	"math/rand"
	"time"
)

func InsertSort(arr *[80000]int) {

	
	// //第一次插入
	// insertVal := arr[1]	//保存无序表第一个元素的值
	// insertIndex := 1	//保存无序表第一个元素的下标

	// for insertIndex >= 0 && insertVal > arr[insertIndex - 1] {
	// 	arr[insertIndex + 1] = arr[insertIndex]
	// 	insertIndex--
	// }
	// fmt.Printf("第一次插入的结果=%v\n", arr)

	for i := 1; i < 80000; i++ {
	
	//第一次插入
	insertVal := arr[i]	//保存无序表第一个元素的值
	insertIndex := i	//保存无序表第一个元素的下标

	for insertIndex > 0 && insertVal > arr[insertIndex - 1] {
		arr[insertIndex ] = arr[insertIndex - 1]
		insertIndex--
	}
	//如果走出for循环，说明已经找到插入位置
	arr[insertIndex] = insertVal
//	fmt.Printf("第 %d 次插入的结果=%v\n", i, arr)

	}
}

func main() {
	//arr := [8]int{23, 0, 12, 56, 34, 87, -1, 23}
	//fmt.Println("原始数组为=", arr)
	var arr[80000] int
	for i:=0; i<80000; i++ {
		arr[i] = rand.Intn(9000000)
	}

	start := time.Now().Unix()
	InsertSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("插入排序，8万数据排序时间 %d 秒！！\n", end - start)

	for j:=0; j < 50; j++ {
		fmt.Println(arr[j])
	}
}