// Упражнение 01: маршруты REST-ресурса на stdlib ServeMux (Решение).
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

func (a *api) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, strings.Join(a.users, ","))
	})
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}
		if id < 1 || id > len(a.users) {
			http.NotFound(w, r)
			return
		}
		fmt.Fprint(w, a.users[id-1])
	})
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil || len(body) == 0 {
			http.Error(w, "empty name", http.StatusBadRequest)
			return
		}
		a.users = append(a.users, string(body))
		w.WriteHeader(http.StatusCreated)
	})
	return mux
}

func main() {
	fmt.Println("--- Проверка 01_routes (Решение) ---")

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
