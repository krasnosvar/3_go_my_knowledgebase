// Пример 01: маленький REST-ресурс в памяти (CRUD).
//
// Поднимаем HTTP-сервер через httptest.NewServer (сам выбирает
// случайный порт на 127.0.0.1) с in-memory хранилищем пользователей.
// Обслуживаем /users и /users/{id} методами GET/POST/PUT/DELETE,
// прогоняем полный цикл CRUD и печатаем статус-коды — программа
// завершается сама, сервер не висит.
//
// Демонстрирует: маршрутизацию по методу, коды 200/201/204/404/405,
// заголовок Location при создании, кодирование/декодирование JSON.
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
	"time"
)

// User — ресурс, которым управляет наш REST-API.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// store — потокобезопасное in-memory хранилище пользователей.
type store struct {
	mu     sync.Mutex
	users  map[int]User
	nextID int
}

func newStore() *store {
	return &store{users: make(map[int]User), nextID: 1}
}

// handler маршрутизирует запрос по пути и методу.
func (s *store) handler(w http.ResponseWriter, r *http.Request) {
	// Разделяем /users и /users/{id}.
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/") // ["users"] или ["users","42"]

	switch {
	case len(parts) == 1 && parts[0] == "users":
		s.handleCollection(w, r)
	case len(parts) == 2 && parts[0] == "users":
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}
		s.handleItem(w, r, id)
	default:
		http.NotFound(w, r) // 404
	}
}

// handleCollection обслуживает /users: GET (список) и POST (создать).
func (s *store) handleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.mu.Lock()
		list := make([]User, 0, len(s.users))
		for _, u := range s.users {
			list = append(list, u)
		}
		s.mu.Unlock()
		writeJSON(w, http.StatusOK, list)

	case http.MethodPost:
		var in User
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "bad body", http.StatusBadRequest)
			return
		}
		s.mu.Lock()
		in.ID = s.nextID
		s.nextID++
		s.users[in.ID] = in
		s.mu.Unlock()
		// 201 Created + заголовок Location на новый ресурс.
		w.Header().Set("Location", fmt.Sprintf("/users/%d", in.ID))
		writeJSON(w, http.StatusCreated, in)

	default:
		// 405: метод не разрешён для этого ресурса.
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleItem обслуживает /users/{id}: GET, PUT, DELETE.
func (s *store) handleItem(w http.ResponseWriter, r *http.Request, id int) {
	switch r.Method {
	case http.MethodGet:
		s.mu.Lock()
		u, ok := s.users[id]
		s.mu.Unlock()
		if !ok {
			http.NotFound(w, r) // 404
			return
		}
		writeJSON(w, http.StatusOK, u)

	case http.MethodPut:
		var in User
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "bad body", http.StatusBadRequest)
			return
		}
		s.mu.Lock()
		_, ok := s.users[id]
		in.ID = id // ID берём из URL, а не из тела
		s.users[id] = in
		s.mu.Unlock()
		if !ok {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, http.StatusOK, in)

	case http.MethodDelete:
		s.mu.Lock()
		_, ok := s.users[id]
		delete(s.users, id)
		s.mu.Unlock()
		if !ok {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusNoContent) // 204: тела нет

	default:
		w.Header().Set("Allow", "GET, PUT, DELETE")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// writeJSON — хелпер: выставляет Content-Type, статус и пишет JSON.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func main() {
	srv := httptest.NewServer(http.HandlerFunc(newStore().handler))
	defer srv.Close()
	fmt.Printf("REST-сервер слушает на %s\n\n", srv.URL)

	client := &http.Client{Timeout: 5 * time.Second}

	// 1) POST /users -> 201 Created + Location.
	body := strings.NewReader(`{"name":"Alice"}`)
	resp := do(client, http.MethodPost, srv.URL+"/users", body)
	fmt.Printf("POST   /users        -> %d %s (Location: %s)\n",
		resp.status, resp.text, resp.location)

	// 2) GET /users/1 -> 200 OK.
	resp = do(client, http.MethodGet, srv.URL+"/users/1", nil)
	fmt.Printf("GET    /users/1      -> %d %s\n", resp.status, resp.text)

	// 3) PUT /users/1 -> 200 OK (замена целиком).
	resp = do(client, http.MethodPut, srv.URL+"/users/1", strings.NewReader(`{"name":"Alice Smith"}`))
	fmt.Printf("PUT    /users/1      -> %d %s\n", resp.status, resp.text)

	// 4) GET /users/999 -> 404 Not Found.
	resp = do(client, http.MethodGet, srv.URL+"/users/999", nil)
	fmt.Printf("GET    /users/999    -> %d (не найден)\n", resp.status)

	// 5) DELETE /users/1 -> 204 No Content.
	resp = do(client, http.MethodDelete, srv.URL+"/users/1", nil)
	fmt.Printf("DELETE /users/1      -> %d (тело пустое: %v)\n", resp.status, resp.text == "")

	// 6) GET /users/1 после удаления -> 404.
	resp = do(client, http.MethodGet, srv.URL+"/users/1", nil)
	fmt.Printf("GET    /users/1      -> %d (уже удалён)\n", resp.status)

	// 7) PATCH /users/1 -> 405 Method Not Allowed.
	resp = do(client, http.MethodPatch, srv.URL+"/users/1", nil)
	fmt.Printf("PATCH  /users/1      -> %d (метод не разрешён, Allow: %s)\n",
		resp.status, resp.allow)

	fmt.Println("\nГотово, программа завершается.")
}

type result struct {
	status   int
	text     string
	location string
	allow    string
}

// do выполняет запрос и собирает статус, тело и интересные заголовки.
func do(c *http.Client, method, url string, body io.Reader) result {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("new request: %v", err)
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("%s %s: %v", method, url, err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return result{
		status:   resp.StatusCode,
		text:     strings.TrimSpace(string(b)),
		location: resp.Header.Get("Location"),
		allow:    resp.Header.Get("Allow"),
	}
}
