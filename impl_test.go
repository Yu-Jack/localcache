package localcache

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite

	cache Cache
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

func (suit *ExampleTestSuite) SetupTest() {
	fmt.Println("prepare cache ...")
	suit.cache = NewCache()
}

func (suit *ExampleTestSuite) TearDownTest() {
	fmt.Println("delete all cache ...")
	suit.cache.DeleteAll()
}

func (suit *ExampleTestSuite) TestLocalcache() {
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
		suit.cache.Set(tc.key, tc.data)

		got := suit.cache.Get(tc.key)

		assert.Equal(suit.T(), tc.expect, got)
	}
}

func (suit *ExampleTestSuite) TestLocalcache_overwriteData() {
	expect := 2
	key := "key1"
	suit.cache.Set(key, 1)
	suit.cache.Set(key, 2)
	got := suit.cache.Get(key)

	assert.Equal(suit.T(), expect, got)
}

func (suit *ExampleTestSuite) TestLocalcache_notFoundData() {
	expect := error(nil)
	key := "key1"
	notFoundKey := "notFoundkey"
	suit.cache.Set(key, 1)

	got := suit.cache.Get(notFoundKey)

	assert.Equal(suit.T(), expect, got)
}

func (suit *ExampleTestSuite) TestLocalcache_expiredData() {
	expiredMilliSecond = 1 * time.Second
	key := "key1"
	expect := error(nil)
	suit.cache.Set(key, 1)
	suit.cache.Set(key, 2)
	time.Sleep(3 * time.Second)

	got := suit.cache.Get(key)

	assert.Equal(suit.T(), expect, got)
}

func (suit *ExampleTestSuite) TestLocalcache_concurrent() {
	expiredMilliSecond = 1 * time.Second
	expect := error(nil)
	key := "key1"
	go func() {
		suit.cache.Set(key, 1)
	}()
	go func() {
		suit.cache.Set(key, 2)
	}()
	time.Sleep(3 * time.Second)

	got := suit.cache.Get(key)

	assert.Equal(suit.T(), expect, got)
}
