package lesson37

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStringsBuilder(t *testing.T) {
	// 此时builder 还未被完全初始化,可以被复制使用
	var builder strings.Builder
	// 此时发生真正的初始化操作,此时不允许被复制使用
	// copyCheck 在该方法中会首先校验结构体中的 指针,如果指针为nil表示首次使用,并进行赋值操作
	// 如果指针不为nil,且指针不等于当前调用者的指针, 会抛出panic
	builder.WriteString("1")
	// 在复制操作之前先进行Reset操作,将其恢复为zero-value
	//builder.Reset()
	bu1 := builder
	bu1.Reset()
	bu1.WriteString("2") // 如果不存在reset操作,这里会抛出panic

	bu1ptrcopy := &bu1
	bu1ptrcopy.WriteString("3")
	assert.Equal(t, bu1.String(), bu1ptrcopy.String())
}

func TestStringsReader(t *testing.T) {
	reader := strings.NewReader("go123456")
	readByte, err := reader.ReadByte()
	t.Log(readByte, err)
	// 对于reader.Len实际就是通过(size - 计数器)得到的
	i := reader.Size() - int64(reader.Len())
	// 当前计数器
	t.Log(i)
	// 这里返回的是copy函数的结果
	at, err := reader.ReadAt([]byte("go123"), 3)
	t.Log(at)
	t.Log(reader.Size())
}
