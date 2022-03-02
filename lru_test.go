package lru

import (
	"fmt"
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	l := NewLRUCache(100)

	for i := 1; i <= 100; i++ {
		v := l.Get(i)
		if v == nil {
			v = i * 10
			l.Put(i, v)
		}
		fmt.Printf("%d: %d\n", i, v)
	}

	fmt.Printf("head: %d\n", l.values.Front().Value)
	fmt.Printf("tail: %d\n", l.values.Back().Value)
	fmt.Printf("lru len: %d\n", l.values.Len())
}
