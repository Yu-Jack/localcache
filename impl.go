// Package localcache use memory to read/write data.
package localcache

import (
	"sync"
	"time"
)

var expiredMilliSecond time.Duration = 30 * time.Second
var checkPeriod time.Duration = 1 * time.Second

func currentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (c *cache) delete(key string) {
	delete(c.data, key)
}

func (c *cache) deleteTimer(key string) {
	var deletedIndex int
	for i, timer := range c.timerList {
		if timer.key == key {
			deletedIndex = i
			break
		}
	}
	c.timerList = append(c.timerList[:deletedIndex], c.timerList[deletedIndex+1:]...)
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
	cd.locker = new(sync.Mutex)
	cd.locker.Lock()
	defer cd.locker.Unlock()

	cd.stored = data
	cd.expired = currentMillis() + int64(expiredMilliSecond)
	if !ok {
		cd.timer = time.NewTimer(expiredMilliSecond)
		c.timerList = append(c.timerList, cacheTimer{
			timer: cd.timer,
			key:   key,
		})
	} else {
		cd.timer.Reset(expiredMilliSecond)
	}
	c.data[key] = cd
}

func (c *cache) listenExpiredTimer() {
	for {
		for _, t := range c.timerList {
			select {
			case <-t.timer.C:
				c.delete(t.key)
				c.deleteTimer(t.key)
			default:
			}
		}
		time.Sleep(checkPeriod)
	}
}

// New create a localcache.
func New() Cache {
	c := &cache{
		data: make(map[string]cacheData),
	}
	go c.listenExpiredTimer()
	return c
}
