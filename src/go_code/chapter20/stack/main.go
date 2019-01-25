package main
import (
	"fmt"
	"errors"
)

type stack struct{
	MaxTop int
	Top int
	arr[5] int
}

func (this *stack) Push(val int) (err error) {
	if this.Top == this.MaxTop - 1 {
		fmt.Println("stack full")
		return errors.New("栈满")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

func (this *stack) List() {
	if this.Top == -1 {
		fmt.Println("stack empty")
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

func (this *stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("栈空")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil

}

func main() {
	s := &stack{
		MaxTop : 5,
		Top : -1,
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	
	s.List()

	s.Pop()
	s.List()
}