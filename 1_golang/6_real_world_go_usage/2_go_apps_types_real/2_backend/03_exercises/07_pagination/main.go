// Упражнение 07: пагинация списка в API.
//
// Задача: реализуйте handler listItems(items []string):
//   - query: page (по умолчанию 1, >= 1), size (по умолчанию 2, 1..50);
//     некорректные → 400;
//   - тело — элементы страницы через запятую (пустая строка за пределами);
//   - заголовок X-Total-Count — общее число элементов.
//
// Запуск: go run ./03_exercises/07_pagination
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

var _ = strings.Join

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
		// TODO: page/size через intParam, срез страницы, заголовок X-Total-Count.
	}
}

func main() {
	fmt.Println("--- Проверка 07_pagination ---")

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
