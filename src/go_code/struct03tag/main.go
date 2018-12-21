package main
import(
	"fmt"
	"encoding/json"
)
type Monster struct{
	Name string `json:"name"`  //反引号 struct tag
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	monster :=Monster{"牛魔王", 500, "芭蕉扇"}
	fmt.Println(monster)

	jsonStr, err := json.Marshal(monster) 
	//json.Marshal函数，引入encoding/json包
	
	if err != nil {
		fmt.Println("json转换错误")
	}
	fmt.Println("json字串", string(jsonStr))

}