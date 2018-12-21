package main
import (
	"fmt"
)
/*课堂练习
1、使用map[string]map[string]string的map类型
2、key：表示用户名，是唯一的，不可以重复
3、如果某个用户名存在，就将其密码改为：”888888“，如果不存在就增加这个用户信息（包括昵称和密码）
4、编写一个函数modifyUser,完成上述功能
*/

func modifyUser(users map[string]map[string]string, name string){
	if users[name] != nil {
		//找到这个用户，更改他的密码
		users[name]["pwd"] = "888888"
	} else {
		//没找到这个用户，创建这个用户，并设置密码和nickname
		users[name] = make(map[string]string, 2)
		users[name]["psd"] = "888888"
		users[name]["nickname"] = "昵称～" + name
	}
}
func main() {
	Users := make(map[string]map[string]string, 10)
	//增加一个已有的用户，调用函数修改其密码
	Users["smith"] = make(map[string]string, 2)
	Users["smith"]["pwd"] = "999999"
	Users["smith"]["nickname"] = "小花猫"
	
	modifyUser(Users, "smith")
	modifyUser(Users, "tom")
	modifyUser(Users, "mary")

	fmt.Println(Users)	

}