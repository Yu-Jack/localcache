package localcache

type Cache interface {
	Get(key string) (data interface{})
	Set(key string, data interface{})
}

type cacheData struct {
	stored  interface{}
	expired int64
}

type cache struct {
	data map[string]cacheData
}
