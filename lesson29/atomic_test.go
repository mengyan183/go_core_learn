package lesson29

import (
	"github.com/stretchr/testify/assert"
	"go/types"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var a int64 = 0
	// 原子读取
	atomic.LoadInt64(&a)
	// 原子赋值
	atomic.StoreInt64(&a, 1)
	// 原子自增
	atomic.AddInt64(&a, 1)
	t.Log(a)
	atomic.CompareAndSwapInt64(&a, a, 3)
	t.Log(a)

	var b uint64
	atomic.StoreUint64(&b, 10)
	// 4 实际为要减去的值
	atomic.AddUint64(&b, ^uint64(4-1))
	assert.Equal(t, uint64(6), b)

	t.Log(^uint8(1))
	t.Log(^uint8(3))
	var c = uint64(10)
	assert.Equal(t, uint64(10-4), atomic.AddUint64(&c, (^uint64(10-6-1))))
	t.Log(c + (^uint64(10 - 6 - 1)))
}

func TestCas(t *testing.T) {
	var a int64 = 0
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for !atomic.CompareAndSwapInt64(&a, 3, 4) {
			t.Log("自旋")
			time.Sleep(time.Second * 1)
		}
		wg.Done()
	}()
	for i := 0; i < 3; i++ {
		a++
		time.Sleep(time.Second)
	}
	wg.Wait()
}

func TestAtomicValue(t *testing.T) {
	var v atomic.Value
	v.Store(1)
	//v.Store(int64(1)) // 由于类型不一致会引发panic异常
	//v.Store(nil) // 不允许存储nil,会引发panic

	// 当使用atomic时一般推荐使用结构体
	type customValue struct {
		v atomic.Value
		t types.Type // 原子值的真实数据类型
	}
}
