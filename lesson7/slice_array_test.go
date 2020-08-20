package lesson7

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestSliceOutIndex(t *testing.T) {
	// 当切片长度和容量不同时
	s := make([]int, 4, 5)
	t.Log(s[0])
	// 当超过切片的长度时,会抛出数组越界异常, 容量表示当前切片指向数组的长度,切片的长度代表的当前可以操作的数组长度
	//t.Log(s[4])
}

func TestCompareSliceAndArray(t *testing.T) {
	// 对于数组和切片声明方式的不同点在于 当使用字面值进行声明时,数组必须指定长度,但切片由于其支持自动扩容,因此不需要显式声明容量
	a := [3]int{1, 2, 3}
	s := []int{1, 2, 3}
	t.Log(len(a), len(s), cap(a), cap(s))
	t.Logf("%T", a)
	t.Logf("%T", s)
	a1 := a
	a1[0] = 0
	// 由于数组是值类型,因此当a赋值给a1时,实际是进行了值复制,开辟了新的数组内存空间,因此当操作赋值后的数组时实际是操作的不同的内存空间;但如果数组的元素为引用类型,虽然操作的数组内存空间不同,但数组的元素空间相同
	assert.Equal(t, a[0], a1[0])
}

func TestMakeSlice(t *testing.T) {
	s := make([]int, 1)
	assert.Equal(t, len(s), cap(s))
	// 容量要求大于等于长度
	s1 := make([]int, 1, 3)
	t.Log(len(s1), cap(s1))
	// 切片的长度限制了切片可以操作数组的长度 对于 s1 长度为1,容量为3(底层数组的长度),对于s1而言,只能看到1个元素,可以操作的数组的索引范围为[0,1)
	// 这里只会迭代 len(s1)次
	for _, v := range s1 {
		t.Log(v)
	}
}

func TestMakeSliceFromAnotherSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := s[3:6]
	assert.Equal(t, 6-3, len(s1))
	assert.Equal(t, cap(s)-3, cap(s1))
	// 当原切片扩容后,s底层数组发生变化,生成了新的数组
	s = append(s, 9)
	s[3] = 0
	// 当发生扩容后,s 和 s1 底层数组就不相同了
	assert.Equal(t, s[3], s1[0])
}

// 正确预估容量
func TestPredictCap(t *testing.T) {
	// 容量在 0-1024范围内时,每次扩容后的容量是上次容量的2倍
	s := make([]int, 0)
	lastCap := cap(s)
	for i := 0; i < 513; i++ {
		s = append(s, i)
		if cap(s) != lastCap {
			if lastCap != 0 {
				assert.Equal(t, 2, cap(s)/lastCap)
			}
			lastCap = cap(s)
		}
	}
	assert.Equal(t, 1024, lastCap)
	// 当容量大于等于 1024时
	s = append(s, make([]int, 511)...)
	assert.Equal(t, len(s), cap(s))
	// 当1024容量需要扩容时,扩容基数是原容量的1.25倍
	s = append(s, 1)
	assert.Equal(t, int(1024*1.25), cap(s))
	t.Log(len(s), cap(s))
	// 当追加元素后的总数量长度超过原有容量的2倍,则切片的容量以最新的长度为基准,最终切片的容量会大于等于最新长度
	s = append(s, make([]int, cap(s)+1)...)
	assert.True(t, cap(s) >= len(s))
}

func TestSliceAddress(t *testing.T) {
	s := make([]int, 1)
	t.Log(unsafe.Pointer(&s), unsafe.Pointer(&(s[0])), len(s), cap(s))
	t.Logf("%p\n", s) //底层数组地址
	s = append(s, 1)
	t.Log(unsafe.Pointer(&s), unsafe.Pointer(&(s[0])), len(s), cap(s))
	t.Logf("%p\n", s)
	s = append(s, 1)
	t.Log(unsafe.Pointer(&s), unsafe.Pointer(&(s[0])), len(s), cap(s))
	t.Logf("%p\n", s)
	s = append(s, 1)
	t.Log(unsafe.Pointer(&s), unsafe.Pointer(&(s[0])), len(s), cap(s))
	t.Logf("%p\n", s)
	s1 := append(s, 1)
	t.Log(unsafe.Pointer(&s1))
	s2 := append(s, 1)
	t.Log(unsafe.Pointer(&s2))
}

func TestAppend(t *testing.T) {
	a := make([]int, 1)
	aptr := &a
	t.Log(unsafe.Pointer(&aptr)) // 切片实例变量地址
	t.Log(unsafe.Pointer(aptr))  //切片实例地址
	t.Log(&a)
	a = append(a, 1)
	aptr1 := &a
	t.Log(unsafe.Pointer(&aptr1))
	t.Log(unsafe.Pointer(aptr1))

	b := append(a, 1)
	bptr := &b
	t.Log(unsafe.Pointer(&bptr))
	t.Log(unsafe.Pointer(bptr))

	//assert.Equal(t,2,len(a))

	t.Logf("%T", a)
}

func TestSliceGrowth(t *testing.T) {
	s := make([]int, 1)
	t.Log(unsafe.Pointer(&s))
	s = append(s, 1, 2)
	t.Log(unsafe.Pointer(&s))
}

func TestAppendNewSlice(t *testing.T) {
	a := make([]int, 1, 5)
	t.Logf("%p", &a)
	// 新的切片实例赋值给已存在的切片实例时,实际是将slice结构体中的数组指针/长度/容量字段赋值给已存在的实例
	a = append(a, 1)
	t.Logf("%p", &a)
	// 生成新的slice实例
	b := append(a, 1)
	t.Logf("%p", &b)
	// 当切片进行赋值操作时,实际是将b实例中的三个属性复制给a中的属性,而不是将变量a指向的内存地址改变
	a = b
	t.Logf("%p", &a)
	t.Log(len(a), len(b), cap(a), cap(b))
}
