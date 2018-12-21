package main
import (
	"fmt"
	"go_code/familyaccountoop/utils"
)

func main() {
	fmt.Println("面向对象的家庭收支记录软件")

	utils.NewFamilyAccount().MainMenu()

}