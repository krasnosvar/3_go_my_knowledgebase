// Упражнение 02: инвалидация кеша на записи (Решение).
package main

import "fmt"

type service struct {
	db      map[int]string
	cache   map[int]string
	dbReads int
}

func (s *service) Get(id int) string {
	if v, ok := s.cache[id]; ok {
		return v
	}
	s.dbReads++
	v := s.db[id]
	s.cache[id] = v
	return v
}

func (s *service) Update(id int, v string) {
	s.db[id] = v
	delete(s.cache, id)
}

func main() {
	fmt.Println("--- Проверка 02_invalidate_on_write (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	s := &service{db: map[int]string{1: "old"}, cache: map[int]string{}}
	check(s.Get(1) == "old", "первое чтение → old")
	s.Update(1, "new")
	_, cached := s.cache[1]
	check(!cached, "после Update ключ удалён из кеша")
	check(s.Get(1) == "new", "чтение после Update → new (не устаревшее)")
	check(s.dbReads == 2, "после инвалидации было повторное чтение из базы")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
