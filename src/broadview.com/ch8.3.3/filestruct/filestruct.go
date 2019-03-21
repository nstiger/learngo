/*
工厂模式输出file1类型
*/
package filestruct

type file1 struct {
	Fd int
	Filename string
}

func Newfile1(fd int, name string) *file1 {
	if fd < 0 {
		return nil 
	}
	return &file1{fd, name}
}