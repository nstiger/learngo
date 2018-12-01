package main
import (
	"fmt"
)

func main() {
	//字符大小写转换
	// var ch byte;
	// fmt.Println("请输入一个字符")
	// fmt.Scanf("%c", &ch)    //应该格式化接收！！！

	// switch ch {
	// case 'a': fmt.Println("A")
	// case 'b': fmt.Println("B")
	// case 'c': fmt.Println("C")
	// case 'd': fmt.Println("D")
	// case 'e': fmt.Println("E")
	// default : fmt.Println("other...")
	//}


	//根据用户指定的月份，打印该月份所属的季节。3、4、5春季，6、7、8夏季，9、10、11秋季，12、1、2冬季
	var month int8
	fmt.Println("请输入当前月份。。。。。。")
	fmt.Scanln(&month)

	switch month {
	case 3,4,5: fmt.Println("春季")
	case 6,7,8: fmt.Println("夏季")
	case 9,10,11: fmt.Println("秋季")
	case 12,1,2: fmt.Println("冬季")
	default : fmt.Println("输入有误")
	}
	

}