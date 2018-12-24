package main
import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpu的数量 =", cpuNum)

	//可以自己设置使用几个cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}