package lesson13

// 指针类型方法和值类型方法对比

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Cat struct {
	name           string // 名字。
	scientificName string // 学名。
	category       string // 动物学基本分类。
}

func New(name, scientificName, category string) Cat {
	return Cat{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.category, cat.name)
}

func TestMethod(t *testing.T) {
	cat := New("little pig", "American Shorthair", "cat")
	cat.SetName("monster") // (&cat).SetName("monster")
	t.Logf("The cat: %s\n", cat)

	cat.SetNameOfCopy("little pig")
	t.Logf("The cat: %s\n", cat)

	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}

	_, ok := interface{}(cat).(Pet)
	t.Logf("Cat implements interface Pet: %v\n", ok)
	assert.False(t, ok)
	_, ok = interface{}(&cat).(Pet)
	t.Logf("*Cat implements interface Pet: %v\n", ok)
	assert.True(t, ok) // 原因在于SetName方法实际只能通过指针类型来调用,对于接口实现而言,通过指针类型可以调用所有接口的方法,因此指针类型才是接口的实现类型
}
