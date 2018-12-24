package main
import (
	"fmt"
	"os"
)
//获取可执行文件的命令行参数

func main(){
	fmt.Println("命令行参数有：", len(os.Args))
	//遍历Args切片，得到所有的命令行参数
	for i, v := range os.Args {
		fmt.Printf("agrs[%v] = %v \n", i, v)
	}
}