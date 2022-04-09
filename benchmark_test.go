package localcache

import (
	"fmt"
	"sync"
	"testing"
)

func Benchmark_V1_Set_And_Get_Different_Key(b *testing.B) {
	c := NewCacheV1()
	for i := 0; i < b.N; i++ {
		c.Set(fmt.Sprintf("%d", i), i)
		c.Get(fmt.Sprintf("%d", i))
	}
}

func Benchmark_V2_Set_And_Get_Different_Key(b *testing.B) {
	c := NewCacheV2()
	for i := 0; i < b.N; i++ {
		c.Set(fmt.Sprintf("%d", i), i)
		c.Get(fmt.Sprintf("%d", i))
	}
}

func Benchmark_V3_Set_And_Get_Different_Key(b *testing.B) {
	c := NewCacheV3()
	for i := 0; i < b.N; i++ {
		c.Set(fmt.Sprintf("%d", i), i)
		c.Get(fmt.Sprintf("%d", i))
	}
}

func Benchmark_V1_Set_And_Get_Same_Key(b *testing.B) {
	c := NewCacheV1()
	for i := 0; i < b.N; i++ {
		c.Set("hi", i)
		c.Get("hi")
	}
}

func Benchmark_V2_Set_And_Get_Same_Key(b *testing.B) {
	c := NewCacheV2()
	for i := 0; i < b.N; i++ {
		c.Set("hi", i)
		c.Get("hi")
	}
}

func Benchmark_V3_Set_And_Get_Same_Key(b *testing.B) {
	c := NewCacheV3()
	for i := 0; i < b.N; i++ {
		c.Set("hi", i)
		c.Get("hi")
	}
}

// There is no v1 concurrent benchmark, because v1 have concurrent map read & write fatal error.

func Benchmark_V2_Set_And_Get_Different_Key_Concurrent(b *testing.B) {
	c := NewCacheV2()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < b.N; i++ {
			c.Set(fmt.Sprintf("%d", i), i)
			c.Get(fmt.Sprintf("%d", i))
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			c.Set(fmt.Sprintf("%d", i), i)
			c.Get(fmt.Sprintf("%d", i))
		}
		wg.Done()
	}()
	wg.Wait()
}

func Benchmark_V3_Set_And_Get_Different_Key_Concurrent(b *testing.B) {
	c := NewCacheV3()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < b.N; i++ {
			c.Set(fmt.Sprintf("%d", i), i)
			c.Get(fmt.Sprintf("%d", i))
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			c.Set(fmt.Sprintf("%d", i), i)
			c.Get(fmt.Sprintf("%d", i))
		}
		wg.Done()
	}()
	wg.Wait()
}
