// Упражнение 08: map под мьютексом + метрики на атомиках (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Cache struct {
	mu     sync.Mutex
	m      map[string]string
	hits   atomic.Int64
	misses atomic.Int64
}

func newCache() *Cache { return &Cache{m: map[string]string{}} }

func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[k] = v
}

func (c *Cache) Get(k string) (string, bool) {
	c.mu.Lock()
	v, ok := c.m[k]
	c.mu.Unlock()
	if ok {
		c.hits.Add(1)
	} else {
		c.misses.Add(1)
	}
	return v, ok
}

func (c *Cache) Stats() (int64, int64) {
	return c.hits.Load(), c.misses.Load()
}

func main() {
	fmt.Println("--- Проверка 08_hybrid (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	c := newCache()
	c.Set("a", "A")
	c.Set("b", "B")

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Get("a")
		}()
	}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Get("zzz")
		}()
	}
	wg.Wait()

	hits, misses := c.Stats()
	check(hits == 100, fmt.Sprintf("hits = %d, ожидалось 100", hits))
	check(misses == 50, fmt.Sprintf("misses = %d, ожидалось 50", misses))
	v, ok := c.Get("b")
	check(ok && v == "B", "Get(b) должен вернуть \"B\"")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
