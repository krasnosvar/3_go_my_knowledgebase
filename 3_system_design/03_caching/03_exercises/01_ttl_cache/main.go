// Упражнение 01: кэш с TTL-протуханием.
//
// Задача: реализуйте методы Set и Get у ttlCache так, чтобы:
//   - Set(key, val, ttl) сохранял значение с моментом истечения now+ttl;
//   - Get(key) возвращал (val, true), если ключ есть И не протух;
//   - протухший ключ трактовался как отсутствующий: (zero, false)
//     (можно удалять его лениво при обращении).
//
// Время берите через переданные в структуру часы now() — это делает
// тест детерминированным (без реального сна). Текущее «время» задаётся
// полем-счётчиком, тест двигает его вручную.
//
// Запуск: go run ./03_exercises/01_ttl_cache/main.go
package main

import (
	"fmt"
	"time"
)

// Держим импорт time живым в заготовке (в решении TTL — это time.Duration).
var _ = time.Second

type item struct {
	value   string
	expires int64 // момент истечения в «тиках» (см. now)
}

// ttlCache — кэш с протуханием по TTL. now() возвращает текущее «время».
type ttlCache struct {
	m   map[string]item
	now func() int64
}

func newTTLCache(now func() int64) *ttlCache {
	return &ttlCache{m: make(map[string]item), now: now}
}

// Set сохраняет значение с истечением через ttl тиков от текущего времени.
func (c *ttlCache) Set(key, val string, ttl int64) {
	// TODO: сохраните значение с полем expires = c.now() + ttl.
}

// Get возвращает значение, если ключ есть и не протух.
func (c *ttlCache) Get(key string) (string, bool) {
	// TODO: верните (val, true) только если ключ есть и c.now() < expires.
	// Протухший ключ считается отсутствующим.
	return "", false
}

func main() {
	fmt.Println("--- Проверка 01_ttl_cache ---")

	var clock int64 // текущее «время» в тиках
	c := newTTLCache(func() int64 { return clock })

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	clock = 0
	c.Set("a", "A", 10) // протухнет в момент 10

	if v, ok := c.Get("a"); !ok || v != "A" {
		check(false, "сразу после Set ключ 'a' должен читаться как \"A\"")
	}

	clock = 9 // ещё жив
	if _, ok := c.Get("a"); !ok {
		check(false, "в момент 9 ключ 'a' ещё не протух")
	}

	clock = 10 // истёк (now >= expires)
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
