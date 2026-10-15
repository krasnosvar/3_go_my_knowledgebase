// Упражнение 03: negative caching (Решение).
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
	if e, ok := c.entries[id]; ok && c.now < e.expiresAt {
		return e.value, e.found
	}
	v, found := c.load(id)
	ttl := int64(missTTL)
	if found {
		ttl = hitTTL
	}
	c.entries[id] = entry{value: v, found: found, expiresAt: c.now + ttl}
	return v, found
}

func main() {
	fmt.Println("--- Проверка 03_negative_cache (Решение) ---")

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
