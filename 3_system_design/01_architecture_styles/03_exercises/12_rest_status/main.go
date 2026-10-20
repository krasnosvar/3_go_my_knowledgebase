// Упражнение 12: REST-хэндлер с правильными статус-кодами.
//
// Задача: реализовать userHandler для in-memory хранилища users
// (уже создан пользователь с id=1). Обслуживаем путь /users/{id}:
//
//	GET  существующий id   -> 200 + JSON пользователя;
//	GET  несуществующий id -> 404;
//	DELETE существующий id -> 204 (без тела), пользователь удаляется;
//	любой другой метод     -> 405 + заголовок Allow: "GET, DELETE".
//
// Разбор пути и id уже сделан за вас — заполните ветки switch по методу.
//
// Запуск: go run ./03_exercises/12_rest_status/main.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	mu    sync.Mutex
	users = map[int]User{1: {ID: 1, Name: "Alice"}}
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Разбор /users/{id} уже готов.
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 2 || parts[0] != "users" {
		http.NotFound(w, r)
		return
	}
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	_ = id // удалите, когда используете id в ветках ниже

	switch r.Method {
	case http.MethodGet:
		// TODO: mu.Lock(); u, ok := users[id]; mu.Unlock();
		//   если !ok -> http.NotFound(w, r); return
		//   иначе -> Content-Type application/json, статус 200,
		//            json.NewEncoder(w).Encode(u)
	case http.MethodDelete:
		// TODO: mu.Lock(); _, ok := users[id]; delete(users, id); mu.Unlock();
		//   если !ok -> http.NotFound(w, r); return
		//   иначе -> w.WriteHeader(http.StatusNoContent)
	default:
		// TODO: w.Header().Set("Allow", "GET, DELETE");
		//   http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// var _ удерживают импорты используемыми, пока TODO не сделан.
var (
	_ = json.NewEncoder
	_ = io.ReadAll
)

func main() {
	fmt.Println("--- Проверка 12_rest_status ---")

	srv := httptest.NewServer(http.HandlerFunc(userHandler))
	defer srv.Close()

	allOK := true

	// 1) GET существующего -> 200 + тело.
	st, body, _ := do(srv.URL+"/users/1", http.MethodGet)
	if st != http.StatusOK {
		allOK = false
		fmt.Printf("FAIL: GET /users/1 -> %d, ожидался 200\n", st)
	}
	var u User
	if json.Unmarshal([]byte(body), &u) != nil || u.Name != "Alice" {
		allOK = false
		fmt.Printf("FAIL: GET /users/1 тело=%q, ожидался пользователь Alice\n", body)
	}

	// 2) GET несуществующего -> 404.
	if st, _, _ = do(srv.URL+"/users/999", http.MethodGet); st != http.StatusNotFound {
		allOK = false
		fmt.Printf("FAIL: GET /users/999 -> %d, ожидался 404\n", st)
	}

	// 3) DELETE существующего -> 204.
	if st, _, _ = do(srv.URL+"/users/1", http.MethodDelete); st != http.StatusNoContent {
		allOK = false
		fmt.Printf("FAIL: DELETE /users/1 -> %d, ожидался 204\n", st)
	}
	// после удаления GET -> 404
	if st, _, _ = do(srv.URL+"/users/1", http.MethodGet); st != http.StatusNotFound {
		allOK = false
		fmt.Printf("FAIL: GET /users/1 после удаления -> %d, ожидался 404\n", st)
	}

	// 4) POST (не поддержан) -> 405 + Allow.
	st, _, allow := do(srv.URL+"/users/1", http.MethodPost)
	if st != http.StatusMethodNotAllowed {
		allOK = false
		fmt.Printf("FAIL: POST /users/1 -> %d, ожидался 405\n", st)
	}
	if allow != "GET, DELETE" {
		allOK = false
		fmt.Printf("FAIL: заголовок Allow=%q, ожидался \"GET, DELETE\"\n", allow)
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

// do выполняет запрос и возвращает статус, тело и заголовок Allow.
func do(url, method string) (int, string, string) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatalf("new request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("%s %s: %v", method, url, err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, strings.TrimSpace(string(b)), resp.Header.Get("Allow")
}
