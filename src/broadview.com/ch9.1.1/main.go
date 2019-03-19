package main
import "fmt"

//Shaper is a interface{}
type Shaper interface{
	Area() float32
}

//Square is a Shaper
type Square struct {
	side float32
}

//Area achieve Shaper interface
func (sq * Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	fmt.Printf("面积为：%f\n", sq1.Area())
	areaIntf := sq1
	fmt.Printf("面积为：%f\n", areaIntf.Area())
}