package main
import (
	"fmt"
	"flag"
)

//用flag解析命令行参数
func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "用户名，默认为localhost")
	flag.IntVar(&port, "p", 3303, "用户名，默认为3303")

	//这里有一个非常重要的操作，转换，必须调用该方法
	flag.Parse()

	fmt.Printf("user=%v pwd=%v host=%v port=%v\n", user, pwd, host, port)
}