package lesson8

import (
	"container/list"
	"testing"
	"unsafe"
)

func TestZeroValue(t *testing.T) {
	var a []int
	var b [1]int
	var c [1]int
	t.Log(unsafe.Pointer(&a))
	t.Log(unsafe.Pointer(&b))
	t.Log(unsafe.Pointer(&c))
}

func TestListDelayInit(t *testing.T) {
	var l list.List
	t.Log(l, unsafe.Pointer(&l))
	l.Init()
	t.Log(l, unsafe.Pointer(&l))
	a := 1
	t.Log(unsafe.Pointer(&a))
	l.PushBack(a)
	t.Log(l, unsafe.Pointer(&l))
}

func TestMoveCustomEle(t *testing.T) {
	var l list.List
	e := list.Element{Value: 1}
	// 这里无法直接操作element,对于element的元素中都包含一个私有list属性,表示当前ele属于哪个list,由于是私有属性,所以无法直接修改ele中list指针的值,而在List暴露出的直接操作Ele的方法中,都增加了 ele中的list必须等价于当前操作的list 限制,因此使用一个list操作一个自定义生成的ele是无法操作成功的
	l.MoveToBack(&e)
	t.Log(l, e)
	t.Log(unsafe.Pointer(&l))
	neptr := l.PushBack(1)
	t.Log(unsafe.Pointer(neptr), unsafe.Pointer(&l))
	t.Log(neptr)
	l.MoveToBack(neptr)
	neptr = l.PushBack(1)
	t.Log(unsafe.Pointer(neptr), unsafe.Pointer(&l))
	t.Log(neptr)
}
