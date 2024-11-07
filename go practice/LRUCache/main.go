package main

import (
	"container/list" // For doubly linked list implementation
	"fmt"
	"sync" // For mutex to make cache thread-safe
)

// Pair represents a key-value pair stored in the cache
// K must be comparable (can be used as map key)
// V can be any type (using 'any' constraint)
type Pair[K comparable, V any] struct {
	key   K
	value V
}

// LRUCache implements a thread-safe Least Recently Used cache
// Generic types K (key) and V (value) allow for flexible usage
type LRUCache[K comparable, V any] struct {
	cache    map[K]*list.Element // Maps keys to doubly linked list nodes
	list     *list.List         // Doubly linked list to maintain access order
	mutex    sync.Mutex         // Ensures thread-safety for cache operations
	capacity int                // Maximum number of items cache can hold
}

// NewLRUCache creates and initializes a new LRU cache with specified capacity
// Returns a pointer to the new cache instance
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		list:     list.New(),                    // Initialize empty doubly linked list
		cache:    make(map[K]*list.Element),     // Initialize empty map
	}
}

// Get retrieves a value from the cache by its key
// Returns the value and true if found, zero value and false if not found
// Also moves accessed item to front of list (marks as most recently used)
func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock() // Ensures mutex is unlocked even if panic occurs

	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)                    // Mark as most recently used
		return elem.Value.(Pair[K, V]).value, true  // Type assert and return value
	}
	var zeroValue V // Return zero value if key not found
	return zeroValue, false
}

// Put adds or updates a key-value pair in the cache
// If key exists: updates value and moves to front
// If key doesn't exist: adds new entry, evicting oldest if at capacity
func (c *LRUCache[K, V]) Put(key K, val V) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		// Key exists: update value and move to front
		c.list.MoveToFront(elem)
		elem.Value = Pair[K, V]{key, val}
	} else {
		// Key doesn't exist: check capacity and add new entry
		if c.list.Len() >= c.capacity {
			// At capacity: remove oldest item (from back of list)
			oldest := c.list.Back()
			if oldest != nil {
				c.list.Remove(oldest)
				delete(c.cache, oldest.Value.(Pair[K, V]).key)
			}
		}
		// Add new item to front of list and map
		elem := c.list.PushFront(Pair[K, V]{key, val})
		c.cache[key] = elem
	}
}

// Remove deletes an item from the cache by its key
// If key exists, removes from both map and list
func (c *LRUCache[K, V]) Remove(key K) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		delete(c.cache, key)    // Remove from map
		c.list.Remove(elem)     // Remove from list
	}
}

func main() {
	// Create new cache with capacity 3
	cache := NewLRUCache[string, int](3)

	// Add three items
	cache.Put("one", 1)
	cache.Put("two", 2)
	cache.Put("three", 3)

	// Access "two", moving it to front of list
	if val, ok := cache.Get("two"); ok {
		fmt.Printf("Key 'two' is in cache with value: %d\n", val)
	}

	// Add "four", which will evict "one" (least recently used)
	cache.Put("four", 4) // This will evict "one" from the cache

	// Verify "one" was evicted
	if _, ok := cache.Get("one"); !ok {
		fmt.Println("Key 'one' is no longer in the cache")
	}

	// Remove "three" explicitly
	cache.Remove("three")

	// Verify "three" was removed
	if _, ok := cache.Get("three"); !ok {
		fmt.Println("Key 'three' has been removed from the cache")
	}
}