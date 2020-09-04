package lesson36

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	s := "go学习"
	// 这里实际获取的字节切片的长度
	t.Log(len(s),len([]rune(s)))
	// 这里实际遍历的是字节切片
	for i, r := range s {
		// 通过打印i 可以看出遍历的不为字符切片,而是字节切片
		//在遍历字节切片时,如果一组字节有对应的Unicode字符,则会将对应的字符返回给变量
		t.Log(i, r)
	}
	rs := []rune(s)
	t.Log(rs)

	str := "Go爱好者"
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	// 对比字节和UTF8字符 切片可以发现,对于中文字符在使用字节进行表示时使用了三个字节进行保存
	// 爱 : 7231(UTF8字符) [e7 88 b1](字节)
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))
}
