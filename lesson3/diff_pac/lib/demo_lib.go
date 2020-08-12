package lib_re

import (
	"fmt"
	"go_core_learn/lesson3/diff_pac/lib/internal"
)

func Hello() {
	fmt.Println("hello")
}
// 对于小写开头的程序实体,只能在当前包下使用,即使引入当前包也不能使用当前包下的私有函数
func hello() {
	internal1.HelloInternal()
}
