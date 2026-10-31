package pokecache

import (
	"sync"
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newCacheEntry := cacheEntry{createdAt: time.Now(), val: val}
	c.cacheEntries[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	test, ok := c.cacheEntries[key]

	return test.val, ok

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	//defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()
		for key, ce := range c.cacheEntries {

			if time.Since(ce.createdAt) >= c.interval {

				delete(c.cacheEntries, key)

			}
		}
		c.mu.Unlock()
	}

}

func NewCache(iv time.Duration) *Cache {
	var theCache Cache
	ec := make(map[string]cacheEntry)
	theCache.interval = iv
	theCache.cacheEntries = ec
	go theCache.reapLoop()
	return &theCache
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	interval     time.Duration
	mu           sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
