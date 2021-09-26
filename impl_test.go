package localcache

import (
	"reflect"
	"testing"
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
