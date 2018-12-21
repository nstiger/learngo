package main
import (
	"fmt"
	"go_code/encapsulation/model"
)

func main() {
	a := model.NewAccount("Gs0001001", 100.00, "888888")
	a.SetAccountNo("GS10000001")
	no := a.GetAccountNo()
	fmt.Println("帐号：", no)
	fmt.Println("余额：", a.GetBalance())
	fmt.Println("密码：", a.GetPwd())

	a.SetBalance(500)
	a.SetPwd("999999")

	fmt.Println("帐号：", no)
	fmt.Println("余额：", a.GetBalance())
	fmt.Println("密码：", a.GetPwd())

	a.SetBalance(19)
	a.SetPwd("8989")
}