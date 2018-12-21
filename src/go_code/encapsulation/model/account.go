package model
import "fmt"

type account struct{
	accountNo string
	balance float64
	pwd string
}

func NewAccount(a string, b float64, p string) *account{
	return &account{
		accountNo : a,
		balance : b,
		pwd : p,
	}
}

func (account *account) SetAccountNo(accNo string) {
	if len(accNo) < 6 && len(accNo) >10 {
		fmt.Println("帐号输入有误，请重新输入6-10位的帐号！")
		return
	} else {
		account.accountNo = accNo
	}
}

func (account *account) GetAccountNo() string {
	return account.accountNo
}

func (account *account) SetBalance(bal float64) {
	if bal <= 20 {
		fmt.Println("您的余额不足......")
	} else {
		account.balance = bal
	}
}

func (account *account) GetBalance() float64 {
	return account.balance
}

func (account *account) SetPwd(p string) {
	if len(p) != 6 {
		fmt.Println("密码设置有误")
		return
	} else {
		account.pwd = p
	}
}

func (account *account) GetPwd() string {
	return account.pwd
}