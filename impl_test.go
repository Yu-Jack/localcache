package localcache

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestLocalcache(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		data   interface{}
		expect interface{}
	}{
		{name: "integer", key: "key1", data: 1, expect: 1},
		{name: "boolean", key: "key1", data: true, expect: true},
		{name: "float", key: "key1", data: 3.1415926, expect: 3.1415926},
		{name: "byte", key: "key1", data: []byte("cool"), expect: []byte("cool")},
		{name: "string", key: "key1", data: "string", expect: "string"},
		{name: "array", key: "key1", data: [3]int{1, 2, 3}, expect: [3]int{1, 2, 3}},
		{name: "slice", key: "key1", data: []int{1, 2, 3}, expect: []int{1, 2, 3}},
		{name: "map", key: "key1", data: map[string]string{"name": "Jack"}, expect: map[string]string{"name": "Jack"}},
		{name: "nested map", key: "key1", data: map[string]map[string]string{"name": {"man": "jack"}}, expect: map[string]map[string]string{"name": {"man": "jack"}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cache := New()
			cache.Set(tc.key, tc.data)

			got := cache.Get(tc.key)

			// use to format the diff instead of reflect.DeepEqual
			diff := cmp.Diff(tc.expect, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestLocalcache_not_found_data(t *testing.T) {
	expect := error(nil)
	key := "key1"
	notFoundKey := "notFoundkey"
	cache := New()
	cache.Set(key, 1)

	got := cache.Get(notFoundKey)

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
