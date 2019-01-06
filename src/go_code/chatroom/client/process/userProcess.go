package process
import (
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	"go_code/chatroom/client/utils"
	"encoding/binary"
	"encoding/json"
	"os"
)

type UserProcess struct {
	//暂时不需要字段
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	//2.准备通过conn给服务器发送消息
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.声明一个LoginMes 结构体变量
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.把data赋给了mes.Data字段
	mes.Data = string(data)
	//6.将mes（结构体）序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn : conn,
	}
	//发送data给服务器
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
	}
	mes, err = tf.ReadPkg()  //mes就是RegisterResMes
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化成RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

//写一个函数，完成登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	//下一步就要开始订协议，
	// fmt.Printf(" userId = %d  userPwd = %s \n", userId, userPwd)
	// return nil 
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	//2.准备通过conn给服务器发送消息
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.声明一个LoginMes 结构体变量
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5.把data赋给了mes.Data字段
	mes.Data = string(data)

	//6.将mes（结构体）序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7.到这个时候，data就是我们要发送的消息
	//7.1首先要把data的长度发送给服务器
	//获取到data的长度，转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	fmt.Printf("客户端， 发送消息的长度=%d, 发送的内容是%s", len(data), string(data))

	//发送消息本身，就是data，已经序列化了
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//休眠20秒
	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20秒...")

	//这里还需要处理服务器返回的消息
	//创建一个Transfer的实例
	tf := &utils.Transfer {
		Conn : conn,
	}
	mes, err = tf.ReadPkg()  
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		//fmt.Println("登录成功")
		//可以显示当前在线用户列表，遍历loginResMes.UsersId
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersId {
			//如果我们要求不显示自己在线，下面我们增加一个代码
			if v == userId {
				continue
			}
			fmt.Println("用户id：\t", v)
			//完成客户端的onlineUser初始化
			user := &message.User{
				UserId : v,
				UserStatus : message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		//这里我们还需要启动一个协程
		//保持和服务器端的通讯，如果服务器有推送的消息给客户端
		//则接收并显示在客户端的终端
		go serverProcessMes(conn)

		//1.显示我们登录成功的菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return

}