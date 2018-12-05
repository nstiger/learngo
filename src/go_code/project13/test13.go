package main
import (
	"fmt"
)









func main() {
	var num1 float32 = 50.0
	var num2 float32 = 10.0
	var result float32 
	var operator byte = '*'

	switch operator {
	case '+': 
		result = num1 + num2
	case '-': 
		result = num1 - num2
	case '*': 
		result = num1 * num2
	case '/': 
		result = num1 / num2
	default :
		fmt.Println("fault.......")
	}

	//fmt.Println(num1, char(operator), num2, "=", result)
	fmt.Printf("%v %c %v = %v\n", num1, operator, num2, result)




}