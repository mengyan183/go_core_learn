package lesson44

import (
	"os"
	"syscall"
	"testing"
)

func TestFile(t *testing.T) {
	// 对于指定的文件路径以及文件名进行操作
	// 如果路径中的文件夹不存在,则会抛出错误 //no such file or directory
	// 如果路径存在,但文件不存在,则会创建一个新的文件,且文件的权限为所有人可操作(0666)
	// 反之,如果文件存在,则会清除文件中的内容
	create, err := os.Create("./file/test.txt")
	t.Log(create, err)
	// 创建新的File指针类型值,第一个参数为文件类型的uintptr数据,第二个数据为文件的路径
	// 如果第一个文件类型的值不正确,则会返回nil,否则会返回一个File指针类型数据
	file := os.NewFile(uintptr(syscall.Stderr), "./file/test.txt")
	t.Log(file)
	// 以只读模式打开一个文件,如果文件不存在则会返回一个错误 // no such file or directory
	open, err := os.Open("./file/test.txt")
	t.Log(open, err)
	// 对只读模式的File指针数据进行写入操作,则会直接抛出错误 //bad file descriptor
	write, err := open.Write([]byte("1234"))
	t.Log(write, err)
	// 第一个参数表示文件的路径,第二个参数表示当前File指针类型值的权限, 第三个参数表示当前文件在操作系统中的权限
	// 对于额外的操作模式,必须搭配基础操作模式(只读/只写/读写),否则会抛出没有操作权限
	//openFile, err := os.OpenFile("./file/test.txt", os.O_RDWR|os.O_APPEND, 0666)
	openFile, err := os.OpenFile("./file/test.txt1", os.O_CREATE, 0666)
	t.Log(openFile, err)
	writeString, err := openFile.WriteString("1234")
	t.Log(writeString, err)


}

//TODO : 创建并操作一个系统进程
