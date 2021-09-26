// Package localcache use memory to read/write data.
package localcache

import (
	"time"
)

var expiredMilliSecond int = 30 * 1000

func makeMilliSecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (c *cache) deleteExpiredData(key string) {
	delete(c.data, key)
}

// Get retrive data with key.
func (c *cache) Get(key string) (data interface{}) {
	cd := c.data[key]
	if cd.expired < makeMilliSecond() {
		c.deleteExpiredData(key)
		return nil
	}
	return c.data[key].stored
}

// Set save data with key, data is stored for 30 seconds.
func (c *cache) Set(key string, data interface{}) {
	cd := c.data[key]
	cd.stored = data
	cd.expired = makeMilliSecond() + int64(expiredMilliSecond)
	c.data[key] = cd
}

// New create a localcache.
func New() (c Cache) {
	c = &cache{
		data: make(map[string]cacheData),
	}
	return c
}
