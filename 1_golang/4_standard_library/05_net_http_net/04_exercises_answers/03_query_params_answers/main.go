// Упражнение 03: разбор query-параметров пагинации (Решение).
package main

import (
	"fmt"
	"net/url"
	"strconv"
)

// parsePaging извлекает page и limit из query-параметров.
func parsePaging(q url.Values) (page, limit int, err error) {
	page, limit = 1, 20
	if s := q.Get("page"); s != "" {
		page, err = strconv.Atoi(s)
		if err != nil || page < 1 {
			return 0, 0, fmt.Errorf("invalid page %q", s)
		}
	}
	if s := q.Get("limit"); s != "" {
		limit, err = strconv.Atoi(s)
		if err != nil || limit < 1 || limit > 100 {
			return 0, 0, fmt.Errorf("invalid limit %q (want 1..100)", s)
		}
	}
	return page, limit, nil
}

func main() {
	fmt.Println("--- Проверка 03_query_params (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	parse := func(raw string) (int, int, error) {
		u, err := url.Parse(raw)
		if err != nil {
			return 0, 0, err
		}
		return parsePaging(u.Query())
	}

	p, l, err := parse("/items")
	check(err == nil && p == 1 && l == 20, fmt.Sprintf("по умолчанию: получено page=%d limit=%d err=%v", p, l, err))

	p, l, err = parse("/items?page=3&limit=50")
	check(err == nil && p == 3 && l == 50, fmt.Sprintf("page=3&limit=50: получено page=%d limit=%d err=%v", p, l, err))

	for _, bad := range []string{"/items?page=0", "/items?page=abc", "/items?limit=0", "/items?limit=101", "/items?limit=-5"} {
		if _, _, err := parse(bad); err == nil {
			check(false, fmt.Sprintf("%s: ожидалась ошибка", bad))
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
