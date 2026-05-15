package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	entries  map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	newEntry := cacheEntry{createdAt: time.Now(), val: val}
	cache.entries[key] = newEntry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry, ok := cache.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()
	for t := range ticker.C {
		expiryTime := t.Add(-cache.interval)
		cache.mu.Lock()
		for key, entry := range cache.entries {
			if entry.createdAt.Before(expiryTime) {
				delete(cache.entries, key)
			}
		}
		cache.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.reapLoop()
	return cache
}
