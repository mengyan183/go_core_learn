package lesson12

import (
	"fmt"
	"testing"
)
// 声明一个函数签名
type Printer func(string) (int, error)

func printToStd(s string) (i int, e error) {
	return fmt.Println(s)
}

func TestFuncSign(t *testing.T) {
	// 对于Printer类型对应的函数签名和printToStd函数实际签名是相同的,因此可以看做为相同的函数类型
	var p Printer
	p = printToStd // 由于函数签名相同,因此可以直接将函数作为值进行传递
	_, _ = p("测试")
}
