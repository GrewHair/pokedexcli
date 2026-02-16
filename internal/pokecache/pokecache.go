package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (cache *Cache) Add(k string, v []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.entries[k] = cacheEntry{
		createdAt: time.Now(),
		val: v,
	}
}

func (cache Cache) Get(key string) ([]byte, bool) {
	//cache.mu.Lock()
	//defer cache.mu.Unlock()
	entry, found := cache.entries[key]
	return entry.val, found
}

//func (cache *Cache) reapLoop() {
//	cache.mu.Lock()
//	defer cache.mu.Unlock()
//	for k, v := range cache.entries {
//		if (time.Now() - v.createdAt) >= 5 * time.Second {
//			delete(cache.entries, k)
//		}
//	}
//}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
	  cache.mu.Lock()
	  defer cache.mu.Unlock()
	  for k, v := range cache.entries {
	  	if time.Now().Sub(v.createdAt) >= 5 * time.Second {
	  		delete(cache.entries, k)
	  	}
	  }
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{}
	go cache.reapLoop(interval)
	return &cache
}
