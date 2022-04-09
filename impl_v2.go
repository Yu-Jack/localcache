// Package localcache use memory to read/write data.
package localcache

import (
	"reflect"
	"time"
)

func (c *cacheV2) deleteKey(key string) {
	delete(c.data, key)
}

func (c *cacheV2) deleteTimer(key string) {
	var deletedIndex int
	for i, timer := range c.timerList {
		if timer.key == key {
			deletedIndex = i
			break
		}
	}
	c.timerList = append(c.timerList[:deletedIndex], c.timerList[deletedIndex+1:]...)
}

func (c *cacheV2) DeleteAll() {
	c.locker.Lock()
	defer c.locker.Unlock()

	keys := reflect.ValueOf(c.data).MapKeys()
	for _, key := range keys {
		c.deleteKey(key.String())
	}
}

// Get retrive data with key.
func (c *cacheV2) Get(key string) (data interface{}) {
	c.locker.Lock()
	defer c.locker.Unlock()
	cd, ok := c.data[key]
	if !ok {
		return nil
	}
	return cd.stored
}

// Set save data with key, data is stored for 30 seconds.
func (c *cacheV2) Set(key string, data interface{}) {
	c.locker.Lock()
	defer c.locker.Unlock()

	cd, ok := c.data[key]
	if !ok {
		cd.timer = time.NewTimer(expiredMilliSecond)
		c.timerList = append(c.timerList, cacheTimer{
			timer: cd.timer,
			key:   key,
		})
	} else {
		cd.timer.Reset(expiredMilliSecond)
	}
	cd.stored = data
	cd.expired = currentMillis() + int64(expiredMilliSecond)
	c.data[key] = cd
}

func (c *cacheV2) listenExpiredTimer() {
	for {
		for _, t := range c.timerList {
			select {
			case <-t.timer.C:
				c.locker.Lock()
				c.deleteKey(t.key)
				c.deleteTimer(t.key)
				c.locker.Unlock()
			default:
			}
		}
		time.Sleep(checkPeriod)
	}
}

// New create a localcache.
func NewCacheV2() Cache {
	c := &cacheV2{
		data: make(map[string]cacheData),
	}
	go c.listenExpiredTimer()
	return c
}
