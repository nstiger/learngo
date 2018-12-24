package main
import (
	"fmt"
	"encoding/json"
)

//将结构体，map，切片进行序列化
type Monster struct {
	Name string `json:"monster_name"`  //反射机制
	Age int `json:"monster_age"`
	Birthday string
	Sal float64
	Skill string
}

func testStruct() {
	monster := Monster {
		Name : "牛魔王",
		Age : 500,
		Birthday : "1900-11-11",
		Sal : 8000.0,
		Skill : "牛魔拳",
	}

	//将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误，err=", err)
	}
	//输出序列化结果
	fmt.Printf("monster序列化的data = %s\n", data)
}

//将map进行序列化
func testMap() {
	var a map[string]interface{}
	//使用map应当先make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "红云洞"

	//将这个map序列化
	//map是引用类型
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误，err=%v", err)
	}
	//输出序列化结果
	fmt.Printf("Map类型序列化结果%v\n", string(data))
	
}

//将slice进行序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "10"
	m2["address"] = [2]string{"墨西哥","夏威夷"}
	slice = append(slice, m2)

	//将这个切片序列化
	//slice是引用类型
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误，err=%v", err)
	}
	//输出序列化结果
	fmt.Printf("Slice类型序列化结果%v\n", string(data))

}

//基本数据类型的序列化， 没有什么实际意义
func testFloat64() {
	var num1 float64 = 2345.67

	//将这个Float64序列化
	//是引用类型
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误，err=%v", err)
	}
	//输出序列化结果
	fmt.Printf("Float64类型序列化结果%v\n", string(data))

}

func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
