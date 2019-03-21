package main
import (
	"encoding/json"
	"fmt"
)

//输出字段名的首字母必须是大写的，如果想用小写的首字母，直接把结构体的字段名改成首字母小写是不行的
//JSON输出的时候必须注意，只有导出的字段才会被输出，如果修改字段名，那么什么都不会被输出
//所以必须通过struct tag定义来实现
type Server struct {
	ServerName string `json:"serverName"`
	ServerIP string	`json:"serverIP"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Local_Web", ServerIP: "172.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Local_DB", ServerIP: "172.0.0.2"})

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err: ", err)
	}

	fmt.Println(string(b))

}