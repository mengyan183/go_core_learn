package lesson16

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

// 当不加 i:=i 时,输出的结果全部为10;当存在 i:=i时,输出的结果为 [0,9]
// 从变量赋值内存以及协程角度解读
// 由于异步协程的执行需要等待主协程释放占用资源时(wg.Wait阻塞),子协程才会开始执行
// 在执行协程时才会执行入栈操作: 当不存在 i:=i时,对于协程方法而言,i为全局变量,此时变量i的值为10,因此输出结果全部为10
// 当 存在 i:=i时,每次循环都是将i对应的副本赋值给新的变量i,此时变量i的值是不会更改的,因此当使用i时,i的值是[0,9]
func TestGo(t *testing.T) {
	var wg sync.WaitGroup
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		//i:=i
		go func(i int) {
			t.Log(unsafe.Pointer(&struct {

			}{}))
			t.Log(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestOrderGo(t *testing.T) {
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {})
}
