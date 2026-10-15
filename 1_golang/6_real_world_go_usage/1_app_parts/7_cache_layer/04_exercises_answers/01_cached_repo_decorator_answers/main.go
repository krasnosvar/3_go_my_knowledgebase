// Упражнение 01: кеш как декоратор репозитория (Решение).
package main

import (
	"errors"
	"fmt"
)

type repo interface {
	Get(id int) (string, error)
}

// slowRepo — «база»: считает обращения.
type slowRepo struct {
	data  map[int]string
	calls int
}

func (s *slowRepo) Get(id int) (string, error) {
	s.calls++
	v, ok := s.data[id]
	if !ok {
		return "", errors.New("not found")
	}
	return v, nil
}

type cachedRepo struct {
	next  repo
	cache map[int]string
}

func (c *cachedRepo) Get(id int) (string, error) {
	if v, ok := c.cache[id]; ok {
		return v, nil
	}
	v, err := c.next.Get(id)
	if err != nil {
		return "", err
	}
	c.cache[id] = v
	return v, nil
}

func main() {
	fmt.Println("--- Проверка 01_cached_repo_decorator (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	db := &slowRepo{data: map[int]string{1: "one"}}
	var r repo = &cachedRepo{next: db, cache: map[int]string{}}

	for i := 0; i < 3; i++ {
		v, err := r.Get(1)
		check(err == nil && v == "one", "Get(1) → one")
	}
	check(db.calls == 1, fmt.Sprintf("три чтения id=1 → одно обращение к базе (было %d)", db.calls))
	_, err1 := r.Get(2)
	_, err2 := r.Get(2)
	check(err1 != nil && err2 != nil, "ошибка источника возвращается")
	check(db.calls == 3, "ошибки не кешируются: каждый промах идёт в базу")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
