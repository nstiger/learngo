package main

import (
	"fmt"
	"broadview.com/ch8.3.3/filestruct"
)

func main() {

	f := filestruct.Newfile1(10, "storm")

	fmt.Printf("f type: %T ; file fd: %d, file name : %v\n", f, f.Fd, f.Filename)
	//fmt.Println("file is :", f)
}