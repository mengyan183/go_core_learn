package lesson15

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

type Named interface {
	// Name 用于获取名字。
	Name() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func New(name string) Dog {
	return Dog{name: name}
}

func TestUnReachAble(t *testing.T) {
	//New("test").SetName("change")//cannot take the address of New("test")
	//对于New("test")实际是临时结果属于不可寻址值,对于不可寻址的值是不能操作的
	d := New("test")
	dptr := uintptr(unsafe.Pointer(&d))
	t.Log(dptr)
	t.Log(uintptr(unsafe.Pointer(&d.name))) //这里实际是获取的是当前结构体实例中name引用的字符串的地址
	t.Log(unsafe.Offsetof(d.name))
	//assert.Equal(t, unsafe.Offsetof(d.name), uintptr(unsafe.Pointer(&d.name))-dptr)
	namePtr := dptr + unsafe.Offsetof(d.name)
	t.Log(namePtr)
	nameP := (*string)(unsafe.Pointer(namePtr))
	t.Logf("%v", nameP)
	t.Log(*nameP)

	New("test").Name()
}

// 通过指针值获取切片
func TestGetPrivate(t *testing.T) {
	s := []int{1, 2, 3}
	sptr := uintptr(unsafe.Pointer(&s))
	t.Log(sptr)
	arrPtr := (*[]int)(unsafe.Pointer(sptr))
	arr := *arrPtr
	t.Logf("%v,%T", arr, arr)
	lenPtrAddr := sptr + unsafe.Sizeof(unsafe.Pointer(sptr))
	lenPtr := (*int)(unsafe.Pointer(lenPtrAddr))
	assert.Equal(t, len(s), *lenPtr)
}
