package main
import (
	"fmt"
	"reflect"
)

func testInt(b interface{}) {
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetInt(110)
}

func main() {
	num := 9
	testInt(&num)
	fmt.Println(num)
}