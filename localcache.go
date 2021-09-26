package localcache

type Cache interface {
	Get(key string) (data interface{})
	Set(key string, data interface{})
}

type cache struct {
	data map[string]interface{}
}
