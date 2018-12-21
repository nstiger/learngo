package main
import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct{

}
//让Phone实现Usb接口的两个方法
func (p Phone) Start(){
	fmt.Println("手机开始工作.......")
}

func (p Phone) Stop(){
	fmt.Println("手机停止工作.......")
}

type Camera struct{

}
//让Camera实现Usb接口的两个方法
func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct{

}
//编写一个方法Working，接收一个Usb接口类型的变量
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	var computer Computer
	var phont Phone
	var camera Camera

	computer.Working(phont)
	computer.Working(camera)
}