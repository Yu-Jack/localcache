package localcache

type Cache interface {
	Get()
	Set()
}

type cache struct{}
