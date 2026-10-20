// Упражнение 10: семантика HTTP-методов (safe / idempotent) (Решение).
package main

import (
	"fmt"
	"net/http"
)

func classify(method string) (safe, idempotent bool) {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		return true, true
	case http.MethodPut, http.MethodDelete:
		return false, true
	default: // POST, PATCH, неизвестные методы
		return false, false
	}
}

func main() {
	fmt.Println("--- Проверка 10_method_semantics (Решение) ---")

	type want struct{ safe, idem bool }
	cases := map[string]want{
		http.MethodGet:     {true, true},
		http.MethodHead:    {true, true},
		http.MethodOptions: {true, true},
		http.MethodPut:     {false, true},
		http.MethodDelete:  {false, true},
		http.MethodPost:    {false, false},
		http.MethodPatch:   {false, false},
		"WAT":              {false, false},
	}

	allOK := true
	for m, w := range cases {
		safe, idem := classify(m)
		if safe != w.safe || idem != w.idem {
			allOK = false
			fmt.Printf("FAIL: %-7s -> safe=%v idem=%v, ожидалось safe=%v idem=%v\n",
				m, safe, idem, w.safe, w.idem)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
