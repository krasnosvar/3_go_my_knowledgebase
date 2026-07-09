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
	// TODO: Инициализация нового объекта LRU
	return nil
}

func (l *LRUCache) Get(key int) int {
	// TODO: Если элемент есть, верните его значение и сделайте "свежим" (переместите в начало списка).
	// Иначе верните -1.
	return -1
}

func (l *LRUCache) Put(key, value int) {
	// TODO: Обновите или добавьте элемент.
	// Если лимит емкости превышен, удалите наименее используемый элемент.
}

func main() {
	fmt.Println("--- Проверка 08_maps_lru_cache ---")

	lru := NewLRU(2)
	if lru == nil {
		fmt.Println("Тест LRU Cache: FAIL (NewLRU вернул nil)")
		return
	}

	lru.Put(1, 10) // {1: 10}
	lru.Put(2, 20) // {1: 10, 2: 20}

	t1 := lru.Get(1) == 10 // {2: 20, 1: 10} — 1 стал свежим

	lru.Put(3, 30)         // Превышение емкости! Вытесняет 2. {1: 10, 3: 30}
	t2 := lru.Get(2) == -1 // 2 удален
	t3 := lru.Get(3) == 30

	lru.Put(4, 40)         // Превышение емкости! Вытесняет 1. {3: 30, 4: 40}
	t4 := lru.Get(1) == -1 // 1 удален
	t5 := lru.Get(4) == 40

	if t1 && t2 && t3 && t4 && t5 {
		fmt.Println("Тест LRU Cache: PASS")
	} else {
		fmt.Printf("Тест LRU Cache: FAIL (t1=%t, t2=%t, t3=%t, t4=%t, t5=%t)\n", t1, t2, t3, t4, t5)
	}
}
