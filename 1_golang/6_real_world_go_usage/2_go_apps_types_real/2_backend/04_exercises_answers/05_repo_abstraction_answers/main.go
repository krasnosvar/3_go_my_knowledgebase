// Упражнение 05: handler зависит от интерфейса хранилища (Решение).
package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
)

var errNotFound = errors.New("not found")

type store interface {
	Save(id, text string) error
	Find(id string) (string, error)
}

type memStore struct{ m map[string]string }

func (s *memStore) Save(id, text string) error {
	s.m[id] = text
	return nil
}

func (s *memStore) Find(id string) (string, error) {
	t, ok := s.m[id]
	if !ok {
		return "", errNotFound
	}
	return t, nil
}

type brokenStore struct{}

func (brokenStore) Save(string, string) error   { return errors.New("disk full") }
func (brokenStore) Find(string) (string, error) { return "", errors.New("conn reset") }

func getNote(s store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := s.Find(r.PathValue("id"))
		switch {
		case errors.Is(err, errNotFound):
			http.NotFound(w, r)
		case err != nil:
			http.Error(w, "internal error", http.StatusInternalServerError)
		default:
			fmt.Fprint(w, t)
		}
	}
}

func main() {
	fmt.Println("--- Проверка 05_repo_abstraction (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	get := func(s store, id string) *httptest.ResponseRecorder {
		mux := http.NewServeMux()
		mux.HandleFunc("GET /notes/{id}", getNote(s))
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest("GET", "/notes/"+id, nil))
		return rec
	}

	mem := &memStore{m: map[string]string{}}
	_ = mem.Save("a", "hello")
	rec := get(mem, "a")
	check(rec.Code == 200 && rec.Body.String() == "hello", "memStore: найденная заметка → 200 hello")
	check(get(mem, "zzz").Code == 404, "memStore: нет заметки → 404")
	check(get(brokenStore{}, "a").Code == 500, "brokenStore: ошибка хранилища → 500")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
