package main

import (
	"fmt"
	"sync"
)

var (
	data = make(map[string]string)
)

type SafeMap struct {
	mu sync.Mutex
	data map[string]string
}

func (m *SafeMap) Set(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *SafeMap) Get(key string) (string, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value, ok := m.data[key]
	return value, ok
}

func main() {
	var wg sync.WaitGroup
	//race condition
	safeMap := SafeMap{data: make(map[string]string)}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			safeMap.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
		}()
	}
	// wg.Wait()
	fmt.Println(data)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			value, ok := safeMap.Get(key)
			if !ok {
				fmt.Println("key not found:", key)
			} else {
				fmt.Printf("%s = %s\n", key, value)
			}
		}()
	}
	// wg.Wait()
	// wg.Wait()
}
