/*
tcp socket编程快速入门
服务器端功能
1、编写一个服务器端程序，在8888端口监听
2、可以和多个客户端创建链接
3、链接成功后，客户端可以发送数据，服务器端接收数据，并显示在终端上
4、先使用telnet来测试，然后客户端程序来测试
*/

package main
import (
	"fmt"
	"net"
	_"io"

)

//写一个协程，处理某个客户端发送的信息
//它需要传入一个net.Conn类型的参数，即链接
func process(conn net.Conn) {
	
	//循环接收客户端发送的数据
	defer conn.Close()  //延时关闭链接，否则链接无限增长
	
	for {
		//创建一个新的切片，准备作为参数传给conn.Read
		buf := make([]byte, 1024)
		//1.等待客户端发送信息
		//2.如果客户端没有write[发送]，那么协程一直阻塞
		//fmt.Printf("服务器在等待客户端%s 发送信息......\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)  //从conn读取
		if err != nil {
			fmt.Println("服务器端的read err=", err)
			return    //一旦对方已关闭连接，就退出for循环，退出协程
		}
		//显示客户端发送的信息到服务器的终端
		//注意只显示到buf[:n],如果显示全部切片内容（1024字节），将出现乱码
		fmt.Print(string(buf[:n])) 
	}
}


func main() {
	fmt.Println("服务器开始监听......")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err =", err)
		return
	}
	defer listen.Close()  //延时关闭链接

	//循环等待客户端来链接

	for {
		//等待客户端链接
		fmt.Println("等待客户端链接......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err =", err)
		} else {
			fmt.Printf("Accept succes conn=%v, 客户端IP：%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备起一个协程，为客户端服务
		go process(conn)
	}


}