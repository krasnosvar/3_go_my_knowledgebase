// Упражнение 04: разбор query-параметров фильтра (Решение).
package main

import (
	"fmt"
	"net/url"
	"strconv"
)

type Filter struct {
	Q      string
	Limit  int
	Active bool
}

func parseFilter(raw string) Filter {
	vals, err := url.ParseQuery(raw)
	if err != nil {
		return Filter{Limit: 20}
	}

	f := Filter{Q: vals.Get("q"), Limit: 20}

	if n, err := strconv.Atoi(vals.Get("limit")); err == nil && n >= 1 {
		f.Limit = n
	}
	if b, err := strconv.ParseBool(vals.Get("active")); err == nil {
		f.Active = b
	}
	return f
}

func main() {
	fmt.Println("--- Проверка 04_parse_filter (Решение) ---")

	allOK := true
	check := func(raw string, want Filter) {
		got := parseFilter(raw)
		if got != want {
			allOK = false
			fmt.Printf("FAIL: %q\n  got : %+v\n  want: %+v\n", raw, got, want)
		}
	}

	check("q=alice&limit=50&active=true",
		Filter{Q: "alice", Limit: 50, Active: true})

	check("", Filter{Q: "", Limit: 20, Active: false})

	check("limit=abc&active=1",
		Filter{Q: "", Limit: 20, Active: true})

	check("limit=0", Filter{Q: "", Limit: 20, Active: false})

	check("q=hello%20world", Filter{Q: "hello world", Limit: 20, Active: false})

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
