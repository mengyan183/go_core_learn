package lesson10

import (
	"math/rand"
	"sync"
	"testing"
)

func TestChannelMutex(t *testing.T) {
	c := make(chan int, 5)
	var wg sync.WaitGroup
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			c <- i
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	for {
		v, ok := <-c
		if !ok {
			break
		}
		t.Log(v)
	}
}

func TestNilChannel(t *testing.T) {
	var c chan int
	c <- 1
	t.Log(<-c)
}

// channel阻塞
func TestMutexChannel(t *testing.T) {
	c := make(chan int, 1)
	c <- 1
	//c <- 2 // 由于channel的容量为1,当channel已满时,再往channel中写入数据就会导致阻塞

	//c1 := make(chan int)
	//t.Log(<-c1)// 由于channel中一直没有数据,所以消费者需要一直阻塞等待

}

// 限定当前函数只能往当前channel中写入数据
func onlyInChannel(inChan chan<- int) {
	defer close(inChan)
	c := cap(inChan)
	for i := 0; i < c; i++ {
		inChan <- rand.Intn(100)
	}
}

func TestOnlyChannelUse(t *testing.T) {
	ch := make(chan int, 10)
	// 双向通道可以直接传递给单向通道
	onlyInChannel(ch)
	//for {
	//	v, ok := <-ch
	//	if !ok {
	//		break
	//	}
	//	t.Log(v)
	//}
	// 使用range 获取channel中的数据,当channel关闭后且channel中没有数据时,range遍历结束后就会自动结束
	for v := range ch {
		t.Log(v)
	}
}
