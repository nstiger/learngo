package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	file, err := os.Open("/home/nanshantiger/文档/我的文件.txt")
	if err !=nil {
		fmt.Println("open file err=", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		fmt.Print(str)
	}
	fmt.Println("文件读取结束！")

}