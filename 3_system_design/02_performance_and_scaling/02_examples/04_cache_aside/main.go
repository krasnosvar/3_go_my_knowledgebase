// Пример 04: паттерн cache-aside поверх «медленной БД».
//
// «БД» — функция с time.Sleep, имитирующая дорогой запрос. Кэш — обычная
// map под RWMutex. На чтении: сначала смотрим в кэш; при промахе грузим из
// «БД», кладём в кэш и возвращаем; при повторном чтении того же ключа —
// попадание без обращения к «БД».
//
// Демонстрирует: цикл miss -> load -> set -> hit, разницу во времени и то,
// что «БД» на повторном ключе не вызывается (счётчик загрузок).
//
// Запуск: go run ./02_examples/04_cache_aside/main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

// slowDB имитирует медленный источник истины.
type slowDB struct {
	latency time.Duration
	mu      sync.Mutex
	calls   int // сколько раз реально сходили в «БД»
}

func (db *slowDB) load(key string) string {
	db.mu.Lock()
	db.calls++
	db.mu.Unlock()

	time.Sleep(db.latency) // «дорогой» запрос
	return "value-of-" + key
}

// cache — простой in-memory кэш (cache-aside им управляет приложение).
type cache struct {
	mu sync.RWMutex
	m  map[string]string
}

func newCache() *cache { return &cache{m: make(map[string]string)} }

func (c *cache) get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.m[key]
	return v, ok
}

func (c *cache) set(key, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = val
}

// getUser реализует cache-aside: кэш -> (при промахе) БД -> запись в кэш.
// Возвращает значение, признак попадания и затраченное время.
func getUser(c *cache, db *slowDB, key string) (val string, hit bool, took time.Duration) {
	start := time.Now()

	if v, ok := c.get(key); ok {
		return v, true, time.Since(start) // (1) hit
	}

	v := db.load(key) // (2) miss -> загрузка из «БД»
	c.set(key, v)     // (3) записываем в кэш (для реального сервиса — с TTL)
	return v, false, time.Since(start)
}

func main() {
	fmt.Println("--- Пример cache-aside ---")

	db := &slowDB{latency: 50 * time.Millisecond}
	c := newCache()

	// Первое чтение ключа "42": промах -> идём в медленную «БД».
	v1, hit1, t1 := getUser(c, db, "42")
	fmt.Printf("1) key=42 hit=%-5v took=%-6v value=%q\n", hit1, t1.Round(time.Millisecond), v1)

	// Повторное чтение того же ключа: попадание -> «БД» не трогаем.
	v2, hit2, t2 := getUser(c, db, "42")
	fmt.Printf("2) key=42 hit=%-5v took=%-6v value=%q\n", hit2, t2.Round(time.Millisecond), v2)

	// Другой ключ: снова промах -> загрузка.
	v3, hit3, t3 := getUser(c, db, "7")
	fmt.Printf("3) key=7  hit=%-5v took=%-6v value=%q\n", hit3, t3.Round(time.Millisecond), v3)

	fmt.Printf("\nОбращений к «БД»: %d (два уникальных ключа -> ровно 2 загрузки)\n", db.calls)
	if !hit1 && hit2 && t2 < t1 {
		fmt.Println("Промах медленный, попадание почти мгновенное — кэш работает.")
	}
	fmt.Println("Готово, программа завершается.")
}
