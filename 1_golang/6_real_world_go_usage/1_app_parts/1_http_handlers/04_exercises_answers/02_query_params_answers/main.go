// Упражнение 02: разбор query-параметров пагинации (Решение).
package main

import (
	"fmt"
	"net/url"
	"strconv"
)

func parsePage(q url.Values) (page, limit int, err error) {
	page, limit = 1, 20
	if s := q.Get("page"); s != "" {
		if page, err = strconv.Atoi(s); err != nil || page < 1 {
			return 0, 0, fmt.Errorf("bad page %q", s)
		}
	}
	if s := q.Get("limit"); s != "" {
		if limit, err = strconv.Atoi(s); err != nil || limit < 1 || limit > 100 {
			return 0, 0, fmt.Errorf("bad limit %q", s)
		}
	}
	return page, limit, nil
}

func main() {
	fmt.Println("--- Проверка 02_query_params (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	parse := func(raw string) (int, int, error) {
		q, _ := url.ParseQuery(raw)
		return parsePage(q)
	}

	p, l, err := parse("")
	check(err == nil && p == 1 && l == 20, "пустой запрос → page=1, limit=20")
	p, l, err = parse("page=3&limit=50")
	check(err == nil && p == 3 && l == 50, "page=3&limit=50 разбирается")
	_, _, err = parse("page=abc")
	check(err != nil, "page=abc → ошибка")
	_, _, err = parse("page=0")
	check(err != nil, "page=0 → ошибка")
	_, _, err = parse("limit=1000")
	check(err != nil, "limit=1000 → ошибка")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
