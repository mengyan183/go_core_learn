package lesson33

import (
	"runtime"
	"sync"
	"testing"
)

type cache struct {
	value int
}

func TestPool(t *testing.T) {
	// 由于sync.Pool并不是开箱即用,因此在创建一个新的缓存池时,要求指定一个初始化函数来限制当前缓存池存储的具体数据类型;否则如果缓存池中没有数据,则会直接返回nil
	p := sync.Pool{
		New: func() interface{} {
			return new(cache)
		},
	}
	get := p.Get() // 实际是执行的删除操作,并返回删除的值
	t.Log(get)
	// 对于private ,只有和当前缓存池属于同一个协程的Put操作才有可能把数据写入;
	p.Put(new(cache)) // 在执行put操作时
}

func TestGoMaxP(t *testing.T) {
	runtime.GOMAXPROCS(5)
	// 这里是获取运行时当前协程最大数量
	num := runtime.GOMAXPROCS(0)
	t.Log(num)
}
//TODO : 如何体现出来private 和 share
func TestShareAndPrivate(t *testing.T) {
	p := sync.Pool{}
	// 在主协程写入10
	p.Put(cache{value: 10})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			p.Put(cache{value: i})
		}
		wg.Done()
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		for true {
			v := p.Get()
			if v == nil {
				break
			}
			t.Log(v)
		}
		wg.Done()
	}()
	wg.Wait()
}
