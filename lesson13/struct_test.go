package lesson13

import (
	"fmt"
	"testing"
)

// AnimalCategory 代表动物分类学中的基本分类法。
type AnimalCategory struct {
	kingdom string // 界。
	phylum  string // 门。
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

func TestS(t *testing.T) {
	a := AnimalCategory{
		kingdom: "测试",
	}
	t.Logf("%s", a) // 对于 %s 会查找当前类型自定义的String方法
}

type Animal struct {
	scientificName string // 学名。
	AnimalCategory        // 动物基本分类。 嵌入字段
}

func (a Animal) String() string {
	return fmt.Sprintf("%s%s", a.scientificName, a.AnimalCategory.String())
}

func TestInternalFunc(t *testing.T) {
	a := Animal{
		AnimalCategory: AnimalCategory{
			kingdom: "测试",
		},
	}
	t.Logf("%s", a) // 当Animal不存在string方法时,实际调用了嵌入字段的String方法; 当Animal存在String方法时实际调用的是当前结构体的String方法,这里实际是屏蔽了嵌入字段的重名方法
}
