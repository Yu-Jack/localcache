// Package localcache use memory to read/write data.
package localcache

import (
	"time"
)

var expiredMilliSecond int = 30 * 1000

func currentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (c *cache) deleteExpired(key string, expired int64) {
	if expired < currentMillis() {
		delete(c.data, key)
	}
}

// Get retrive data with key.
func (c *cache) Get(key string) (data interface{}) {
	cd, ok := c.data[key]
	if !ok {
		return nil
	}
	c.deleteExpired(key, cd.expired)
	return c.data[key].stored
}

// Set save data with key, data is stored for 30 seconds.
func (c *cache) Set(key string, data interface{}) {
	cd := c.data[key]
	cd.stored = data
	cd.expired = currentMillis() + int64(expiredMilliSecond)
	c.data[key] = cd
}

// New create a localcache.
func New() (c Cache) {
	c = &cache{
		data: make(map[string]cacheData),
	}
	return c
}
