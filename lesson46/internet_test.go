package lesson46

import (
	"context"
	"net"
	"net/http"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestNetDial(t *testing.T) {
	dial, err := net.Dial("tcp", "127.0.0.1:80")
	t.Log(dial, err)
	// 第一个参数 为 socket实例通信域
	// 第二个参数 为 socket的类型
	// 第三个参数 为 socket的协议,对于第三个参数,当第一二个参数指定好之后,当第三个参数为0时,内核程序会自动选择合适的协议
	socket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	t.Log(socket, err)
	// 这里的timeout表示当前连接建立的最大超时时间
	timeout, err := net.DialTimeout("tcp", "192.168.89.61:9200", time.Second)
	t.Log(timeout, err)
	if timeout != nil {
		// 设置当前连接的
		err := timeout.SetDeadline(time.Time{})
		t.Log(err)
	}
}

func TestHttp(t *testing.T) {
	// 这里直接调用实际有内置http.Client的缺省值
	get, err := http.Get("http://www.baidu.com")
	t.Log(get,err)
	// 支持开箱即用
	var c http.Client
	resp, err := c.Get("http://www.google.cn")
	t.Log(resp,err)
	c =http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0, // 当前client的超时限制包含连接建立时间以及请求时间; 当数据为0时,表示永远不是超时
	}
	server := http.Server{
		Addr:              "",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	err = server.ListenAndServe()
}

func TestServer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	server := http.Server{
		Addr: "http://www.baidu.com",
	}
	server.RegisterOnShutdown(func() {
		t.Log("server关闭前的回调")
		wg.Done()
	})
	err := server.Shutdown(context.Background())
	t.Log(err)
	wg.Wait()
}
