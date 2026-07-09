// Упражнение 04: инвалидация кэша по ключу (Решение).
package main

import "fmt"

type store struct {
	db      map[string]string
	cache   map[string]string
	dbReads int
}

func newStore() *store {
	return &store{
		db:    map[string]string{"k": "v1"},
		cache: make(map[string]string),
	}
}

func (s *store) Get(key string) (string, bool) {
	if v, ok := s.cache[key]; ok {
		return v, true // попадание — db не трогаем
	}
	v, ok := s.db[key]
	if !ok {
		return "", false
	}
	s.dbReads++
	s.cache[key] = v
	return v, true
}

func (s *store) Update(key, val string) {
	s.db[key] = val
	delete(s.cache, key) // инвалидация: следующий Get возьмёт свежее
}

func main() {
	fmt.Println("--- Проверка 04_invalidation (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	s := newStore()

	if v, ok := s.Get("k"); !ok || v != "v1" {
		check(false, "первый Get должен вернуть v1 из db")
	}
	check(s.dbReads == 1, fmt.Sprintf("после первого Get dbReads=1, получено %d", s.dbReads))

	if v, ok := s.Get("k"); !ok || v != "v1" {
		check(false, "повторный Get должен вернуть v1 из кэша")
	}
	check(s.dbReads == 1, "повторный Get не должен читать db (должно быть попадание)")

	s.Update("k", "v2")

	if v, ok := s.Get("k"); !ok || v != "v2" {
		check(false, "после Update Get должен вернуть свежее v2")
	}
	check(s.dbReads == 2, fmt.Sprintf("после инвалидации Get должен снова читать db, dbReads=2, получено %d", s.dbReads))

	if _, ok := s.Get("nope"); ok {
		check(false, "несуществующий ключ не должен находиться")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
