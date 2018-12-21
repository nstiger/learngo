package main
import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)



//定义Student结构体
type Student struct{
	Name string
	Age int
	Score float64
}

//定义Student切片
type StudentSlice []Student


//实现sort.Sort的三个方法, Len\Less\Swap
//func Sort(data interface)

func (ss StudentSlice) Len() int {
	return len(ss)
} 

func (ss StudentSlice) Less(i, j int) bool {
	return ss[i].Score > ss[j].Score
} 

func (ss StudentSlice) Swap(i, j int) {
	// temp := ss[i]
	// ss[i] = ss[j]
	// ss[j] = ss[i]
	ss[i], ss[j] = ss[j], ss[i]
} 


func main() {

	//将Student结构体变量按照Score字段由高到低排序
	
	//测试是否可以对结构体切片进行排序
	var stus StudentSlice
	for i := 0; i < 20; i++ {
		stu := Student{
			Name : fmt.Sprintf("学生～00%d", i),
			Age : rand.Intn(30),
			Score : rand.Float64() * 100,
		}
		//把float64转换成小数点后保留2位，比较复杂！
		//这个方式的转换比较好，是处理了四舍五入的！
		//func ParseFloat(s string, bitSize int) (f float64, err error)
		stu.Score, _ = strconv.ParseFloat(fmt.Sprintf("%.2f",stu.Score), 64)
		//将stu append到stus切片变量中
		stus = append(stus, stu)
	}

	//排序前
	fmt.Println("排序前")
	for _ , v := range stus{
		fmt.Println(v)
	}

	//排序
	sort.Sort(stus)

	//排序后
	fmt.Println("------------排序后--------------")
	for _, v := range stus{
		fmt.Println(v)
	}

}