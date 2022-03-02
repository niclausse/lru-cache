package lru

import (
	"container/list"
	"sync"
)

type Cache struct {
	cap      int        // capacity of lru cache
	values   *list.List // use doubly linked list for data storage
	cacheMap map[interface{}]*list.Element
	lock     sync.Mutex
}

type entry struct {
	key, value interface{}
}

func NewLRUCache(size int) *Cache {
	return &Cache{
		cap:      size,
		values:   list.New(),
		cacheMap: make(map[interface{}]*list.Element, size),
	}
}

func (c *Cache) Get(key interface{}) interface{} {
	e, ok := c.cacheMap[key]
	if !ok {
		return nil
	}

	c.values.MoveToFront(e)
	return e.Value.(*entry).value
}

func (c *Cache) Put(key, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	e, ok := c.cacheMap[key]
	if ok {
		e.Value.(*entry).value = value
		c.values.MoveToFront(e)
		return
	}

	// 链表实际长度已达预设容量
	if c.values.Len() == c.cap {
		tail := c.values.Back()
		c.values.Remove(tail)
		delete(c.cacheMap, tail.Value.(*entry).key)
	}

	c.cacheMap[key] = c.values.PushFront(&entry{key: key, value: value})
}
