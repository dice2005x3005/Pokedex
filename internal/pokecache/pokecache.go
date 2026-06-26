package pokecache

import (
	"time"
	"sync"
)


type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(t time.Duration) *Cache {
	c := &Cache {
		entries: map[string]cacheEntry{},
		mu: sync.Mutex{},
	}
	go c.reapLoop(t)
	return c
}

func (c *Cache) Add(a string, v []byte) {
	c.mu.Lock()
	c.entries[a] = cacheEntry{createdAt: time.Now(), val: v}
	c.mu.Unlock()
}

func (c *Cache) Get(k string) ([]byte, bool) {
	c.mu.Lock()
	v, ok := c.entries[k]
	c.mu.Unlock()
	return v.val, ok
}

func (c *Cache) reapLoop(t time.Duration) {
	ticker := time.NewTicker(t)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.entries {
		if time.Since(v.createdAt) > t {
			delete(c.entries, k)
		}
		c.mu.Unlock()
	}
	}
}