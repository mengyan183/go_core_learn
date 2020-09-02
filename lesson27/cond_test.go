package lesson27

import (
	"sync"
	"testing"
	"time"
)

// 利用条件变量实现协调多协程发取信件操作
func TestCond(t *testing.T) {
	var wg sync.WaitGroup
	var mu sync.RWMutex
	// 信箱
	mail := false
	// 两个条件变量
	// 发送信条件变量
	sendCond := sync.NewCond(&mu)
	// 接收信条件变量, 对于接收实际是只读操作,因此只需要使用读锁就可以
	receiveCond := sync.NewCond(mu.RLocker())
	// 最大发送接收次数
	max := 5
	wg.Add(2)
	// 发送人协程
	go func(i int) {
		for ; i > 0; i-- {
			time.Sleep(time.Second * 3)
			mu.Lock() // 这里的lock表示的是拥有访问当前mail的权利
			// 如果信箱不为空,则需要等待
			//for mail { // 这是使用for 的原因在于如果共享变量存在多个值,即使当前阻塞协程被唤醒,但并不能表示当前条件满足,因此需要再次做判断; 即使共享变量不存在多个值,如果全部被阻塞的协程被唤醒,也只可能有一个协程成功进入临界区,不可能允许多个协程同时进入临界区,否则就有可能导致panic; 且协程的唤醒不只能通过程序唤醒,操作系统也可以直接唤醒阻塞的协程
			if mail {
				// 发送者等待
				t.Log("sendCond准备进入等待队列")
				sendCond.Wait()
				t.Log("sendCond进入等待队列")
			}
			mail = true
			t.Log("发送信件成功")
			mu.Unlock()
			// 通知发送者
			receiveCond.Signal()
			t.Log("唤醒receiveCond")
		}
		wg.Done()
	}(max)

	go func(i int) {
		for ; i > 0; i-- {
			mu.RLock()
			//for !mail {
			if !mail {
				//接收者等待
				t.Log("receiveCond准备进入等待队列")
				receiveCond.Wait() // 如果没有被唤醒会一直阻塞在此
				t.Log("receiveCond进入等待队列")
			}
			mail = false
			t.Log("获取信件成功")
			mu.RUnlock()
			// 通知接收者
			sendCond.Signal()
			t.Log("唤醒sendCond")
		}
		wg.Done()
	}(max)
	wg.Wait()
}

func TestCondWaitAndSignal(t *testing.T) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	con := sync.NewCond(&mu)
	wg.Add(2)
	go func() {
		for true {
			mu.Lock()
			t.Log("con准备进入持续阻塞等待状态")
			con.Wait()// 如果没有signal ,当前con会一直阻塞在此
			t.Log("con进行等待状态.......")
			mu.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for true {
			time.Sleep(time.Second * 5)
			mu.Lock()
			t.Log("准备通知con")
			con.Signal()
			t.Log("通知con结束")
			mu.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
}
