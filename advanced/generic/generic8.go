package main

import (
	"fmt"
)

type Cache[K comparable, V any] struct {
	store map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{store: make(map[K]V)}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.store[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	value, exists := c.store[key]
	return value, exists
}

func main() {
	cache := NewCache[string, int]()
	cache.Set("one", 1)
	cache.Set("two", 2)

	if value, found := cache.Get("one"); found {
		fmt.Println("Value for 'one':", value)
	} else {
		fmt.Println("'one' not found in cache")
	}

	if value, found := cache.Get("three"); found {
		fmt.Println("Value for 'three':", value)
	} else {
		fmt.Println("'three' not found in cache")
	}
}
