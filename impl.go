// Package localcache use memory to read/write data.
package localcache

import (
	"sync"
	"time"
)

var (
	expiredMilliSecond time.Duration = 30 * time.Second
	checkPeriod        time.Duration = 1 * time.Second
)

func currentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (c *cache) deleteKey(key string) {
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
	locker := c.lock(key)
	defer locker.Unlock()

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

func (c *cache) listenExpiredTimer() {
	for {
		for _, t := range c.timerList {
			select {
			case <-t.timer.C:
				locker := c.lock(t.key)
				c.deleteLockKey(t.key)

				c.deleteKey(t.key)
				c.deleteTimer(t.key)

				locker.Unlock()
			default:
			}
		}
		time.Sleep(checkPeriod)
	}
}

// lock cache data per key, instead of whole cache store
func (c *cache) lock(key string) *sync.Mutex {
	locker, ok := c.locker[key]
	if !ok {
		locker = new(sync.Mutex)
		c.locker[key] = locker
	}
	locker.Lock()
	return locker
}

func (c *cache) deleteLockKey(key string) {
	delete(c.locker, key)
}

// New create a localcache.
func New() Cache {
	c := &cache{
		data:   make(map[string]cacheData),
		locker: make(map[string]*sync.Mutex),
	}
	go c.listenExpiredTimer()
	return c
}
