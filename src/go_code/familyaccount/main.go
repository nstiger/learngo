package main
import (
	"fmt"
)

func main() {
	//定义一个变量用于存放余额
	var balance float64 = 10000
	//定义一个变量用于接收金额
	var money float64
	//定义一个变量用于接收说明
	var note string

	//定义一个变量用于拼接title
	details := "收支\t账户金额\t收支金额\t说  明"

	//定义一个变量用来接收键盘输入值
	 var key string = ""

	//定义一个变量用来退出for循环
	var loop bool = true

	//定义一个变量查看是否发生了收支记录
	flag := false
	
	for {
	
		//显示主菜单
		fmt.Println("\n-------------------家庭收支记账软件-------------------")
		fmt.Println("                   1.收支明细")
		fmt.Println("                   2.收入登记")
		fmt.Println("                   3.支出登记")
		fmt.Println("                   4.退出")

		fmt.Print("欢迎使用本软件，请选择您需要的操作（1-4）:")
	
		 fmt.Scanln(&key)
		 
	 	switch key {
			case "1":
				fmt.Println("---------当前收支明细记录-----------")
				if flag {
					fmt.Println(details)
				} else {
					fmt.Println("目前还没有收支记录，来一笔吧.....")
				}
				
		
			case "2":
				fmt.Println("本次收入金额：")
				fmt.Scanln(&money)
				//修改余额
				balance += money
				fmt.Println("本次收入说明：")
				fmt.Scanln(&note)
				//拼接details字符串
				details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
				flag = true  //发生了收支记录
			case "3":
				fmt.Println("本次支出金额：")
				fmt.Scanln(&money)
				//修改余额
				if money > balance {
					fmt.Println("您的余额不足......")
					break
				}
				balance -= money
				fmt.Println("本次支出说明：")
				fmt.Scanln(&note)
				//拼接details字符串
				details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
				flag = true //发生了收支记录
				
			case "4":
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
					loop = false
				}
			default:
				fmt.Println("您的输入有误，请重新输入......")
		}

		if !loop {
			break //退出for循环
		} 
	}
	fmt.Println("您已退出本软件！")
}