package lesson26

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var i = 0

func TestShare(t *testing.T) {
	//i := 0
	var wg sync.WaitGroup
	n := 1000 //当n越大,抛出错误的几率越大
	wg.Add(n)
	for a := 0; a < n; a++ {
		go func() {
			i++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(i)
	assert.Equal(t, n, i) // 这里有可能会抛出错误
}

func TestShareWithLock(t *testing.T) {
	var mu sync.Mutex
	var wg sync.WaitGroup
	n := 1000 //当n越大,抛出错误的几率越大
	wg.Add(n)
	for a := 0; a < n; a++ {
		go func() { //当对当前代码进行加锁操作保证操作共享变量的协程互斥性,达到共享变量的一致性
			mu.Lock()
			defer mu.Unlock()
			i++
			wg.Done()
		}()
	}
	wg.Wait()
	assert.Equal(t, n, i)
}

func TestRLockUnlock(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	var rw sync.RWMutex
	for a := 0; a < 10; a++ {
		go func() {
			defer func() {
				recover() // 对于运行时系统抛出的panic不允许捕获
			}()
			rw.RUnlock() // 对未加锁的锁进行解锁操作会立即引发panic导致程序终止
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestRWLock(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	var rw sync.RWMutex
	for a := 0; a < 10; a++ {
		if a%2 == 0 {
			go func() {
				rw.RLock()
				defer rw.RUnlock()
				t.Log(i) // 由于读锁之间不会互斥,因此存在并发读的情况,会读取到相同的数据
				wg.Done()
			}()
		} else {
			go func() {
				rw.Lock()
				defer rw.Unlock()
				i++
				wg.Done()
			}()
		}

	}
	wg.Wait()
}
