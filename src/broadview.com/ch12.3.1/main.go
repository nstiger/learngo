/*
排序包sort实现了四种基本排序算法：插入排序、归并排序、堆排序和快速排序，但是这四种排序方法是不公开的，
它们只被sort包内部使用

*/
package main
import (
	"fmt"
	"sort"
)

func main() {
	//三个int切片的方法：sort.Ints(), sort.Reverse(), sort.SearchInts()
	p := []int{77, 66, 55, 99, 87, 0, 3, 65, 4}
	fmt.Println(p)
	//对[]int切片排序更常使用sort.Ints(),而不是直接使用IntSlice类型
	sort.Ints(p)
	fmt.Println(p)

	fmt.Println("--------------")
	s := []int{7, 66, 15, 99, 87, 0, 3, 65, 4}
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	//sort.Reverse(sort.IntSlice(s))
	fmt.Println(s)

	fmt.Println("----------string---------")
	//与Sort(),IsSorted(),Search()对应的三个字符串切片[]string的方法：Strings(),StringsAreSorted(), SearchStrings()
	str := []string{"study", "golang", "php3.5", 
				"python", "microsoft", "oracle", "php2.7", "Apple", "Zippo"}
	fmt.Println(str)
	fmt.Println(sort.StringsAreSorted(str))
	sort.Strings(str)
	fmt.Println(str)

	num := sort.SearchStrings(str, "php3")
	fmt.Println(num)

}

