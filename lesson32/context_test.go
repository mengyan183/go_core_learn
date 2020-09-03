package lesson32

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	// 创建context根节点
	rootContext := context.Background()
	// 创建具有撤销功能的context
	cancel, cancelFunc := context.WithCancel(rootContext)
	var total int64 = 10
	var sum int64 = 0
	var i int64 = 0
	for ; i < total; i++ {
		go func() {
			// 通过原子增操作判断是否达到执行上限
			if atomic.AddInt64(&sum, int64(1)) == total {
				// 执行撤销函数,为当前创建的cancel中channel中写入一个空的结构体实例
				cancelFunc()
			}
			t.Log("执行一次")
		}()
	}

	value := context.WithValue(cancel, 1, 1)
	t.Log(value.Value(1))
	t.Log(cancel.Err())
	// 当前channel会持续阻塞直到接收到消息或channel关闭为止
	<-cancel.Done()
	t.Log(cancel.Err())

	timeout, _ := context.WithTimeout(rootContext, time.Second)
	t.Log(timeout.Err())
	<-timeout.Done()
	t.Log(timeout.Err())
}

func TestValueSearch(t *testing.T) {
	root := context.Background()

	leftContext := context.WithValue(root, "left", "left")
	rightContext, cancelFunc := context.WithCancel(root)

	cancel, _ := context.WithCancel(rightContext)
	subContext := context.WithValue(leftContext, "right", "right")
	// 从下往上开始查找
	value := subContext.Value("left")
	t.Log(value)
	// 从上往下开始撤销
	cancelFunc()
	t.Log(cancel.Err())
}
