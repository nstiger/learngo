package model
import (
	"net"
	"go_code/chatroom/common/message"
)

//因为在客户端，我们很多地方会使用到CurUser，
type CurUser struct {
	Conn net.Conn
	message.User
}