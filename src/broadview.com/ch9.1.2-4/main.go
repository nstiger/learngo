package main

import (
	"fmt"
	"math"
)

//Square is a Shaper
type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	
	//areaIntf的类型是否是Square
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("areaIntf的类型是：%T\n", t)
	}
	if u, ok :=areaIntf.(*Circle); ok {
		fmt.Printf("areaIntf的类型是：%T\n", u)
	} else {
		fmt.Println("areaIntf不含类型为Circle的变量")
	}

	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Square类型的 %T 值为：%v\n", t, t)
	case *Circle:
		fmt.Printf("Circle类型的 %T 值为：%v\n", t, t)
	case nil:
		fmt.Printf("nil值：发生了意外。\n")
	default:
		fmt.Printf("未知类型 %T\n", t)
	}

}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}


