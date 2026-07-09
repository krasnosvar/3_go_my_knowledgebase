package main

import (
	"container/list"
	"fmt"
)

type entry struct {
	key   int
	value int
}

type LRUCache struct {
	capacity  int
	evictList *list.List
	items     map[int]*list.Element
}

func NewLRU(capacity int) *LRUCache {
	return &LRUCache{
		capacity:  capacity,
		evictList: list.New(),
		items:     make(map[int]*list.Element),
	}
}

func (l *LRUCache) Get(key int) int {
	if elem, ok := l.items[key]; ok {
		l.evictList.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1
}

func (l *LRUCache) Put(key, value int) {
	if elem, ok := l.items[key]; ok {
		l.evictList.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	elem := l.evictList.PushFront(&entry{key: key, value: value})
	l.items[key] = elem

	if l.evictList.Len() > l.capacity {
		last := l.evictList.Back()
		if last != nil {
			l.evictList.Remove(last)
			delete(l.items, last.Value.(*entry).key)
		}
	}
}

func main() {
	fmt.Println("--- Проверка 08_maps_lru_cache (Решение) ---")

	lru := NewLRU(2)
	if lru == nil {
		fmt.Println("Тест LRU Cache: FAIL (NewLRU вернул nil)")
		return
	}

	lru.Put(1, 10)
	lru.Put(2, 20)

	t1 := lru.Get(1) == 10

	lru.Put(3, 30)
	t2 := lru.Get(2) == -1
	t3 := lru.Get(3) == 30

	lru.Put(4, 40)
	t4 := lru.Get(1) == -1
	t5 := lru.Get(4) == 40

	if t1 && t2 && t3 && t4 && t5 {
		fmt.Println("Тест LRU Cache: PASS")
	} else {
		fmt.Printf("Тест LRU Cache: FAIL (t1=%t, t2=%t, t3=%t, t4=%t, t5=%t)\n", t1, t2, t3, t4, t5)
	}
}
