package main
import (
	"fmt"
	"os"
	"bufio"
)

//打开已存在的文件，在原来的内容后面追加“ABC，ENGLISH”

func main() {
	filePath := "/home/nanshantiger/文档/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err%v\n", err)
		return
	}
	//及时关闭文件
	defer file.Close()

	str := "ABC，ENGLISH\r\n"  //支持不同编辑器的换行
	//写入时使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i :=0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()

}