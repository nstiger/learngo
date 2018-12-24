package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)

//统计文件中出现的英文字母、数字、空格和其他字符的个数
//思路：打开一个文件，创建一个Reader
//每读取一行，就统计该行有多少个英文字母、数字、空格和其他字符的个数
//然后将结果保存到一个结构体

//定义一个结构体，用于保存统计结果
type CharCount struct{
	Chcount int
	Numcount int
	Spacecount int
	Othercount int
}

func main() {
	fileName := "/home/nanshantiger/文档/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}

	defer file.Close()    //延迟关闭文件

	//定义CharCount的实例
	var count CharCount

	//创建一个Reader
	reader := bufio.NewReader(file)
	//开始循环读取fileName的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//遍历str，进行统计
		//为了兼容中文字符，可以将str转为[]rune
		//str = []rune(str)
		for _, v := range str {
			switch  {
			case v >= 'a' && v <= 'z':
				fallthrough  //穿透
			case v >= 'A' && v <= 'Z':
				count.Chcount++
			case v == ' ' || v == '\t':
				count.Spacecount++
			case v >= '0' && v <= '9' :
				count.Numcount++
			default:
				count.Othercount++
			}
		}
	}

	//输出统计结果
	fmt.Println("————————统计结果————————")
	fmt.Println("文件中包含英文字母：", count.Chcount)
	fmt.Println("文件中包含数字：", count.Numcount)
	fmt.Println("文件中包含空格：", count.Spacecount)
	fmt.Println("文件中包含其他字符：", count.Othercount)

	//linux ubuntu， 换行符为2个字符，中文一个汉字为1个字符

}