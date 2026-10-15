// Упражнение 03: REST-хэндлер с правильными статус-кодами (Решение).
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

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		u, ok := users[id]
		mu.Unlock()
		if !ok {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(u)

	case http.MethodDelete:
		mu.Lock()
		_, ok := users[id]
		delete(users, id)
		mu.Unlock()
		if !ok {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, DELETE")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	fmt.Println("--- Проверка 03_rest_status (Решение) ---")

	srv := httptest.NewServer(http.HandlerFunc(userHandler))
	defer srv.Close()

	allOK := true

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

	if st, _, _ = do(srv.URL+"/users/999", http.MethodGet); st != http.StatusNotFound {
		allOK = false
		fmt.Printf("FAIL: GET /users/999 -> %d, ожидался 404\n", st)
	}

	if st, _, _ = do(srv.URL+"/users/1", http.MethodDelete); st != http.StatusNoContent {
		allOK = false
		fmt.Printf("FAIL: DELETE /users/1 -> %d, ожидался 204\n", st)
	}
	if st, _, _ = do(srv.URL+"/users/1", http.MethodGet); st != http.StatusNotFound {
		allOK = false
		fmt.Printf("FAIL: GET /users/1 после удаления -> %d, ожидался 404\n", st)
	}

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
