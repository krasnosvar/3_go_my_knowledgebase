// Упражнение 03: разбор query-параметров пагинации.
//
// Задача: реализуйте parsePaging(q), где q — r.URL.Query():
//   - page: по умолчанию 1, допустимо >= 1;
//   - limit: по умолчанию 20, допустимо 1..100;
//   - параметр задан, но не число или вне диапазона → ошибка
//     (хендлер потом ответит 400 Bad Request);
//   - отсутствие параметра определяйте через q.Has(name) или q.Get(name) == "".
//
// Запуск: go run ./03_exercises/03_query_params/main.go
package main

import (
	"fmt"
	"net/url"
	"strconv"
)

var _ = strconv.Atoi // держим импорт живым в заготовке

// parsePaging извлекает page и limit из query-параметров.
func parsePaging(q url.Values) (page, limit int, err error) {
	// TODO: значения по умолчанию, strconv.Atoi, проверки диапазонов.
	return 0, 0, nil
}

func main() {
	fmt.Println("--- Проверка 03_query_params ---")

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
