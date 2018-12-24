package main
import (
	"fmt"
)

/*
管道channel基本介绍
管道是引用类型
*/

type Cat struct {
	Name string
	Age int
}

func main() {

	//1.定义
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("管道intChan的地址=%v 管道本身的地址%v=\n", intChan, &intChan)

	//2.向管道写入数据
	intChan<- 10
	num := 50
	intChan<- num

	//3.看看管道的大小和容量cap
	fmt.Printf("管道intChan的大小%v 管道的容量cap=%v\n", len(intChan), cap(intChan))

	//4.从管道中取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("管道intChan的大小%v 管道的容量cap=%v\n", len(intChan), cap(intChan))	

	//存取map类型的管道
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["city1"] = "北京"
	m1["city2"] = "上海"

	mapChan <- m1

	m2 := make(map[string]string, 10)
	m2["hero1"] = "宋江"
	m2["hero2"] = "武松"

	mapChan <- m2

	//var m3, m4 map[string]string
	m3 := make(map[string]string, 10)
	m3 = <-mapChan
	m4 := <-mapChan
	fmt.Printf("从mapchan取出m3=%v\n", m3)
	fmt.Printf("从mapchan取出m4=%v\n", m4)

	//存取结构体类型的管道
	var catChan chan Cat
	catChan = make(chan Cat, 10)
	cat1 := Cat{
		Name:"tom",
		Age : 18,
	}
	cat2 := Cat{Name:"jerry", Age:20,}
	catChan <- cat1
	catChan <- cat2

	cat3 := <-catChan
	fmt.Printf("从catChan管道中取出cat3 %v\n", cat3)

	//存取指针类型的管道
	var catChan2 chan *Cat
	catChan2 = make(chan *Cat, 10)
	catChan2 <- &cat1
	catChan2 <- &cat2
	//取出
	cat11 := <- catChan2
	cat22 := <- catChan2

	fmt.Println(cat11, cat22)
	fmt.Println(*cat11, *cat22)

	//存取空接口类型的管道
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)

	allChan <- cat1
	allChan <- cat2
	allChan <- 10
	allChan <- "jack"

	//取出
	cat211 := <- allChan
	cat212 := <- allChan
	v1 := <- allChan
	v2 := <- allChan

	fmt.Println(cat211, cat212, v1, v2)

	fmt.Printf("cat211的类型%T， cat211的值%v\n", cat211, cat211)

	/*以下用法编译错误: 
	fmt.Printf("cat211.Name", cat211.Name)
	*/
	//使用类型断言！
	a := cat211.(Cat)
	fmt.Println("a=", a.Name)
 	
}