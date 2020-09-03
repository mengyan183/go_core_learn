package lesson31

import (
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup // 在内部存在一个计数器
	wg.Add(2) // 原子操作
	for i := 0; i < 2; i++ {
		//go wg.Done()
	}
	wg.Add(-2) //当计数器不为0时,允许传递负值,但要求计数器最终结果值不能为负值,否则会抛出panic
	wg.Add(2)
	wg.Add(-2)
	wg.Wait()
}

func TestOnce(t *testing.T) {
	var o sync.Once
	o.Do(func() {
		t.Log("执行一次")
	})

	o.Do(func() {
		t.Log("第二次")
	})
}