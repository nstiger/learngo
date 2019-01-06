package process
import (
	"fmt"
	"go_code/chatroom/common/message"
	"encoding/json"
	"go_code/chatroom/client/utils"
)

type SmsProcess struct {

}

//发送群聊的消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//1.创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType
	//2.创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content  //内容
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3.将SmsMes序列化
	data, err :=json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}
	mes.Data = string(data)

	//4.将mes再次序列化
	data, err =json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json.Marshal fail =", err.Error())
		return
	}

	//将mes发送给服务器
	tf := &utils.Transfer {
		Conn : CurUser.Conn,
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}
	return
}