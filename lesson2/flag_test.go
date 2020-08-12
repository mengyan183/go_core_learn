package lesson2

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var s *string

// 方式3: 定义全局命令行解析变量
var cmd = flag.NewFlagSet("", flag.ExitOnError)

func init() {

	// 解析启动命令中的name 参数,如果不存在参数,则参数的默认值为everyOne
	//s = flag.String("name", "everyOne", "terminal input")
	// 方式3 使用自定义的全局变量替换flag
	s = cmd.String("name", "everyOne", "terminal input")

	// 方式2 : 通过修改flag 内部变量 CommandLine,实现修改Usage的操作
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError) // 第二个参数代表的是异常处理机制,属于flag内部已经预设好的几种处理机制
	// 设置Usage参数
	flag.CommandLine.Usage = func() {
		// 这里会修改命令行中参数的用处注释
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", "flag.NewFlag")
		flag.PrintDefaults()
	}
}

// 从启动命令获取输入数据
// 命令行启动 go test -v -run TestGetDataFromCommand flag_test.go -name=命令行输入的数据
func TestGetDataFromCommand(t *testing.T) {
	// 方式1
	// 需要定义在Parse函数之前
	//flag.Usage = func() {
	//	// 这里会修改命令行中参数的用处注释
	//	_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", "flag")
	//	flag.PrintDefaults()
	//}
	// 解析启动命令中的数据
	//flag.Parse() // 这里内部实际调用的是CommandLine.Parse
	// 方式3
	_ = cmd.Parse(os.Args[1:])
	t.Log(*s)
}
