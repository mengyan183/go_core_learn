package lesson12

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 定义函数签名
type operate func(x, y int) int

func add(x, y int) int {
	return x + y
}
func by(x, y int) int {
	return x * y
}

func calculate(x, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

func TestOperate(t *testing.T) {
	i, err := calculate(1, 2, add)
	if err != nil {
		t.Error(err)
	} else {
		assert.Equal(t, 3, i)
	}
}

type calculateFunc func(x, y int) (int, error)

// 构建闭包函数, 对于closure 返回的函数而言,由于引用外部数据 op,导致了函数的不确定性,因此称为闭包函数
func closure(op operate) calculateFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func TestClosure(t *testing.T) {
	i, err := closure(add)(1, 2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 3, i)
}
