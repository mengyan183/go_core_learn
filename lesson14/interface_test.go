package lesson14

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type pet interface {
	eat()
}

type cat struct {
	name  string
	color string
}

// 方法签名一致且方法名一致,称cat为pet的实现类型
func (c *cat) eat() {
	fmt.Printf("名字为%s且颜色为%s的猫吃东西", c.name, c.color)
}

func TestInterfaceNil(t *testing.T) {
	var p pet
	assert.Equal(t, nil, p)
	p = nil
	assert.Equal(t, nil, p)
	var c cat
	t.Log(c) // 不为nil
	var cptr *cat
	assert.True(t, cptr == nil)
	p = cptr
	assert.False(t, p == nil)
	// 对于接口类型变量和接口实现类型变量进行比较时,如果接口变量的类型和当前实现类型一致,则会比较接口变量动态值和当前实现类型变量的值
	assert.True(t, p == cptr)
}
