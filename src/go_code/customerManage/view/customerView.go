package main
import (
	"fmt"
	"go_code/customerManage/service"
	"go_code/customerManage/model"
)
type customerView struct{
	//用于接收客户输入的菜单选项
	key string
	//用于for循环的退出
	loop bool

	//为了调用service包中的customerService结构体，添加customerService字段
	customerService *service.CustomerService 
}

//显示所有客户信息
func (this *customerView) list() {
	//首先，获取当前客户的所有信息，在切片中。调用customerService结构体的List方法返回
	customers := this.customerService.List()
	//开始显示
	fmt.Println("\n---------------------------客户信息------------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	//在这里遍历客户的信息
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("--------------------------客户信息完成----------------------------")
	fmt.Println()
}

func (this *customerView) add() {
	fmt.Println("--------------------添加客户----------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	
	//构建一个新的customer实例
	customer := model.NewCustomer2(name, gender, age, phone, email)

	//调用Add
	if this.customerService.Add(customer) {
		fmt.Println("--------------------添加完成----------------------")
	} else {
		fmt.Println("--------------------添加失败----------------------")
	}
} 

//得到用户输入id，修改用户信息
func (this *customerView) update() {
	fmt.Println("---------------修改客户信息------------------")
	fmt.Println("请选择待修改客户编号(-1退出):")

	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return  //放弃删除操作
	}

	//1.首先，获取当前客户的所有信息，在切片中。调用customerService结构体的List方法返回
	customers := this.customerService.List()

	//2.调用FindById方法，找到id对应记录的index（切片索引）
	index := this.customerService.FindById(id)
	

	//需要一个方法，得到id客户的初值

	fmt.Printf("姓名[%v]：", customers[index].Name)
	name := customers[index].Name
	fmt.Scanln(&name)
	fmt.Printf("性别[%v]：", customers[index].Gender)
	gender := customers[index].Gender
	fmt.Scanln(&gender)
	fmt.Printf("年龄[%v]：", customers[index].Gender)
	age := customers[index].Age
	fmt.Scanln(&age)
	fmt.Printf("电话[%v]：", customers[index].Phone)
	phone := customers[index].Phone
	fmt.Scanln(&phone)
	fmt.Printf("邮箱[%v]：", customers[index].Email)
	email := customers[index].Email
	fmt.Scanln(&email)

	//调用customerService Update方法
	if this.customerService.Update(id, name, gender, age, phone, email) {
		fmt.Println("-----------------修改成功------------------")
	}

}

//得到用户输入id，删除该id所在记录
func (this *customerView) delete() {
	fmt.Println("----------------------删除客户----------------------")
	fmt.Println("请输入待删除的客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return  //放弃删除操作
	}

	fmt.Println("确认是否删除(Y/N):")
	//这里可以加入一个for循环，判断是否输入的y或n，否则要求重新输入

	choice := ""
	fmt.Scanln(&choice)

	if choice == "y" || choice == "Y" {
		//调用customerService的Delete方法
		if this.customerService.Delete(id) {
			fmt.Println("-----------------------删除成功------------------------")
		} else {
			fmt.Println("-----------------------删除失败，id号不存在--------------")
		}
	}
}

//封装退出功能
func (this *customerView) exit() {
	fmt.Println("您确定要退出吗？(Y/N):")

	for {
		fmt.Scanln(&this.key)
			if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
				break
			} else {
				fmt.Println("输入有误，您确定要退出吗？(Y/N):")
			}
	}

	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}


//用于显示主菜单，绑定customerView结构体(指针）,在主函数中用实例来调用本函数就可以显示主菜单
func (this *customerView) mainMenu(){
	for {
		fmt.Println("\n---------------------客户关系管理系统----------------------")
		fmt.Println("                     1.添 加 客 户")
		fmt.Println("                     2.修 改 客 户")
		fmt.Println("                     3.删 除 客 户")
		fmt.Println("                     4.客 户 列 表")
		fmt.Println("                     5.退      出")
		fmt.Print("请选择(1~5)：")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
		//	fmt.Println("tianjia")
			this.add()
		case "2":
		//	fmt.Println("xiugai")
			this.update()
		case "3":
		//	fmt.Println("shanchu")
			this.delete()
		case "4":
		//	fmt.Println("kehuliebiao")
			this.list()
		case "5":
			//this.loop = false
			this.exit() 
		default:
			fmt.Println("您的输入有误，请重新输入!")
		}
		//如果loop为false则退出for循环
		if !this.loop {
			break
		}
	}

}


func main() {
	//创建customerView的实例，为什么实例名和结构体名可以一样？？？？
	customerView :=customerView{
		key : "",
		loop : true,
	}
	//在这里完成对customerView结构体中customerService字段(是一个指针类型)的初始化。
	//在service包中有NewCustomerService方法可以完成它的初始化
	//该方法正好返回一个customerService指针
	customerView.customerService = service.NewCustomerService()
	//调用mainMenu方法显示主菜单
	customerView.mainMenu()
//	fmt.Println("ok")
}
