package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go 向redis写入和读取数据
	//1.连接到redis服务
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	defer conn.Close()   //!!!关闭连接

	//2.通过go 向redis写入数据string[key-val]
	_, err1 := conn.Do("Set", "name", "tomjerry猫猫")
	if err1 != nil {
		fmt.Println("set err:", err1)
		return
	}

	//3.通过go 向redisd获取数据string[key-val]
	//使用redis包提供的方法直接将读取的数据转成string
	r, err1 := redis.String(conn.Do("Get", "name"))
	if err1 != nil {
		fmt.Println("set err:", err1)
		return
	}


	fmt.Println("caozuo success......", r)
}