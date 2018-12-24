package main
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	filePath := "/home/nanshantiger/文档/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err%v\n", err)
		return
	}
	//及时关闭文件
	defer file.Close()

	str := "Hello, Gardon\r\n"  //支持不同编辑器的换行
	//写入时使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i :=0; i < 5; i++ {
		writer.WriteString(str)
	}

	writer.Flush()

}