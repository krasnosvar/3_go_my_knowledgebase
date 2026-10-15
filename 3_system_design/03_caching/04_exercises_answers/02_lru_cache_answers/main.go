// Упражнение 02: LRU-кэш (Решение).
package main

import (
	"container/list"
	"fmt"
)

var _ = list.New

type entry struct {
	key   string
	value int
}

type lruCache struct {
	capacity int
	ll       *list.List
	items    map[string]*list.Element
}

func newLRU(capacity int) *lruCache {
	return &lruCache{
		capacity: capacity,
		ll:       list.New(),
		items:    make(map[string]*list.Element),
	}
}

func (c *lruCache) Get(key string) (int, bool) {
	if el, ok := c.items[key]; ok {
		c.ll.MoveToFront(el)
		return el.Value.(*entry).value, true
	}
	return 0, false
}

func (c *lruCache) Put(key string, value int) {
	if el, ok := c.items[key]; ok {
		el.Value.(*entry).value = value
		c.ll.MoveToFront(el)
		return
	}
	el := c.ll.PushFront(&entry{key: key, value: value})
	c.items[key] = el

	if c.ll.Len() > c.capacity {
		back := c.ll.Back()
		if back != nil {
			c.ll.Remove(back)
			delete(c.items, back.Value.(*entry).key)
		}
	}
}

func main() {
	fmt.Println("--- Проверка 02_lru_cache (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	c := newLRU(2)

	c.Put("a", 1)
	c.Put("b", 2)

	if v, ok := c.Get("a"); !ok || v != 1 {
		check(false, "Get(a) должен вернуть 1")
	}

	c.Put("c", 3)

	if _, ok := c.Get("b"); ok {
		check(false, "'b' должен быть вытеснен как least recently used")
	}
	if v, ok := c.Get("a"); !ok || v != 1 {
		check(false, "'a' должен остаться (был освежён)")
	}
	if v, ok := c.Get("c"); !ok || v != 3 {
		check(false, "'c' должен присутствовать")
	}

	c.Put("a", 10)
	if v, ok := c.Get("a"); !ok || v != 10 {
		check(false, "Put по существующему ключу должен обновлять значение")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
