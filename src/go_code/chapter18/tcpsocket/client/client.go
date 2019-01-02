/*
客户端功能
1、编写客户端程序，能链接到服务器端的8888端口
2、客户端可以发送单行数据，然后退出
3、能通过终端输入数据（输入一行发送一行），并发送到服务器端
4、在终端输入exit，表示退出程序
*/

package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)


func main() {
	conn, err := net.Dial("tcp", "192.168.1.105:8888")
	if err != nil {
		fmt.Println("client dial err= ", err)
		return
	}
	fmt.Println("connect succes! conn=", conn)

	//功能一，客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin)   //os.Stdin代表标准输入

	//服务器端输入exit表示退出程序，for循环从这里开始
	for {
		//从终端拿到一行输入，准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ReadString err=", err)
		} 
		

		//判断line是否等于exit，等于退出
		//由于line是带"\n"的，先去掉它再做判断
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出程序")
			break
		}
		//再将line发送给服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
	}

}