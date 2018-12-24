package main
import (
	"fmt"
	"os"
	"bufio"
)

//打开已存在的文件，用10个字符串“你好”，替换文件中的内容

func main() {
	filePath := "/home/nanshantiger/文档/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err%v\n", err)
		return
	}
	//及时关闭文件
	defer file.Close()

	str := "你好，北京昌平\r\n"  //支持不同编辑器的换行
	//写入时使用带缓存的*writer
	writer := bufio.NewWriter(file)
	for i :=0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()

}