package lesson10

import (
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
