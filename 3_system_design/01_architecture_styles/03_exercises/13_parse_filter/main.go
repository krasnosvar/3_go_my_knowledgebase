// Упражнение 13: разбор query-параметров фильтра.
//
// Задача: реализовать parseFilter — по "сырой" query-строке (часть URL
// после '?') собрать структуру Filter:
//
//	q      — строковый поиск (параметр "q"), по умолчанию "";
//	limit  — целое (параметр "limit"), по умолчанию 20;
//	         при отсутствии/ошибке/значении < 1 -> 20;
//	active — булево (параметр "active"), по умолчанию false;
//	         "true"/"1" -> true, иначе false.
//
// Подсказка: url.ParseQuery(raw) -> url.Values; Values.Get(name);
//
//	strconv.Atoi, strconv.ParseBool.
//
// Запуск: go run ./03_exercises/13_parse_filter/main.go
package main

import (
	"fmt"
	"net/url"
	"strconv"
)

// Filter — распарсенные параметры фильтра.
type Filter struct {
	Q      string
	Limit  int
	Active bool
}

// parseFilter — реализуйте согласно условию.
func parseFilter(raw string) Filter {
	// TODO:
	//   1) vals, err := url.ParseQuery(raw); при ошибке верните дефолт
	//      Filter{Limit: 20};
	//   2) f := Filter{Q: vals.Get("q"), Limit: 20};
	//   3) limit: n, err := strconv.Atoi(vals.Get("limit"));
	//      если err == nil && n >= 1 -> f.Limit = n;
	//   4) active: b, err := strconv.ParseBool(vals.Get("active"));
	//      если err == nil -> f.Active = b.
	return Filter{}
}

// var _ удерживает импорт strconv используемым, пока TODO не сделан.
var _ = strconv.Atoi

func main() {
	fmt.Println("--- Проверка 13_parse_filter ---")

	allOK := true
	check := func(raw string, want Filter) {
		got := parseFilter(raw)
		if got != want {
			allOK = false
			fmt.Printf("FAIL: %q\n  got : %+v\n  want: %+v\n", raw, got, want)
		}
	}

	// Все параметры заданы.
	check("q=alice&limit=50&active=true",
		Filter{Q: "alice", Limit: 50, Active: true})

	// Пустая строка -> дефолты.
	check("", Filter{Q: "", Limit: 20, Active: false})

	// Кривой limit -> дефолт 20; active=1 -> true.
	check("limit=abc&active=1",
		Filter{Q: "", Limit: 20, Active: true})

	// limit < 1 -> дефолт 20.
	check("limit=0", Filter{Q: "", Limit: 20, Active: false})

	// URL-кодирование пробела в q.
	check("q=hello%20world", Filter{Q: "hello world", Limit: 20, Active: false})

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

// пример корректной query-строки для справки.
var _ = url.Values{}
