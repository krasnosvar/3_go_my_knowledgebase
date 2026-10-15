// Упражнение 07: пагинация списка в API (Решение).
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

func intParam(r *http.Request, name string, def, lo, hi int) (int, bool) {
	s := r.URL.Query().Get(name)
	if s == "" {
		return def, true
	}
	n, err := strconv.Atoi(s)
	if err != nil || n < lo || n > hi {
		return 0, false
	}
	return n, true
}

func listItems(items []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, ok1 := intParam(r, "page", 1, 1, 1<<30)
		size, ok2 := intParam(r, "size", 2, 1, 50)
		if !ok1 || !ok2 {
			http.Error(w, "bad pagination", http.StatusBadRequest)
			return
		}
		w.Header().Set("X-Total-Count", strconv.Itoa(len(items)))
		start := (page - 1) * size
		if start >= len(items) {
			return
		}
		end := min(start+size, len(items))
		fmt.Fprint(w, strings.Join(items[start:end], ","))
	}
}

func main() {
	fmt.Println("--- Проверка 07_pagination (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	h := listItems([]string{"a", "b", "c", "d", "e"})
	get := func(q string) *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		h(rec, httptest.NewRequest("GET", "/items?"+q, nil))
		return rec
	}

	rec := get("")
	check(rec.Code == 200 && rec.Body.String() == "a,b", "по умолчанию page=1,size=2 → a,b")
	check(rec.Header().Get("X-Total-Count") == "5", "X-Total-Count = 5")
	check(get("page=3").Body.String() == "e", "page=3 → e")
	check(get("page=2&size=3").Body.String() == "d,e", "page=2&size=3 → d,e")
	rec = get("page=9")
	check(rec.Code == 200 && rec.Body.String() == "", "за пределами → пусто")
	check(get("size=0").Code == 400 && get("page=-1").Code == 400, "некорректные параметры → 400")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
