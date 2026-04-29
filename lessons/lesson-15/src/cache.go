package cache

import (
	"sync"
	"time"
)

type Item struct {
	Value      interface{}
	Expiration int64
}

type Cache struct {
	items           map[string]Item
	mu              sync.RWMutex
	cleanupInterval time.Duration
	stopCleanup     chan struct{}
}

func NewCache(cleanupInterval time.Duration) *Cache {
	c := &Cache{
		items:           make(map[string]Item),
		cleanupInterval: cleanupInterval,
		stopCleanup:     make(chan struct{}),
	}

	if cleanupInterval > 0 {
		go c.gc()
	}
	return c
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	var expiration int64
	if ttl > 0 {
		expiration = time.Now().Add(ttl).UnixNano()
	}

	c.mu.Lock()
	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	if item.Expiration > 0 && time.Now().UnixNano() > item.Expiration {
		return nil, false
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.items, key)
	c.mu.Unlock()
}

func (c *Cache) gc() {
	ticker := time.NewTicker(c.cleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		case <-c.stopCleanup:
			return
		}
	}
}

func (c *Cache) DeleteExpired() {
	now := time.Now().UnixNano()
	c.mu.Lock()
	for k, v := range c.items {
		if v.Expiration > 0 && now > v.Expiration {
			delete(c.items, k)
		}
	}
	c.mu.Unlock()
}

func (c *Cache) Close() {
	close(c.stopCleanup)
}
