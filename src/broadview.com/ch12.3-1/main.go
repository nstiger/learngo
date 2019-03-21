/*
对数据集合排序不必考虑应当选择哪种排序方法，只需实现sort.Interface定义的三个方法：
Len():获取数据集合的长度
Less():比较两个数据元素的大小
Swap():交换两个元素的位置
*/
package main

import (
	"fmt"
	"sort"
)

type StuScore struct {
	name string
	score int
}

type StuScores []StuScore

func (s StuScores) Len() int {
	return len(s)
}

func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"王语嫣", 95}, 
		{"慕容复", 91}, 
		{"乔峰", 96}, 
		{"段誉", 90},
	}
	fmt.Println("-------------------")
	for _, v := range stus {
		fmt.Println(v.name, ": ", v.score)
	}
	fmt.Println()

	//StuScores类型已经实现了sort.Interface接口
	sort.Sort(stus)
	fmt.Println("-----------排序之后-----------")
	for _, v := range stus {
		fmt.Println(v.name, ": ", v.score)
	}

	//判断是否已经排好顺序
	fmt.Println("是否已经排序？", sort.IsSorted(stus))

}