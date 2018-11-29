package main
import (
	"fmt"
)

func main() {
	var days int =97
	var weeks int
	weeks = days / 7
	days = days % 7

	fmt.Printf("离放假还有%d周，另%d天", weeks, days)

	var a int = 9
	var b int = 3

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("a= ,b=", a , b)
}