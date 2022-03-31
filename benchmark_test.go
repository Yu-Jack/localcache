package localcache

import (
	"fmt"
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
