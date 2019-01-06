package main
import (
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/process"
	"go_code/chatroom/server/utils"
	"io"
)

type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	//看看能否收到客户端发送的群发消息
	fmt.Println("mes=", mes)

	switch mes.Type {
		case message.LoginMesType:
			//处理登录类型的消息
			//创建一个UserProcess的实例
			up := &process2.UserProcess {
				Conn : this.Conn, 
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType:
			//处理注册类型的消息
			up := &process2.UserProcess {
				Conn : this.Conn, 
			}
			err = up.ServerProcessRegister(mes)
		case message.SmsMesType :
			//处理群发消息类型的消息,创建一个SmsProcess实例完成转发群聊消息
			smsProcess := &process2.SmsProcess{}
			smsProcess.SendGroupMes(mes)


		default :
			fmt.Println("消息类型不存在，无法处理...")
		
	}
	return
}

func (this *Processor) process2() (err error) {

	//循环读取客户端发送的信息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message,Err
		//调用readPkg()
		//创建一个Transfer实例，完成读包任务
		tf := &utils.Transfer {
			Conn : this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也正常退出...")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err 
			}
		}
		//fmt.Println("mes =", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
	
}