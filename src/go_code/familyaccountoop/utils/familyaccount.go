package utils
import (
	"fmt"
)

type FamilyAccount struct{
	//定义一个字段用于存放余额
	balance float64
	//定义一个字段用于接收金额
	money float64
	//定义一个字段用于接收说明
	note string

	//定义一个字段用于拼接title
	details string            //:= "收支\t账户金额\t收支金额\t说  明"

	//定义一个字段用来接收键盘输入值
	key string         //=""

	//定义一个字段用来退出for循环
	loop bool            //true

	//定义一个字段查看是否发生了收支记录
	flag bool         //false

}

//工厂模式，构造函数
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		balance : 10000.0,
		money : 0.0,
		note : "",
		details :  "收支\t账户金额\t收支金额\t说  明",
		key : "",
		loop : true,
		flag : false,
	}
} 

//定义一个方法绑定FamilyAccount结构体，用于显示主菜单
func (this *FamilyAccount) MainMenu() {

	//显示主菜单
	for {
	
		fmt.Println("\n-------------------家庭收支记账软件-------------------")
		fmt.Println("                   1.收支明细")
		fmt.Println("                   2.收入登记")
		fmt.Println("                   3.支出登记")
		fmt.Println("                   4.退出")

		fmt.Print("请选择（1-4）:")

		fmt.Scanln(&this.key)
		
			switch this.key {
				case "1":
					this.ShowDetails()
				case "2":
					this.Income()
				case "3":
					this.Pay()
				case "4":
					this.Exit()
				default:
					fmt.Println("您的输入有误，请重新输入......")
			}
	
			if !this.loop {
				break //退出for循环
			}
	} 	
}

//定义一个方法绑定FamilyAccount结构体，用于显示收入明细
func (this *FamilyAccount) ShowDetails() {
	fmt.Println("---------当前收支明细记录-----------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("目前还没有收支记录，来一笔吧.....")
	}


}

//定义一个方法绑定FamilyAccount结构体，用于记录收入
func (this *FamilyAccount) Income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	//修改余额
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	//拼接details字符串
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true  //发生了收支记录
}

//定义一个方法绑定FamilyAccount结构体，用于记录支出
func (this *FamilyAccount) Pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	//修改余额
	if this.money > this.balance {
		fmt.Println("您的余额不足......")
		//break
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	//拼接details字符串
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true //发生了收支记录
}

//定义一个方法绑定FamilyAccount结构体，用于退出程序
func (this *FamilyAccount) Exit() {
//完善退出功能
	fmt.Println("您确定要退出吗？y/n")

	choice := ""
	for {
		fmt.Scanln(&choice)
			if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("您的输入有误，请重新输入y/n")
	}

	if choice == "y" {
		this.loop = false
	}
}


