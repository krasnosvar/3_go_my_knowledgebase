// Упражнение 01: маршруты REST-ресурса на stdlib ServeMux (Go 1.22+).
//
// Задача: в (a *api) routes() зарегистрируйте:
//   - "GET /users"       → 200, тело — имена через запятую в порядке ID ("ann,bob");
//   - "GET /users/{id}"  → 200 и имя; нечисловой id → 400; нет такого → 404;
//   - "POST /users"      → имя берётся из тела запроса (io.ReadAll), пустое → 400,
//     иначе добавить в a.users и ответить 201.
//
// Запуск: go run ./03_exercises/01_routes
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

type api struct{ users []string } // индекс+1 = ID

var _ = io.ReadAll
var _ = strconv.Atoi

func (a *api) routes() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: зарегистрируйте три маршрута.
	return mux
}

func main() {
	fmt.Println("--- Проверка 01_routes ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	a := &api{users: []string{"ann"}}
	mux := a.routes()
	do := func(method, path, body string) *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest(method, path, strings.NewReader(body)))
		return rec
	}

	check(do("POST", "/users", "bob").Code == 201, "POST /users bob → 201")
	check(do("POST", "/users", "").Code == 400, "POST пустое имя → 400")
	rec := do("GET", "/users", "")
	check(rec.Code == 200 && rec.Body.String() == "ann,bob", "GET /users → ann,bob")
	rec = do("GET", "/users/2", "")
	check(rec.Code == 200 && rec.Body.String() == "bob", "GET /users/2 → bob")
	check(do("GET", "/users/9", "").Code == 404, "GET /users/9 → 404")
	check(do("GET", "/users/x", "").Code == 400, "GET /users/x → 400")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
