// Пример 02: LRU-кэш с ограничением ёмкости и вытеснением.
//
// Классическая реализация: map для O(1) доступа + container/list
// (двусвязный список) для порядка использования. На каждый Get/Put
// элемент переносится в голову списка (most recently used); при
// переполнении удаляется хвост (least recently used).
//
// Демонстрирует: как при заполнении кэша вытесняется давно не
// использованный ключ и как Get «спасает» ключ от вытеснения.
//
// Запуск: go run ./02_examples/02_lru_eviction/main.go
package main

import (
	"container/list"
	"fmt"
)

// entry хранится как значение элемента списка (нужен ключ, чтобы при
// удалении хвоста знать, что удалять из map).
type entry struct {
	key   string
	value int
}

// lruCache — LRU-кэш фиксированной ёмкости.
type lruCache struct {
	capacity int
	ll       *list.List               // порядок: голова = свежий, хвост = старый
	items    map[string]*list.Element // key -> узел списка
}

func newLRU(capacity int) *lruCache {
	return &lruCache{
		capacity: capacity,
		ll:       list.New(),
		items:    make(map[string]*list.Element),
	}
}

// Get возвращает значение и помечает ключ как недавно использованный.
func (c *lruCache) Get(key string) (int, bool) {
	if el, ok := c.items[key]; ok {
		c.ll.MoveToFront(el) // освежаем: теперь это most recently used
		return el.Value.(*entry).value, true
	}
	return 0, false
}

// Put вставляет/обновляет ключ; при переполнении вытесняет LRU-элемент.
// Возвращает вытесненный ключ (или "" если вытеснения не было).
func (c *lruCache) Put(key string, value int) (evicted string) {
	if el, ok := c.items[key]; ok {
		el.Value.(*entry).value = value
		c.ll.MoveToFront(el)
		return ""
	}

	el := c.ll.PushFront(&entry{key: key, value: value})
	c.items[key] = el

	if c.ll.Len() > c.capacity {
		back := c.ll.Back() // least recently used
		if back != nil {
			e := back.Value.(*entry)
			c.ll.Remove(back)
			delete(c.items, e.key)
			return e.key
		}
	}
	return ""
}

// keysMRUtoLRU возвращает ключи в порядке от свежего к старому (для наглядности).
func (c *lruCache) keysMRUtoLRU() []string {
	out := make([]string, 0, c.ll.Len())
	for el := c.ll.Front(); el != nil; el = el.Next() {
		out = append(out, el.Value.(*entry).key)
	}
	return out
}

func main() {
	fmt.Println("--- Пример LRU-вытеснения (ёмкость 3) ---")

	c := newLRU(3)

	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)
	fmt.Printf("после Put a,b,c:            %v (свежий -> старый)\n", c.keysMRUtoLRU())

	// Обращаемся к "a" — теперь "b" самый давний.
	if v, ok := c.Get("a"); ok {
		fmt.Printf("Get(a)=%d -> a освежён:      %v\n", v, c.keysMRUtoLRU())
	}

	// Вставляем "d": ёмкость превышена -> вытесняется хвост ("b").
	ev := c.Put("d", 4)
	fmt.Printf("Put(d) -> вытеснен %q:       %v\n", ev, c.keysMRUtoLRU())

	// "b" больше нет в кэше.
	_, okB := c.Get("b")
	// "a", "c", "d" — на месте.
	_, okA := c.Get("a")
	_, okC := c.Get("c")
	_, okD := c.Get("d")

	fmt.Printf("\nналичие: a=%v c=%v d=%v b=%v\n", okA, okC, okD, okB)
	if ev == "b" && !okB && okA && okC && okD {
		fmt.Println("LRU вытеснил именно давно не использованный ключ — верно.")
	}
	fmt.Println("Готово, программа завершается.")
}
