package lesson42

import (
	"bufio"
	"bytes"
	"testing"
)

func TestReader(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBufferString("i love go"))
	readByte, err := reader.ReadByte()
	t.Log(readByte, err)
	// 指定缓冲区切片长度,最小切片长度为16
	reader = bufio.NewReaderSize(bytes.NewBufferString("i love go"), 2)
	_, _ = reader.ReadByte()
	readBytes := make([]byte, 3)
	// 当读取数据长度超过切片的容量时,有可能发生直接调用数据 ,当存在 b.r == b.w 切片中没有数据,会切片中的数据全部都已经被读完时,如果要读取的长度超过切片的长度,就会直接调用底层读取器Reader直接读取数据
	read, err := reader.Read(readBytes)
	t.Log(string(readBytes),read, err)
	read, err = reader.Read(readBytes)
	t.Log(string(readBytes),read, err)
}

func TestWriter(t *testing.T) {
	writer := bufio.NewWriter(bytes.NewBufferString(""))
	write, err := writer.Write(make([]byte, 5))
	t.Log(write,err)
}

func TestOtherReader(t *testing.T) {
	reader := bufio.NewReader(bytes.NewBufferString("i love go"))
	// 并不会修改已读计数器
	peek, err := reader.Peek(0)
	t.Log(string(peek),err)

	reader.ReadBytes(1)
	reader.ReadSlice(1)
}
//TODO : scanner学习
func TestScanner(t *testing.T) {
	scanner := bufio.NewScanner(bytes.NewBufferString("i love go"))
	scan := scanner.Scan()
	if scan{
		text := scanner.Text()
		t.Log(text)
	}
}
