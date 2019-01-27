package main
import (
	"fmt"
	"errors"
	"strconv"
)

type Stack struct{
	MaxTop int
	Top int
	arr[20] int
}

func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop - 1 {
		fmt.Println("stack full")
		return errors.New("栈满")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("stack empty")
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("栈空")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil

}

//判断一个字符是不是运算符[+ - * /]
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//运算的方法
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
		case 42:
			res = num2 * num1
		case 43:
			res = num2 + num1
		case 45:
			res = num2 - num1
		case 47:
			res = num2 / num1
		default:
			fmt.Println("运算符错误.")
	}
	return res
}

//编写一个方法，返回某个运算符的优先级[程序员定义]
func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		return 1
	} else if oper == 43 || oper == 45 {
		return 0
	} 
	return res
}

func main() {
	//数栈
	numStack := &Stack{
		MaxTop : 20,
		Top : -1,
	}
	//符号栈
	operStack := &Stack{
		MaxTop : 20,
		Top : -1,
	}

	exp := "30+30*6-4-6"
	//定义一个index，帮助扫描exp
	index := 0

	//为了配合运算，定义
	num1 := 0
	num2 := 0
	oper := 0

	result := 0

	keepNum :=""

	for {
		//我们需要增加一个逻辑，处理多位数的问题
		ch := exp[index : index+1]
		//把字符串ch转成int, "+"==>43
		temp := int([]byte(ch)[0])  //字符对应的ASCII码值

		if operStack.IsOper(temp) {
			//说明是操作符
			//如果operStack是一个空栈，就直接入栈
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				//如果发现operStack栈顶的运算符优先级大于等于当前准备入栈的运算符的优先级，
				//就从符号栈pop出，并从数栈中也pop两个数，进行运算，运算后的结果再重新入栈到
				//数栈，当前运算符号再入符号栈
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()

					result = operStack.Cal(num1, num2, oper)
					numStack.Push(result)
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}

			}

		} else {
			//是数字
			//处理多位数的思路
			//1.定义一个变量keepNum string,做拼接
			keepNum += ch
			//2.每次要向index的前面一位字符测试一下，看看是不是运算符，然后处理
			//如果已经到了表达式的最后
			if index == len(exp) - 1 {
				val, _ := strconv.ParseInt(ch, 10, 64)
				numStack.Push(int(val)) //3的ASCII码转成3
			} else {
				//向index后面测试看看 是不是运算符 [index]
				if operStack.IsOper(int([]byte(exp[index+1:index+2])[0])) {
					val, _ :=strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}

		//继续扫描，判断index是否已经扫描到了表达式的最后
		if index + 1 == len(exp) {
			break
		}
		index++		
	}

	//如果扫描表达式完毕，依次从符号栈取出符号，然后从数栈中取出两个数
	//运算后的结果入数栈，直到符号栈为空
	for {
		if operStack.Top == -1 {
			break //退出条件
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()

		result = operStack.Cal(num1, num2, oper)
		numStack.Push(result)
	}
	//如果我们的算法没有问题，表达式也是正确的，则结果就是numStack最后剩下的数
	res, _ := numStack.Pop()
	fmt.Println("结果是：", res)

}