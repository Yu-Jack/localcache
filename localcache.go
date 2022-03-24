package localcache

import (
	"sync"
	"time"
)

type Cache interface {
	Get(key string) (data interface{})
	Set(key string, data interface{})
	DeleteAll()
}

type cacheData struct {
	stored  interface{}
	expired int64
	timer   *time.Timer
}

type cacheTimer struct {
	timer *time.Timer
	key   string
}

type cache struct {
	data      map[string]cacheData
	timerList []cacheTimer

	locker      map[string]*sync.Mutex
	cacheLocker sync.Mutex
}
