package main
import (
	"fmt"
)
func main() {

	//先定义一个切片，数据类型为map
	//再为切片make空间
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "曹操"
		monsters[0]["age"] = "56"
		monsters[0]["power"] = "魏"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "孙权"
		monsters[1]["age"] = "29"
		monsters[1]["power"] = "吴"
	}

	fmt.Println(monsters)
	
	//以下增加切片的方法将导致越界，切片只定义了两个map数据
	// if monsters[2] == nil {
	// 	monsters[2] = make(map[string]string)
	// 	monsters[2]["name"] = "刘备"
	// 	monsters[2]["age"] = "47"

	newMonster := map[string]string {
		"name" : "刘备",
		"age" : "47",
		"power" : "蜀",
	}
	//用append方法把新的map加入切片
	monsters = append(monsters, newMonster)
	
	
	newMonster01 := map[string]string {
		"name" : "李世民",
		"age" : "22",
		"power" : "唐",
	}
	monsters = append(monsters, newMonster01)

	//注意打印结果，切片是有序的但map是无序的

	fmt.Println(monsters)

}