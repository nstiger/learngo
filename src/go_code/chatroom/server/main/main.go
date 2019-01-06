package main
import (
	"fmt"
	"net"
	"time"
	"go_code/chatroom/server/model"
)

func init() {
	//当服务器启动时，我们就去初始化我们的redis连接池
	initPool("localhost:6379", 16, 0, 300 * time.Second)
	initUserDao()
}

// //封装一个函数，把读取和返回消息放在一个函数里
// func readPkg(conn net.Conn) (mes message.Message, err error) {

// 	buf := make([]byte, 8096)

// 	fmt.Println("读取客户端发送的数据...")
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		//err = errors.New("read pkg header error")
// 		return
// 	}

// 	//把buf[:4](切片)转成uint32
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[0:4])

// 	//从conn里面读取pkgLen个字节，存入buf
// 	n, err := conn.Read(buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		//err = errors.New("read pkg body error")
// 		return
// 	}

// 	//把buf反序列化成 -> message.Message
// 	//mes前面要加上取地址符&!!技术就是一扇窗户纸！
// 	err = json.Unmarshal(buf[:pkgLen], &mes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal fail, err= ", err)
// 		return
// 	}
// 	return
// }

// func writePkg(conn net.Conn, data []byte) (err error) {

// 	//1.先发送长度给对方
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
// 	//发送长度
// 	n, err := conn.Write(buf[:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write(bytes) fail", err)
// 		return
// 	}
// 	//再发送消息本身
// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Write(bytes) fail", err)
// 		return
// 	}

// 	return
// }

// //编写一个函数，专门处理登录类型的请求
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	//核心代码...
// 	//1.先从mes中取出mes.Data，并直接反序列化成LoginMes
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal fail err=", err)
// 		return
// 	}

// 	//1.先声明一个resMes 
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType

// 	//2.再声明一个LoginResMes,并完成赋值
// 	var loginResMes message.LoginResMes


// 	//判断是否合法用户，先用固定的语法，不走数据库
// 	//如果用户id=100, 密码=123456,认为合法，否则不合法
// 	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
// 		//合法
// 		loginResMes.Code = 200
// 	} else {
// 		//不合法
// 		loginResMes.Code = 500	//状态码，表示该用户不存在
// 		loginResMes.Error = "该用户不存在，请注册再使用"
// 	}

// 	//3.将loginResMes序列化
// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail, err=", err)
// 		return
// 	}

// 	//4.将data(现在是一个切片)赋值给resMes
// 	resMes.Data = string(data)

// 	//5.将resMes序列化 ，准备发送
// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail, err=", err)
// 		return
// 	}

// 	//6.发送data，将其封装到一个WritePkg函数中
// 	err = writePkg(conn, data)
// 	return

// }

// //编写一个ServerProcessMes函数
// //功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
// func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 		case message.LoginMesType:
// 			//处理登录类型的消息
// 			err = serverProcessLogin(conn, mes)
// 		case message.RegisterMesType:
// 			//处理注册类型的消息
// 		default :
// 			fmt.Println("消息类型不存在，无法处理...")
		
// 	}
// 	return
// }

func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()
	//这里调用总控，创建一个实例
	processor := &Processor {
		Conn : conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误，err=", err)
		return
	}

	// //循环读取客户端发送的信息
	// for {
	// 	//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message,Err
	// 	//调用readPkg()
	// 	mes, err := readPkg(conn)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			fmt.Println("客户端退出，服务器端也正常退出...")
	// 			return
	// 		} else {
	// 			fmt.Println("readPkg err=", err)
	// 			return
	// 		}
	// 	}
	// 	//fmt.Println("mes =", mes)
	// 	err = serverProcessMes(conn, &mes)
	// 	if err != nil {
	// 		return
	// 	}
	// }
}

//这里我们编写一个函数，完成对UserDao的初始化任务
func initUserDao() {
	//这里pool本身就是一个全局的变量
	//这里要注意一个初始化顺序问题
	//initPool, 再initUserDao
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	//当服务器启动时，我们就去初始化我们的redis连接池
	initPool("localhost:6379", 16, 0, 300 * time.Second)
	initUserDao()
	//提示信息
	fmt.Println("服务器[新的结构]正在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	//defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	//一旦监听成功，就等待客户端来链接服务器
	for {
		fmt.Println("正在等待客户端的链接....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
	//一旦链接成功，就启动一个协程保持与客户端的通讯
		go process(conn)
	}
	
}