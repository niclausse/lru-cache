package lru

import (
	"container/list"
	"errors"
	"sync"
)

var (
	ErrNotExist = errors.New("not exists")
)

type Cache struct {
	cap      int        // capacity of lru cache
	values   *list.List // use doubly linked list for data storage
	cacheMap map[interface{}]*list.Element
	lock     sync.Mutex
}

func NewLRUCache(size int) *Cache {
	return &Cache{
		cap:      size,
		values:   list.New(),
		cacheMap: make(map[interface{}]*list.Element, size),
	}
}

func (c *Cache) Get(key interface{}) (interface{}, error) {
	e, ok := c.cacheMap[key]
	if ok {
		return e.Value, nil
	}

	return nil, ErrNotExist
}

func (c *Cache) Put(key, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	// 链表实际长度已达预设容量
	if c.values.Len() == c.cap {
		tail := c.values.Back()
		c.values.Remove(tail)
		delete(c.cacheMap, key)
	}

	c.cacheMap[key] = c.values.PushFront(value)
}
