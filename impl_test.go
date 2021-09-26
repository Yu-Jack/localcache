package localcache

import "testing"

func TestLocalcache(t *testing.T) {
	cache := New()
	cache.Get()
	cache.Set()
}
