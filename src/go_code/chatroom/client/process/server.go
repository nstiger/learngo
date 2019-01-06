package process
import (
	"fmt"
	"os"
	"net"
	"encoding/json"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"bufio"
)

func ShowMenu() {
	fmt.Println("---------------欢迎xxx成功登录-----------------")
	fmt.Println("---------------1.显示在线用户列表---------------")
	fmt.Println("---------------2.发送消息-----------------")
	fmt.Println("---------------3.信息列表-----------------")
	fmt.Println("---------------4.退出系统-----------------")
	fmt.Println("请选择（1-4）：")
	var key int
//	var content string
	var content *bufio.Reader

	//因为我们总会使用到SmsProcess实例，因此我们将其定义在swtich外部
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		// fmt.Println("请输入你想对大家说点什么：")
		// fmt.Scanf("%s\n", &content)
		// smsProcess.SendGroupMes(content)

		//尝试解决输入空格出错问题
		fmt.Println("想对大家说点啥？")
		content = bufio.NewReader(os.Stdin)
		str, _ := content.ReadString('\n')
		smsProcess.SendGroupMes(str)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入有误，请重新输入")
		
	}
}

//和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个Transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn : conn,
	}

	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err:= tf.ReadPkg()
		if err !=nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//如果读到消息，又是下一步处理逻辑
		switch mes.Type {
		case message.NotifyUserStatusMesType :	//有人上线了
			//1.取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2.把这个用户的信息，状态保存到客户map[int]User中
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType :	//有人群发消息了
			outputGroupMes(&mes)

		default :
			fmt.Println("服务器端返回了未知类型的消息")
		}
		//fmt.Printf("mes=%v\n", mes)

	}
}