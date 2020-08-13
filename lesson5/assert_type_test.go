package lesson5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// aliases 类型别名
type AliasString = string

// 类型重定义
type ReString string
type ReString2 string
type ReStringCopy ReString

func TestUnderlyingType(t *testing.T) {
	s := "1"
	rs := ReString(s)
	rs2 := ReString2(rs)
	rsc := ReStringCopy(s)
	t.Log(rs, rs2, rsc)
	// 对于潜在类型相同的数据只支持类型强制转换以及类型断言; 对于赋值/对比等操作会抛出编译错误
	//t.Log(rsc == rs)

	//相同潜在类型的 切片类型
	var reStringSlice []ReString
	var reString2Slice []ReString2
	t.Log(reString2Slice, reStringSlice)
	// 不支持集合类型强制转换,会抛出编译错误,原因在于对于集合类型的数据其潜在类型是不同的
	//reString2Slice = ([]ReString2)(reStringSlice)
	var stringSlice []string
	t.Log(stringSlice)
	// 编译错误
	//reStringSlice = ([]ReString)(stringSlice)
}

func TestTypeAliasesVsReType(t *testing.T) {
	var s string = "1"
	// 类型别名实际和原类型等价可以进行强转
	s1 := AliasString(s)
	t.Log(s1)
	_, ok := interface{}(s).(string)
	assert.True(t, ok)
	reString := ReString(s)
	t.Log(reString)
	_, ok = interface{}(s).(string)
	t.Log(ok)

	var aliasString AliasString = s
	t.Log(aliasString)
	t.Log(s == aliasString)
	// 对于类型重定义操作,对于类型重定义后,重定义的类型和重定义后的类型实际是两个完全不同的类型,因此不支持任何比较以及赋值操作,但支持强制类型转换
	//var reString1 ReString = s
	//t.Log(reString1)
	//t.Log(s == reString1)
}

func TestAssertType(t *testing.T) {
	//a := 1
	a := "1"
	// 需要通过第二个返回字段判断断言是否成功
	i, ok := interface{}(a).(int)
	if ok {
		t.Log(i, "类型为int,断言成功")
	} else {
		t.Log(i, "当前类型不为int")
	}
}

func TestAssertTypeWithSwitch(t *testing.T) {
	a := 1
	switch v := interface{}(a).(type) {
	case int:
		t.Log("int", v)
	case string:
		t.Log("string", v)
	default:
		t.Log("unknown type : ", a)
	}
}

func TestBigDataToSmallData(t *testing.T) {
	src := int16(-255)
	i := int8(src) // 转换后的数据和原始数据不一致; 原因在于当转换为int8时会对-255的16位补码直接截取最后8位,因此得到的结果为1
	//i := int8(int16(-255)) // 编译会直接抛出错误
	assert.Equal(t, -255, i)
}

func TestStringConvert(t *testing.T) {
	s := string(-1)
	t.Log(s)
	// 这里实际是执行的ascii转码
	s = string(97)
	t.Log(s)
}
