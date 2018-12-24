package main
import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "/home/nanshantiger/文档/test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file err", err)
	}

	//fmt.Printf("%s", content) 效果一样
	fmt.Printf("%v", string(content))

	//不必显式open和close文件，readfile函数封装了Open和Close方法
}