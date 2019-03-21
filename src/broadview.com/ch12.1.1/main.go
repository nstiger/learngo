package main

import (
	"log"
	"strings"
	"os"
	"fmt"
	"io"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil 
	}
	return p, err
}

func main() {
	//从标准输入读取
	data, err := ReadFrom(os.Stdin, 11)
	if err != nil {
		fmt.Println("读取数据错误")
	}

	fmt.Printf("标准输入为：%v\n", string(data))

	//从字符串读取
	data, err = ReadFrom(strings.NewReader("from string"), 12)
	fmt.Printf("从字符串读取：%v\n", string(data))

	//从文件读取
	file, err := os.Open("/home/nanshantiger/111.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err = ReadFrom(file, 100)
	file.Close()
	fmt.Printf("从文件读取：%v\n", string(data))

}