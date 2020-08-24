package lesson10

import (
	"math/rand"
	"testing"
)

func TestSelectChannel(t *testing.T) {
	chArr := []chan int{
		make(chan int, 10),
		make(chan int, 10),
		make(chan int, 10),
	}
	// 当channel中没有数据时或为nil,select会执行default;
	for index, ch := range chArr {
		c := cap(ch)
		for i := 0; i < c; i++ {
			ch <- rand.Intn(100)
		}
		// 当channel为nil时,select case不会执行阻塞的channel
		//chArr[index] = nil
		t.Log(index)
		// 对于修改ch变量,实际并不会修改数组存储的channel; 且对于range返回的值变量对于for循环而讲实际是一个全局变量
		// 且修改ch变量为nil,只是修改了当前ch变量的指向并不会真正修改数组中存储的channel变量真正的指向
		//ch = nil
	}
	for _, ch := range chArr {
		t.Log(ch)
	}
	// 如果channel为nil,对应的case是不会被执行到的
	select {
	// default 的执行顺序和其代码位置无关,只有case都不能匹配时,default才会执行
	default:
		t.Log("执行默认")
	case v := <-chArr[0]:
		t.Log(v)
	case v := <-chArr[1]:
		t.Log(v)
	case v := <-chArr[2]:
		t.Log(v)
		//当不存在default,会阻塞等待其中任意一个case对应的channel执行;
		//当存在default时,按照case代码执行顺序,如果以上case都被阻塞,则会执行default

	}
}

func TestSelect(t *testing.T) {
}
