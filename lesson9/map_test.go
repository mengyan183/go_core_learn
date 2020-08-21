package lesson9

import "testing"

func TestKeyType(t *testing.T) {
	//m := make(map[interface{}]int)
	//m[[]int{1, 2}] = 1 // 这里会抛出运行时异常
	//m[[1][]int{[]int{1, 2, 3}}] = 1 // key为数组类型,但数组中的元素类型为切片类型,会抛出运行时异常

	var m1 map[int]int
	t.Log(m1)    //m1的值为nil
	t.Log(m1[1]) // 支持读取等操作
	m1[1] = 1     // 会抛出运行时panic, 不支持写入操作
}
