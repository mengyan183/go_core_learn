package lesson17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 对于range 而言,右侧的数据实际就是参数传递,因此在遍历过程中实际是操作的数据副本
// 而如果在for代码块中使用遍历的数据,实际操作的不是传入的副本,而是原始数据
func TestRangeArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	for i, v := range a {
		a[i] += v
	}
	t.Log(a)
}

func TestRangeSlice(t *testing.T) {
	s := []int{1, 2, 3}
	for i, v := range s {
		if i == 0 {
			s[i] += v
		} else {
			s[i] += s[i-1]
		}
	}
	t.Log(s)

	numbers2 := []int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	t.Log(numbers2) // [22,3,6,10,15,21]
	assert.Equal(t, numbers2, []int{22, 3, 6, 10, 15, 21})
}

// 测试当switch表达式为无类型常量时,自动类型转换
func TestSwitchConstant(t *testing.T) {
	//arr := [...]int64{1, 2, 3, 4}
	switch 1 {
	//case arr[0], arr[1], arr[2]: // 会抛出编译失败,原因在于 switch表达式为无类型常量,"1"被自动转换为默认类型(int),而由于case表达式结果值类型为(int64);由于 switch 表达式的类型和case类型不一致因此编译不通过

	}
}

func TestCaseConstant(t *testing.T) {
	arr := [...]int8{1, 2, 3, 4}
	switch arr[0] {
	case 1: // 编译通过,但运行时会抛出panic异常 (mismatched types int64 and int),

	}
}

func TestCaseSameConstantVal(t *testing.T) {
	switch 1 {
	case 1, 2, 3: // 对于case表达式中,使用字面数据,如果出现重复的字面数据,会直接编译不通过
	//case 2, 3, 4:
	default:

	}
}

func TestCaseSameVal(t *testing.T) {
	arr := [...]int{1, 1, 1, 2, 2, 2}

	switch 1 {
	case arr[0], arr[3]:// 对于case表达式中如果出现结果值出现重复,允许编译通过也会执行通过
		t.Log(1)
	case arr[1], arr[4],1:
		t.Log(2)
	}
}
