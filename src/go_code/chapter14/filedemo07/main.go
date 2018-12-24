package main
import (
	"fmt"
	"io/ioutil"

)

func main() {
	file1Path := "/home/nanshantiger/文档/abc.txt"
	file2Path := "/home/nanshantiger/文档/test.txt"

	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("read file error:", err)
		return
	} 

	//func WriteFile(filename string, data []byte, perm os.FileMode) error
	error := ioutil.WriteFile(file2Path, data, 0666)
	if error != nil {
		fmt.Println("Write file error:", error)
		return
	}
	fmt.Println("操作成功")
}