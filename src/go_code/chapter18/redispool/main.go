/*
连接池的使用案例
用init函数初始化连接池
*/
package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

//程序启动时，初始化链接池
func init() {
	pool = &redis.Pool{
		MaxIdle: 8,			//最大空闲链接数
		MaxActive : 0,		//表示和数据库的最大链接数，0表示没有限制
		IdleTimeout: 100,	//最大空闲时间
		//初始化链接的代码，链接哪个ip的redis
		Dial: func() (redis.Conn, error) {		
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	//先从pool取出一个链接
	conn := pool.Get()
	defer conn.Close()

	//操作redis
	_, err:= conn.Do("Set", "name", "汤姆猫")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	//取出数据
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	
	fmt.Println("r =", r)

}