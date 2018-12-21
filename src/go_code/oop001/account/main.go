package main
import (
	"fmt"
)
type Account struct{
	AccountNo string
	Pwd string
	Balance float64
}

func (a *Account) Deposit(pwd string, money float64) {
	if  pwd != a.Pwd {
		fmt.Println("帐号密码有误，请重新输入...\n")
		return
	} 
	if money <= 0 {
		fmt.Println("您输入的金额有误......")
		return
	} else {
		a.Balance += money
		fmt.Println("您已存入：", money, "您的当前余额是：", a.Balance)
		fmt.Println()
	}
}

func (a *Account) Withdraw(pwd string, money float64) {
	if pwd != a.Pwd {
		fmt.Println("帐号密码有误，请重新输入...")
		return
	}
	if money <= 0 || money > a.Balance {
		fmt.Println("您输入的金额有误......")
		return
	}
	a.Balance -= money
	fmt.Println("您本次取款：", money, "您的当前余额是：", a.Balance)
	fmt.Println()
}

func (a *Account) Query(pwd string) {
	if pwd != a.Pwd {
		fmt.Println("您输入的密码有误，请查对后重新输入...")
		fmt.Println()
	} else {
		fmt.Println("*********您的账户余额是：", a.Balance)
		fmt.Println()
	}	
}
func main() {
	var a = Account{
		AccountNo : "GS1001",
		Pwd : "888888",
		Balance : 100,
	}

	for {
		fmt.Println("请输入您需要的服务(1-4)：")
		fmt.Println("----------------------")
		fmt.Println("*      1、存 款      *")
		fmt.Println("*      2、取 款      *")
		fmt.Println("*      3、查 询      *")
		fmt.Println("*      4、退 出      *")
		fmt.Println("----------------------")
		var i int = 0
		var password string
		var inputmoney float64
		fmt.Scanln(&i)

		switch i {
		case 1: {
			fmt.Println("请输入您的密码：")
			fmt.Scanln(&password)
			fmt.Println("请输入您要存入的金额：")
			fmt.Scanln(&inputmoney)
			a.Deposit(password, inputmoney)
		}
		case 2: {
			fmt.Println("请输入您的密码：")
			fmt.Scanln(&password)
			fmt.Println("请输入您要取出的金额：")
			fmt.Scanln(&inputmoney)
			a.Withdraw(password, inputmoney)
		}
		case 3: {
			fmt.Println("输入您的查询密码：")
			fmt.Scanln(&password)
			a.Query(password)
		}
		case 4: return
			
		default: {
				fmt.Println("您的输入有误，请输入正确的数字......")

			}
		}
	}



	a.Query("888888")

	a.Deposit("888888", 3500)
	a.Withdraw("888888", 170)
	a.Query("888888")
	a.Query("66666")


}