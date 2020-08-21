package lesson8

import (
	"container/ring"
	"testing"
)

func TestRing(t *testing.T) {
	var r ring.Ring
	t.Log(r.Len())
	r1 := ring.New(10)
	t.Log(r1.Len())
}
