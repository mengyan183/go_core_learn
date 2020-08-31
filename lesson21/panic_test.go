package lesson21

import "testing"

func TestPanic(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	//panic: runtime error: index out of range [4] with length 4 [recovered]
	//	panic: runtime error: index out of range [4] with length 4  引发panic的详细错误信息
	// /go_core_learn/lesson21/panic_test.go:7 +0x1b // 当前代码相对于所属函数入口程序偏移量
	t.Log(arr[4]) // 会抛出panic
	panic("错误信息")
}

func TestPanicRecovery(t *testing.T) {
	defer func() {
		msg := recover()
		t.Log(msg)
	}()
	panic("错误信息")
}

func TestMultiDefer(t *testing.T) {
	for i := 0; i < 10; i++ {
		defer func(i int) { // defer 语句执行时会被添加到当前函数关联的链表中,使用的数据结构为栈,满足先进后出的原则
			// 在当前函数返回时会调用当前函数关联的defer函数链表,按照函数出栈规则依次执行
			// 对于defer右侧的函数中的参数列表,在当前延迟函数压入栈时会同时将参数一起压入栈(参数发生值复制)
			t.Log(i)
		}(i)
	}
}
