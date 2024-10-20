package main

import (
	"container/list"
	"fmt"
	"sync"
)

type Pair[K comparable, V any] struct {
	key   K
	value V
}

type LRUCache[K comparable, V any] struct {
	cache    map[K]*list.Element
	list     *list.List
	mutex    sync.Mutex
	capacity int
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[K]*list.Element),
	}
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		return elem.Value.(Pair[K, V]).value, true
	}
	var zeroValue V
	return zeroValue, false
}

func (c *LRUCache[K, V]) Put(key K, val V) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		elem.Value = Pair[K, V]{key, val}
	} else {
		if c.capacity <= c.list.Len() {
			oldest := c.list.Back()
			if oldest != nil {
				c.list.Remove(oldest)
				delete(c.cache, oldest.Value.(Pair[K, V]).key)
			}
		}
	}
}

func (c *LRUCache[K, V]) Remove(key K) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if elem, ok := c.cache[key]; ok {
		delete(c.cache, key)
		c.list.Remove(elem)
	}
}

func main() {
	cache := NewLRUCache[string, int](3)
	cache.Put("one", 1)
	cache.Put("two", 2)
	cache.Put("three", 3)
	if val, ok := cache.Get("two"); ok {
		fmt.Printf("Key 'two' is in cache with value: %d\n", val)
	}
	cache.Put("four", 4) // This will evict "one" from the cache
	if _, ok := cache.Get("one"); !ok {
		fmt.Println("Key 'one' is no longer in the cache")
	}
	cache.Remove("three")
	if _, ok := cache.Get("three"); !ok {
		fmt.Println("Key 'three' has been removed from the cache")
	}
}
