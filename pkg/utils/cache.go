package utils

import (
	"sync"
	"time"
)

// Represents an item in the cache
type cacheItem struct {
	cachedItem		interface{}
	timeout			time.Time``
}


// Interface defining what capabilities a cache should have
type Cache interface {

	GetItem(key string) (interface{}, bool)
	AddItem(key string, item interface{}, expirySec time.Duration)
	BasicAddItem(key string, item interface{})
}


// Basic implementation of a cache that uses a sync.Map
type InMemoryCache struct {
	underlying			sync.Map
}

// Retrieves an item from the cache if it is available (and not expired)
func (c InMemoryCache) GetItem(key string) (item interface{}, found bool) {
	cItem, found := c.underlying.Load(key)
	if !found {
		return nil, false
	} else {
		if cItem.(cacheItem).timeout.Before(time.Now()) {
			c.underlying.Delete(key)
			return nil, false
		} else {
			return cItem.(cacheItem).cachedItem, true
		}
	}
}

// Adds an item to the cache
func (c InMemoryCache) AddItem(key string, item interface{}, expirySec time.Duration) {
	c.underlying.Store(key, cacheItem{cachedItem: item, timeout: time.Now().Add(time.Second * expirySec)})
}

// Adds an item to the cache without an expiry
func (c InMemoryCache) BasicAddItem(key string, item interface{}) {
	c.underlying.Store(key, cacheItem{cachedItem: item, timeout: time.Now().AddDate(100, 0, 0)})
}