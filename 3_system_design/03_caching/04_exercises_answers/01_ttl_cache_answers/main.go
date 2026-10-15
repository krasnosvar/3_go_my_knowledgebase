// Упражнение 01: кэш с TTL-протуханием (Решение).
package main

import (
	"fmt"
	"time"
)

var _ = time.Second

type item struct {
	value   string
	expires int64
}

type ttlCache struct {
	m   map[string]item
	now func() int64
}

func newTTLCache(now func() int64) *ttlCache {
	return &ttlCache{m: make(map[string]item), now: now}
}

func (c *ttlCache) Set(key, val string, ttl int64) {
	c.m[key] = item{value: val, expires: c.now() + ttl}
}

func (c *ttlCache) Get(key string) (string, bool) {
	it, ok := c.m[key]
	if !ok {
		return "", false
	}
	if c.now() >= it.expires {
		delete(c.m, key) // ленивое удаление протухшего
		return "", false
	}
	return it.value, true
}

func main() {
	fmt.Println("--- Проверка 01_ttl_cache (Решение) ---")

	var clock int64
	c := newTTLCache(func() int64 { return clock })

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	clock = 0
	c.Set("a", "A", 10)

	if v, ok := c.Get("a"); !ok || v != "A" {
		check(false, "сразу после Set ключ 'a' должен читаться как \"A\"")
	}

	clock = 9
	if _, ok := c.Get("a"); !ok {
		check(false, "в момент 9 ключ 'a' ещё не протух")
	}

	clock = 10
	if _, ok := c.Get("a"); ok {
		check(false, "в момент 10 ключ 'a' должен считаться протухшим")
	}

	if _, ok := c.Get("missing"); ok {
		check(false, "несуществующий ключ не должен находиться")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
