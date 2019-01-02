package main
import (
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	_"net"
	"encoding/binary"
	"encoding/json"
	_"errors"
	_"io"
)


//封装一个函数，把读取和返回消息放在一个函数里
func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)

	fmt.Println("读取客户端发送的数据...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	//把buf[:4](切片)转成uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//从conn里面读取pkgLen个字节，存入buf
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}

	//把buf反序列化成 -> message.Message
	//mes前面要加上取地址符&!!技术就是一扇窗户纸！
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err= ", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {

	//1.先发送长度给对方
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
	//再发送消息本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	return
}