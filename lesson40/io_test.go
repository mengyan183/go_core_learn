package lesson40

import (
	"io"
	"sync"
	"testing"
)

type customReader struct {
}

func (c *customReader) Read(p []byte) (n int, err error) {
	return len(p), nil
}

type customWriter struct {
}

func (c *customWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func TestIoCopy(t *testing.T) {
	// 对于io复制函数
	written, err := io.Copy(&customWriter{}, &customReader{})
	t.Log(written, err)
}

func TestLimitReader(t *testing.T) {
	reader := io.LimitedReader{
		N: 10,
	}
	// 对于read方法会对属性N进行校验,N小于等于0,表示已经不能读取;反之 则将读取数据,并将N减少读取的长度
	read, err := reader.Read(make([]byte, 5))
	t.Log(reader, read, err)
}

func TestSectionReader(t *testing.T) {
	startIndex := 0
	len := 10
	//对于SectionReader,可以限制指定开始读取位置以及最大读取长度
	// 类似于 切片,只允许在指定的区块中执行读取操作
	io.NewSectionReader(nil, int64(startIndex), int64(len))
}

func TestTeeReader(t *testing.T) {
	reader := io.TeeReader(nil, nil)
	read, err := reader.Read(make([]byte, 10))
	t.Log(read, err)
}

func TestMulti(t *testing.T) {
	io.MultiReader(nil, nil)
	io.MultiWriter(nil, nil)
}

// 实现建议的阻塞通道
func TestPipe(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		defer wg.Done()
		_, _ = pipeWriter.Write([]byte("i love go"))
		_ = pipeWriter.Close()
	}()
	go func() {
		defer wg.Done()
		b := make([]byte, 20)
		_, _ = pipeReader.Read(b)
		_ = pipeReader.Close()
		t.Log(string(b))
	}()
	wg.Wait()
}
