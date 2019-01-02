/*
反射最佳实践1：使用反射来遍历结构体的字段，调用结构体的方法并获取结构体标签的值
*/

package main
import (
	"fmt"
	"reflect"
)

//定义一个结构体
type Monster struct {
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex string
} 


//方法：返回两个数的和
func (m Monster) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

//方法：接收4个值，给s赋值
func (m Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

//方法：显示s的值
func (m Monster) Print() {
	fmt.Println("----------start------------")
	fmt.Println(m)
	fmt.Println("-----------end-------------")
}

//定义函数，接收空接口参数，调用反射机制获取Type、Kind、值等
func TestStruct(a interface{}) {
	//获取a的Type
	rTyp := reflect.TypeOf(a)
	//获取a的Value
	rVal := reflect.ValueOf(a)
	//获取a的Kind
	rKd := rVal.Kind()

	//如果传入的不是结构体，就退出
	if rKd != reflect.Struct {
		fmt.Println("传入参数不是结构体类型，退出！")
		return
	}

	//获取该结构体有几个字段
	num := rVal.NumField()
	fmt.Printf("rVal的字段共：%v个\n", num)

	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %v, 值为 %v\n", i, rVal.Field(i))
		//获取到struct的标签，需要用reflect.Type来获取tag标签的值
		//func (tag StructTag) Get(key string) string
		//Get方法返回标签字符串中键key对应的值。如果标签中没有该键，会返回""。如果标签不符合标准格式，Get的返回值是不确定的。
		tagVal := rTyp.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field为 %d， tag是 ：%v\n", i, tagVal)
		}
	}

	//获取该结构体有多少个方法
	//func (v Value) NumMethod() int
	//返回v持有值的方法集的方法数目。
	numOfMethod := rVal.NumMethod()
	fmt.Printf("结构体拥有%d个方法\n", numOfMethod)

	//方法的排序默认按照函数名的排序(ASCII码)
	rVal.Method(1).Call(nil)  //获取到第二个方法，调用它
	//调用结构体的第一个方法
	var params []reflect.Value  //声明切片
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))

	res := rVal.Method(0).Call(params) //传入参数是[]reflect.Value, 返回[]reflect.Value
	fmt.Println("res=", res[0].Int())  //返回结果，返回的结果是[]reflect.Value

}




func main() {
	//创建结构体实例
	var a Monster= Monster {
		Name : "黄鼠狼精",
		Age : 400,
		Score : 30.8,

	}

	//将monster实例传给TestStruct
	TestStruct(a)
}