package lesson25

import "testing"

func TestFunction(t *testing.T) {
	// 功能测试
	for i := 0; i < 10; i++ {
		t.Log(i)
	}
}

//go test -bench=. -run=BenchmarkPerformance -cpu=1,2,4 -count 限制当前测试函数执行次数
// 当设置-cpu为整数列表,在执行基准测试时,会按照整数列表依次设置当前基准测试最大P数量,并依次执行基准测试
func BenchmarkPerformance(b *testing.B) {
	// 基准测试
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			//b.Log(i)
		}
	}
}
