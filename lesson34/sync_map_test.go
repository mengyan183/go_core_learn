package lesson34

import (
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	m := sync.Map{}

	m.Store(1,2) // 如果key在read中存在,且value不为逻辑删除状态,执行更新操作时,性能会很高效

	m.Load(1) // 如果数据在read中存在且不为逻辑删除状态时,查询的性能也会很高效
	m.Delete(1)
}
