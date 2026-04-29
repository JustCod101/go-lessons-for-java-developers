package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := NewCache(time.Second)
	defer c.Close()

	c.Set("foo", "bar", time.Millisecond*100)

	val, found := c.Get("foo")
	if !found || val.(string) != "bar" {
		t.Fatal("expected to find foo=bar")
	}

	time.Sleep(time.Millisecond * 200)
	_, found = c.Get("foo")
	if found {
		t.Fatal("expected foo to be expired")
	}
}

func BenchmarkCacheGet(b *testing.B) {
	c := NewCache(0)
	c.Set("key", "value", 0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Get("key")
	}
}

func BenchmarkCacheSet(b *testing.B) {
	c := NewCache(0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Set("key", "value", 0)
	}
}

func BenchmarkParallelGet(b *testing.B) {
	c := NewCache(0)
	c.Set("key", "value", 0)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get("key")
		}
	})
}
