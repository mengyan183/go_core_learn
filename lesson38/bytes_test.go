package lesson38

import (
	"bytes"
	"strconv"
	"testing"
)

func TestBytesBuffer(t *testing.T) {
	// 开箱即用属于懒加载
	var b bytes.Buffer
	b.WriteString("i love golang")
	// Len方法实际是返回的剩余未读取的字节长度
	i := b.Len()
	t.Log(i)
	c := b.Cap()
	t.Log(c)
	readRune, size, err := b.ReadRune()
	t.Log(readRune, size, err)
	// 当执行读取操作后,该值会发生变化
	i = b.Len()
	t.Log(i, b.String())
	b.Truncate(1)
	t.Log(b.String())

	// 读回退操作
	err = b.UnreadRune()
	t.Log(err)
}

func TestBufferGrow(t *testing.T) {
	var b bytes.Buffer
	// 默认容器容量为64
	for i := 0; i < 32; i++ {
		b.WriteString(strconv.Itoa(i))
	}
	t.Log(b.Len())
	// 当扩容长度小于当前剩余容量
	b.Grow(3)
	readBytes := make([]byte, b.Len()/2)
	read, err := b.Read(readBytes)
	t.Log(read, err)
	c := b.Cap()
	t.Log(c, b.Len())
	// 当扩容长度+未读数据长度 < 容量/2 且 扩容长度 > 剩余容量
	growLen := (c / 2) - b.Len()
	t.Log(growLen)
	b.Grow(growLen)
	// 生成新的切片
}
