package lesson4

import (
	"github.com/stretchr/testify/assert"
	. "go_core_learn/lesson4/other_pac"
	"io"
	"os"
	"testing"
)

var a = 1

//var A = 1 // 如果在把引入包作为当前包使用时,当存在同名全局变量时会编译出错提示变量重名,原因在于当把引入包作为当前包使用时,因此它的全局变量作用范围都是当前包,在相同的作用域下不能存在同名变量

// 变量声明

func TestDetermineParam(t *testing.T) {
	// 当变量的作用域不同时,可以出现重名
	var a = a
	{
		// 对于不同作用域的同名变量,其生效范围只在其作用域内,不会影响作用域外的同名变量
		a := 2
		assert.Equal(t, 2, a)
	}
	assert.Equal(t, 1, a)
	t.Log(a)
	var c int
	t.Log(c)
	b := 1
	t.Log(b)
	// 对于a/b而言,变量的类型已经通过自动推导确定,且go作为静态语言,因此在后续的操作中不能更改已经确定的类型
	//a = "1" // 编译期就会有编译错误
	t.Log(A)
	Hello()
	v, ok := interface{}(A).(int)
	t.Log(v, ok)

}

func TestReDetermine(t *testing.T) {
	var e error                                             // 声明一个变量
	n, e := io.WriteString(os.Stdout, "Hello, everyone!\n") // 利用短声明的方式,在声明一个新变量时对一个旧变量进行重声明操作
	t.Log(n)
	t.Error(e)
}
