// Package localcache use memory to read/write data.
package localcache

import (
	"reflect"
	"sync"
	"time"
)

var (
	expiredMilliSecond time.Duration = 30 * time.Second
	checkPeriod        time.Duration = 100 * time.Millisecond
)

func currentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (c *cacheV1) deleteKey(key string) {
	delete(c.data, key)
}

func (c *cacheV1) deleteTimer(key string) {
	var deletedIndex int
	for i, timer := range c.timerList {
		if timer.key == key {
			deletedIndex = i
			break
		}
	}
	c.timerList = append(c.timerList[:deletedIndex], c.timerList[deletedIndex+1:]...)
}

func (c *cacheV1) DeleteAll() {
	c.cacheLocker.Lock()
	defer c.cacheLocker.Unlock()

	keys := reflect.ValueOf(c.data).MapKeys()
	for _, key := range keys {
		c.deleteKey(key.String())
	}
}

// Get retrive data with key.
func (c *cacheV1) Get(key string) (data interface{}) {
	locker := c.lock(key)
	defer locker.Unlock()
	cd, ok := c.data[key]
	if !ok {
		return nil
	}
	return cd.stored
}

// Set save data with key, data is stored for 30 seconds.
func (c *cacheV1) Set(key string, data interface{}) {
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

func (c *cacheV1) listenExpiredTimer() {
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
func (c *cacheV1) lock(key string) *sync.Mutex {
	locker, ok := c.locker[key]
	if !ok {
		locker = new(sync.Mutex)
		c.locker[key] = locker
	}
	locker.Lock()
	return locker
}

func (c *cacheV1) deleteLockKey(key string) {
	delete(c.locker, key)
}

// New create a localcache.
func NewCacheV1() Cache {
	c := &cacheV1{
		data:   make(map[string]cacheData),
		locker: make(map[string]*sync.Mutex),
	}
	go c.listenExpiredTimer()
	return c
}
