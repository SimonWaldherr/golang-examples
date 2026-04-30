// Description: sync.Mutex - protect shared state from concurrent goroutines
// Tags: mutex, sync, goroutine, concurrent, race condition, lock, unlock, RWMutex
package main

import (
	"fmt"
	"sync"
)

// SafeCounter is a goroutine-safe counter backed by a mutex.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current counter value for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

// ReadHeavyCache demonstrates sync.RWMutex: multiple readers or one writer.
type ReadHeavyCache struct {
	mu    sync.RWMutex
	cache map[string]string
}

func (rh *ReadHeavyCache) Set(key, val string) {
	rh.mu.Lock() // exclusive write lock
	defer rh.mu.Unlock()
	rh.cache[key] = val
}

func (rh *ReadHeavyCache) Get(key string) (string, bool) {
	rh.mu.RLock() // shared read lock
	defer rh.mu.RUnlock()
	v, ok := rh.cache[key]
	return v, ok
}

func main() {
	// --- Mutex example ---
	counter := SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc("hits")
		}()
	}
	wg.Wait()
	fmt.Println("hits:", counter.Value("hits")) // always 1000

	// --- RWMutex example ---
	cache := ReadHeavyCache{cache: make(map[string]string)}
	cache.Set("lang", "Go")
	cache.Set("version", "1.22")

	// Spawn many concurrent readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if v, ok := cache.Get("lang"); ok {
				fmt.Printf("reader %d: lang=%s\n", id, v)
			}
		}(i)
	}
	wg.Wait()
}
