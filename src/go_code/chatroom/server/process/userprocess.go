package process2
import (
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	"encoding/json"
	"go_code/chatroom/server/utils"
	"go_code/chatroom/server/model"
	
)

//定义结构体，绑定以下函数
type UserProcess struct {
	//字段？
	Conn net.Conn
	//增加一个字段，表示该链接属于谁
	UserId int
}

//这里我们编写通知所有在线的用户的方法
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历onlineUsers, 然后一个一个的发送 NOtifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤掉自己
		if id == userId {
			continue
		}
		//开始通知(单独写一个方法)
		up.NotifyMeOnline(userId)
	}

}

func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装我们的NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)

	//将mes再次序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}

	//准备发送，创建Transfer实例
	tf := &utils.Transfer{
		Conn : this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}

}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data，并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	//我们需要到redis数据库去完成验证
	//1.使用model.MyUserDao到redis去验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."

		}
	} else {
		registerResMes.Code = 200
	}

	//3.将registerResMes序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}

	//4.将data(现在是一个切片)赋值给resMes
	resMes.Data = string(data)

	//5.将resMes序列化 ，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}

	//6.发送data，将其封装到一个WritePkg函数中
	//因为使用分层模式（MVC），我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

//编写一个函数，专门处理登录类型的请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心代码...
	//1.先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1.先声明一个resMes 
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.再声明一个LoginResMes,并完成赋值
	var loginResMes message.LoginResMes

	//我们需要到redis数据库去完成验证
	//1.使用model.MyUserDao到redis去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		} 
		} else {
		loginResMes.Code = 200
		//这里，因为用户登录成功，我们就把该登录成功的用户放入到userMgr中
		//同时，将登录成功的用户userId 赋给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他的在线用户，我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前在线用户的id放入到loginResMes.UsersId
		//遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user, "登录成功")
	}

	// //判断是否合法用户，先用固定的语法，不走数据库
	// //如果用户id=100, 密码=123456,认为合法，否则不合法
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	//合法
	// 	loginResMes.Code = 200
	// } else {
	// 	//不合法
	// 	loginResMes.Code = 500	//状态码，表示该用户不存在
	// 	loginResMes.Error = "该用户不存在，请注册再使用"
	// }

	//3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}

	//4.将data(现在是一个切片)赋值给resMes
	resMes.Data = string(data)

	//5.将resMes序列化 ，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}

	//6.发送data，将其封装到一个WritePkg函数中
	//因为使用分层模式（MVC），我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return

}
