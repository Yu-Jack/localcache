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

// One lock for one cache data.
type cacheV1 struct {
	data      map[string]cacheData
	timerList []cacheTimer

	locker      map[string]*sync.Mutex
	cacheLocker sync.Mutex
}

// Only one lock for whole cache data.
type cacheV2 struct {
	data      map[string]cacheData
	timerList []cacheTimer

	locker sync.Mutex
}

// Only one lock for whole cache data.
// User time.AfterFunc to trigger deletion.
type cacheV3 struct {
	data map[string]cacheData

	locker sync.Mutex
}
