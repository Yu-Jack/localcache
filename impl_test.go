package localcache

import (
	"reflect"
	"testing"
	"time"
)

func TestLocalcache_integer(t *testing.T) {
	expect := 1
	key := "key1"
	cache := New()
	cache.Set(key, expect)

	got := cache.Get(key)

	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("expected: %v, got: %v", expect, got)
	}
}

func TestLocalcache_string(t *testing.T) {
	expect := "string_data"
	key := "key1"
	cache := New()
	cache.Set(key, expect)

	got := cache.Get(key)

	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("expected: %v, got: %v", expect, got)
	}
}

func TestLocalcache_map(t *testing.T) {
	expect := map[string]string{
		"name": "Jack",
	}
	key := "key1"
	cache := New()
	cache.Set(key, expect)

	got := cache.Get(key)

	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("expected: %v, got: %v", expect, got)
	}
}

func TestLocalcache_expired_data(t *testing.T) {
	expiredMilliSecond = 1 * 1000
	key := "key1"
	cache := New()
	expect := error(nil)
	cache.Set(key, 1)
	time.Sleep(2 * time.Second)

	got := cache.Get(key)

	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("expected: %v, got: %v", expect, got)
	}
}
