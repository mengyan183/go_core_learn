package lesson40

import (
	"bytes"
	"testing"
)

// io核心接口 io.Reader/Writer/Closer

// 扩展 io.ByteReader
func TestIoReaderImpl(t *testing.T) {
	// 对于其实现有buffer中的数据读取
	bufferString := bytes.NewBufferString("i love go")
	// 对于buffer中的ReadByte实际就是io.ByteReader的实现类型
	readByte, err := bufferString.ReadByte()
	t.Log(readByte, err)
	// 对于buffer中的ReadRune实际就是io.RuneReader的实现
	readRune, size, err := bufferString.ReadRune()
	t.Log(readRune, size, err)
	// 对于UnreadRune实际就是对于io.RuneScanner的实现
	err = bufferString.UnreadRune()

}
