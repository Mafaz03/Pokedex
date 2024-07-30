package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	CreatedAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapEvery(interval)
	return c
}

func (c *Cache) Add(key string, val []byte){
	c.cache[key] = cacheEntry{
		val: val, 
		CreatedAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) reapEvery(interval time.Duration) {
	ticker := time.Tick(interval)
	for range ticker {
		c.reap(interval)
	}
}


func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range(c.cache){
		if v.CreatedAt.Before(timeAgo){
			delete(c.cache, k)
		}
	}
}