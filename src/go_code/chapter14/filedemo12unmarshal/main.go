package main
import (
	"fmt"
	"encoding/json"
)

//将结构体，map，切片进行反序列化
type Monster struct {
	Name string 
	Age int 
	Birthday string
	Sal float64
	Skill string
}

func UnmarshalStruct() {
	//获取json字符串
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"1900-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("monster %v\n", monster)

}

//反序列化map
func UnmarshalMap() {
	//获取json字符串
	str := "{\"address\":\"红云洞\",\"age\":30,\"name\":\"红孩儿\"}"
	//定义map变量
	var a map[string]interface{}
	//反序列化map不需要make，make已经封装到Unmarshal函数里面了	
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err= %v\n", err)
	}
	fmt.Printf("map a = %v\n", a)
}

//反序列化slice
func UnmarshalSlice() {
	//获取json字符串
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," + 
	"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"10\",\"name\":\"tom\"}]"

	//定义slice变量
	var slice []map[string]interface{}
	//反序列化slice同样不需要make
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err= %v\n", err)
	}
	fmt.Printf("map slice = %v\n", slice)	

}


func main() {

	UnmarshalStruct()
	UnmarshalMap()
	UnmarshalSlice()
}