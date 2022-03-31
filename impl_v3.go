// Package localcache use memory to read/write data.
package localcache

import (
	"reflect"
	"time"
)

func (c *cacheV3) deleteKey(key string) {
	c.locker.Lock()
	defer c.locker.Unlock()

	delete(c.data, key)
}

func (c *cacheV3) DeleteAll() {
	keys := reflect.ValueOf(c.data).MapKeys()
	for _, key := range keys {
		c.deleteKey(key.String())
	}
}

// Get retrive data with key.
func (c *cacheV3) Get(key string) (data interface{}) {
	cd, ok := c.data[key]
	if !ok {
		return nil
	}
	return cd.stored
}

// Set save data with key, data is stored for 30 seconds.
func (c *cacheV3) Set(key string, data interface{}) {
	c.locker.Lock()
	defer c.locker.Unlock()

	cd, ok := c.data[key]
	if !ok {
		cd.timer = time.AfterFunc(expiredMilliSecond, func() {
			c.deleteKey(key)
		})
	} else {
		cd.timer.Reset(expiredMilliSecond)
	}
	cd.stored = data
	cd.expired = currentMillis() + int64(expiredMilliSecond)
	c.data[key] = cd
}

// New create a localcache.
func NewCacheV3() Cache {
	c := &cacheV3{
		data: make(map[string]cacheData),
	}
	return c
}
