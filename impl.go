// Package localcache use memory to read/write data.
package localcache

import (
	"time"
)

var expiredMilliSecond int = 30 * 1000

func makeMilliSecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

var c = &cache{
	data: make(map[string]cacheData),
}

func (c *cache) get(key string) (data interface{}) {
	cd, ok := c.data[key]
	if !ok {
		return nil
	}
	if cd.expired < makeMilliSecond() {
		delete(c.data, key)
		return nil
	}
	return c.data[key].stored
}

func (c *cache) set(key string, data interface{}) {
	cd := c.data[key]
	cd.stored = data
	cd.expired = makeMilliSecond() + int64(expiredMilliSecond)
	c.data[key] = cd
}

// Get retrive data with key.
func Get(key string) (data interface{}) {
	return c.get(key)
}

// Set save data with key, data is stored for 30 seconds.
func Set(key string, data interface{}) {
	c.set(key, data)
}
