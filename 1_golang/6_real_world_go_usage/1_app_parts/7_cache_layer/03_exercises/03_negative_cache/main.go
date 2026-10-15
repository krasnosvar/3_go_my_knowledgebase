// Упражнение 03: negative caching.
//
// Если клиенты массово спрашивают несуществующие ключи, каждый промах
// бьёт в базу. Решение — кешировать и «нет такого» (с коротким TTL).
//
// Задача: реализуйте (c *cache) Get(id int) (string, bool):
//   - запись в c.entries хранит found и expiresAt; если запись есть и
//     c.now < expiresAt — ответить из кеша (в т.ч. «не найдено»);
//   - иначе спросить c.load(id); найденное кешировать на hitTTL,
//     ненайденное — на missTTL.
//
// Время — поле c.now (тики), тест двигает его вручную.
//
// Запуск: go run ./03_exercises/03_negative_cache
package main

import "fmt"

const (
	hitTTL  = 100
	missTTL = 10
)

type entry struct {
	value     string
	found     bool
	expiresAt int64
}

type cache struct {
	entries map[int]entry
	load    func(id int) (string, bool)
	now     int64
}

func (c *cache) Get(id int) (string, bool) {
	// TODO: проверьте свежую запись; иначе load и сохранение с нужным TTL.
	return c.load(id)
}

func main() {
	fmt.Println("--- Проверка 03_negative_cache ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	loads := 0
	c := &cache{entries: map[int]entry{}, load: func(id int) (string, bool) {
		loads++
		if id == 1 {
			return "one", true
		}
		return "", false
	}}

	for i := 0; i < 5; i++ {
		_, ok := c.Get(404)
		check(!ok, "404 не найден")
	}
	check(loads == 1, fmt.Sprintf("5 промахов по 404 → одна загрузка (было %d)", loads))
	c.now = 10
	c.Get(404)
	check(loads == 2, "после missTTL отрицательная запись протухла")
	c.Get(1)
	c.now = 50
	v, ok := c.Get(1)
	check(ok && v == "one" && loads == 3, "положительная запись живёт hitTTL")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
