// Упражнение 02: разбор query-параметров пагинации.
//
// Задача: реализуйте parsePage(q url.Values) (page, limit int, err error):
//   - отсутствующий page → 1, отсутствующий limit → 20 (значения по умолчанию);
//   - page должен быть целым >= 1, limit — целым в диапазоне 1..100;
//   - нечисловое или вне диапазона значение → ненулевая ошибка.
//
// Запуск: go run ./03_exercises/02_query_params
package main

import (
	"fmt"
	"net/url"
	"strconv"
)

var _ = strconv.Atoi

func parsePage(q url.Values) (page, limit int, err error) {
	// TODO: прочитайте q.Get("page") и q.Get("limit"), подставьте значения
	// по умолчанию и провалидируйте диапазоны через strconv.Atoi.
	return 0, 0, nil
}

func main() {
	fmt.Println("--- Проверка 02_query_params ---")

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
