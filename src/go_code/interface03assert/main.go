package main
import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct{
	Name string

}
//让Phone实现Usb接口的两个方法
func (p Phone) Start(){
	fmt.Println("手机开始工作.......")
}

func (p Phone) Stop(){
	fmt.Println("手机停止工作.......")
}

//Phone特殊的方法Call
func (p Phone) Call(){
	fmt.Println("手机 在打电话.......")
}

type Camera struct{
	Name string
}
//让Camera实现Usb接口的两个方法
func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

//相机的特殊方法Take
func (c Camera) Take() {
	fmt.Println("相机拍照咔咔卡......")
}

type Computer struct{
	

}
//编写一个方法Working，接收一个Usb接口类型的变量
func (c Computer) Working(usb Usb) {
	usb.Start()

	//这里要增加一个类型断言，当usb变量是Phone类型时，调用Phone类型的Call方法
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}

	if camera, ok := usb.(Camera); ok {
		camera.Take()
	}

	usb.Stop()
}

func main() {
	var computer Computer
	// var phone Phone
	// var camera Camera

	//	computer.Working(phone)
	//  computer.Working(camera)

	var usbArr [3]Usb 
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}

	//遍历usbArr数组，如果是Phone类型变量，则除了Usb接口声明的方法之外，还要调用Phone特有的方法
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
	
}