package localcache

func (c *cache) Get(key string) (data interface{}) {
	return c.data[key]
}

func (c *cache) Set(key string, data interface{}) {
	c.data[key] = data
}

func New() (c Cache) {
	c = &cache{
		data: make(map[string]interface{}),
	}
	return c
}
