// Упражнение 02: LRU-кэш (Put/Get с вытеснением).
//
// Задача: реализуйте методы Get и Put у lruCache фиксированной ёмкости:
//   - Get(key) возвращает (value, true) и делает ключ самым свежим;
//     если ключа нет — (0, false);
//   - Put(key, value) вставляет/обновляет ключ (делая его самым свежим);
//     если после вставки размер превысил capacity — вытесняет самый
//     давно использованный (least recently used) ключ.
//
// Подсказка: map[string]*list.Element для O(1) доступа + container/list
// для порядка использования. Голова списка — свежий, хвост — LRU.
//
// Запуск: go run ./03_exercises/02_lru_cache/main.go
package main

import (
	"container/list"
	"fmt"
)

// Держим импорт container/list живым в заготовке.
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

// Get возвращает значение и освежает ключ (делает most recently used).
func (c *lruCache) Get(key string) (int, bool) {
	// TODO: найдите элемент в items; если есть — MoveToFront и верните value.
	return 0, false
}

// Put вставляет/обновляет ключ; при переполнении вытесняет LRU-элемент.
func (c *lruCache) Put(key string, value int) {
	// TODO:
	//  1) если ключ уже есть — обновите value и MoveToFront;
	//  2) иначе PushFront нового entry и запишите в items;
	//  3) если ll.Len() > capacity — удалите ll.Back() и его ключ из items.
}

func main() {
	fmt.Println("--- Проверка 02_lru_cache ---")

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

	if v, ok := c.Get("a"); !ok || v != 1 { // освежаем "a" -> LRU теперь "b"
		check(false, "Get(a) должен вернуть 1")
	}

	c.Put("c", 3) // переполнение -> вытесняется "b"

	if _, ok := c.Get("b"); ok {
		check(false, "'b' должен быть вытеснен как least recently used")
	}
	if v, ok := c.Get("a"); !ok || v != 1 {
		check(false, "'a' должен остаться (был освежён)")
	}
	if v, ok := c.Get("c"); !ok || v != 3 {
		check(false, "'c' должен присутствовать")
	}

	c.Put("a", 10) // обновление существующего ключа
	if v, ok := c.Get("a"); !ok || v != 10 {
		check(false, "Put по существующему ключу должен обновлять значение")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
