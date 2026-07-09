// Пример 03: идемпотентность методов и Idempotency-Key.
//
// Сервер (httptest, случайный порт на 127.0.0.1) хранит пользователей.
// Показываем три вещи:
//  1. PUT идемпотентен: два одинаковых PUT дают одно и то же состояние;
//  2. "голый" POST не идемпотентен: два POST создают ДВА ресурса (дубликат);
//  3. POST с заголовком Idempotency-Key дедуплицируется: повтор с тем же
//     ключом возвращает тот же результат и НЕ создаёт дубликат.
//
// Программа сама прогоняет сценарий и завершается — сервер не висит.
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

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type store struct {
	mu       sync.Mutex
	users    map[int]User
	nextID   int
	idemKeys map[string]int // Idempotency-Key -> уже созданный ID
}

func newStore() *store {
	return &store{users: make(map[int]User), nextID: 1, idemKeys: make(map[string]int)}
}

// createPOST — создание. Не идемпотентно, если нет Idempotency-Key.
func (s *store) createPOST(w http.ResponseWriter, r *http.Request) {
	var in User
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "bad body", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Если пришёл Idempotency-Key и мы его уже видели — вернём прежний результат.
	if key := r.Header.Get("Idempotency-Key"); key != "" {
		if id, ok := s.idemKeys[key]; ok {
			writeJSON(w, http.StatusOK, s.users[id]) // повтор: тот же ресурс, без дубликата
			return
		}
		// Первый раз с этим ключом — создадим и запомним связь ключ->ID.
		in.ID = s.nextID
		s.nextID++
		s.users[in.ID] = in
		s.idemKeys[key] = in.ID
		writeJSON(w, http.StatusCreated, in)
		return
	}

	// Без ключа — каждый POST создаёт НОВЫЙ ресурс (возможен дубликат).
	in.ID = s.nextID
	s.nextID++
	s.users[in.ID] = in
	writeJSON(w, http.StatusCreated, in)
}

// putItem — замена ресурса по ID из URL. Идемпотентно по определению.
func (s *store) putItem(w http.ResponseWriter, r *http.Request, id int) {
	var in User
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "bad body", http.StatusBadRequest)
		return
	}
	s.mu.Lock()
	in.ID = id
	s.users[id] = in // повтор с тем же телом -> то же состояние
	if id >= s.nextID {
		s.nextID = id + 1
	}
	s.mu.Unlock()
	writeJSON(w, http.StatusOK, in)
}

func (s *store) count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.users)
}

func (s *store) handler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	switch {
	case len(parts) == 1 && parts[0] == "users" && r.Method == http.MethodPost:
		s.createPOST(w, r)
	case len(parts) == 2 && parts[0] == "users" && r.Method == http.MethodPut:
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}
		s.putItem(w, r, id)
	default:
		http.NotFound(w, r)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func main() {
	st := newStore()
	srv := httptest.NewServer(http.HandlerFunc(st.handler))
	defer srv.Close()
	fmt.Printf("Сервер слушает на %s\n\n", srv.URL)

	client := &http.Client{Timeout: 5 * time.Second}

	// --- 1) PUT идемпотентен ---
	fmt.Println("1) PUT /users/1 дважды с одним телом:")
	u1 := putUser(client, srv.URL+"/users/1", `{"name":"Alice"}`)
	u2 := putUser(client, srv.URL+"/users/1", `{"name":"Alice"}`)
	fmt.Printf("   первый:  %+v\n   повтор:  %+v\n", u1, u2)
	fmt.Printf("   -> результат одинаковый: %v; всего пользователей: %d\n\n", u1 == u2, st.count())

	// --- 2) POST без ключа НЕ идемпотентен: два запроса -> дубликат ---
	fmt.Println("2) POST /users дважды с одним телом БЕЗ Idempotency-Key:")
	p1 := postUser(client, srv.URL+"/users", `{"name":"Bob"}`, "")
	p2 := postUser(client, srv.URL+"/users", `{"name":"Bob"}`, "")
	fmt.Printf("   первый:  %+v\n   повтор:  %+v\n", p1, p2)
	fmt.Printf("   -> созданы РАЗНЫЕ ID (дубликат): %v; всего пользователей: %d\n\n",
		p1.ID != p2.ID, st.count())

	// --- 3) POST с Idempotency-Key дедуплицируется ---
	fmt.Println("3) POST /users дважды с одним Idempotency-Key: k-123:")
	k1 := postUser(client, srv.URL+"/users", `{"name":"Carol"}`, "k-123")
	k2 := postUser(client, srv.URL+"/users", `{"name":"Carol"}`, "k-123")
	fmt.Printf("   первый:  %+v\n   повтор:  %+v\n", k1, k2)
	fmt.Printf("   -> тот же ID, дубликата нет: %v; всего пользователей: %d\n\n",
		k1.ID == k2.ID, st.count())

	fmt.Println("Готово, программа завершается.")
}

func putUser(c *http.Client, url, body string) User {
	req, _ := http.NewRequest(http.MethodPut, url, strings.NewReader(body))
	return send(c, req)
}

func postUser(c *http.Client, url, body, idemKey string) User {
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	if idemKey != "" {
		req.Header.Set("Idempotency-Key", idemKey)
	}
	return send(c, req)
}

func send(c *http.Client, req *http.Request) User {
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("%s %s: %v", req.Method, req.URL, err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	var u User
	if err := json.Unmarshal(b, &u); err != nil {
		log.Fatalf("decode: %v (body=%q)", err, b)
	}
	return u
}
