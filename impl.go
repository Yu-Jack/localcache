// Package localcache use memory to read/write data.
package localcache

import (
	"time"
)

var expiredMilliSecond time.Duration = 30 * time.Second

func currentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (c *cache) delete(key string) {
	delete(c.data, key)
}

// Get retrive data with key.
func (c *cache) Get(key string) (data interface{}) {
	cd, ok := c.data[key]
	if !ok {
		return nil
	}
	return cd.stored
}

// Set save data with key, data is stored for 30 seconds.
func (c *cache) Set(key string, data interface{}) {
	cd, ok := c.data[key]
	cd.stored = data
	cd.expired = currentMillis() + int64(expiredMilliSecond)
	if !ok {
		cd.timer = time.NewTimer(expiredMilliSecond)
	} else {
		cd.timer.Reset(expiredMilliSecond)
	}
	c.data[key] = cd

	go func() {
		<-cd.timer.C
		c.delete(key)
	}()
}

// New create a localcache.
func New() Cache {
	c := &cache{
		data: make(map[string]cacheData),
	}
	return c
}
