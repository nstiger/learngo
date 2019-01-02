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

	//2.通过go 向redis写入数据Hash[key:[string]string]
	_, err1 := conn.Do("HMSet", "user02", "name", "johnson", "age", 22)
	if err1 != nil {
		fmt.Println("Hset err:", err1)
		return
	}

	//3.通过go 向redisd获取数据
	//使用redis包提供的方法直接将读取的数据转成string或int
	r, err1 := redis.Strings(conn.Do("HMGet","user02", "name", "age"))
	if err1 != nil {
		fmt.Println("hget err:", err1)
		return
	}

	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i , v)
	}  
}