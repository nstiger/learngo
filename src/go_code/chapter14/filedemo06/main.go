package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)

//打开已存在的文件，将原来的内容显示到终端，并在原来的内容后面追加“hello，北京！”

func main() {
	filePath := "/home/nanshantiger/文档/abc.txt"

	file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err%v\n", err)
		return
	}
	//及时关闭文件
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	str := "Hello, 北京！\r\n"  //支持不同编辑器的换行
	//写入时使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i :=0; i < 5; i++ {
		writer.WriteString(str)
	}

	writer.Flush()

}