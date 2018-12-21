package main
import (
	"fmt"
)
func main() {
	//三种定义和赋值方式
	//第一种定义方式
	var monsters map[string]string
	monsters = make(map[string]string, 3)  //3不是必须的
	monsters["no1"] = "牛魔王"
	monsters["no2"] = "红孩儿"
	monsters["no3"] = "玉山公主"
	monsters["no4"]	= "狐狸精"
	monsters["no5"]	= "无极国王"
	monsters["no5"] = "五级国王"
	fmt.Println(monsters)

	//map的删除操作
	//1、用delete内建函数,如果key存在，就删除，否则不操作
	//delete(map, "key")
	//2、遍历所有的key，删除所有，或者用make方法，让原来的成为垃圾被回收
	delete(monsters, "no4")
	fmt.Println("删除了no4", monsters)


	//第二种定义方式
	heroes := map[string]string {
		"z01" : "宋江",
		"z02" : "武松",
		"z03" : "卢俊义",    //这个逗号不可少
	}
	fmt.Println(heroes)

	//第三种方式定义
	var students map[string]map[string]string
	students = make(map[string]map[string]string)

	students["s01"] = make(map[string]string)
	students["s01"]["name"] = "张无忌"
	students["s01"]["sex"] = "男"
	students["s01"]["address"] = "北京"

	students["s02"] = make(map[string]string)
	students["s02"]["name"] = "赵敏"
	students["s02"]["sex"] = "女"
	students["s02"]["address"] = "上海"

	students["s03"] = make(map[string]string)
	students["s03"]["name"] = "周芷若"
	students["s03"]["sex"] = "女"
	students["s03"]["address"] = "成都"
	
	fmt.Println(students)
	fmt.Println(students["s03"])
	fmt.Println(students["s03"]["address"])

	//map增加和修改，如果有key，则修改; 无key，则增加
	//map查找，不能用for遍历，因为key不是连续整数，只能用for-range
	//如果monsters这个map中存在no1, val是对应的值，findRes（可任意取名）返回true，否则返回false
	val, findRes := monsters["no1"]  //val是对应的value,findRes是bool值，找到了为true, 否则为false 
	if findRes {
		fmt.Printf("找到了monster no1, val =%v\n", val)
	} else {
		fmt.Printf("没有找到no1\n")
	}

	vals, findRess := students["s02"]
	if findRess {
		fmt.Printf("找到了s02,是%v\n", vals)
	} else {
		fmt.Println("没找到s02")
	}

	//map的遍历，k是key，v是value
	for k, v := range heroes {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	fmt.Println("********************* for - range ****************************")
	//value是map类型的遍历，两重for-range
	for k1, v1 := range students {
		fmt.Printf("k1=%v\n", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\tk2=%v, \tv2=%v\n", k2, v2)
		}
	}

	//map的长度 func len(v Type) int
	fmt.Println("students map的长度是：", len(students))
}