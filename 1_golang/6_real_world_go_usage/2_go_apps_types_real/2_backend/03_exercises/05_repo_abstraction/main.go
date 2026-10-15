// Упражнение 05: handler зависит от интерфейса хранилища.
//
// Задача:
//   - реализуйте memStore (Save/Find) — in-memory реализацию интерфейса store;
//     Find отсутствующего ключа → errNotFound;
//   - реализуйте handler getNote(s store): GET /notes/{id} → 200 и текст,
//     errNotFound → 404, другая ошибка → 500.
//
// В тесте handler проверяется и с memStore, и с «сломанным» brokenStore —
// подмена реализации без изменения handler'а.
//
// Запуск: go run ./03_exercises/05_repo_abstraction
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
	// TODO: сохраните в map.
	return nil
}

func (s *memStore) Find(id string) (string, error) {
	// TODO: верните errNotFound для отсутствующего id.
	return "", nil
}

type brokenStore struct{}

func (brokenStore) Save(string, string) error   { return errors.New("disk full") }
func (brokenStore) Find(string) (string, error) { return "", errors.New("conn reset") }

func getNote(s store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: s.Find(r.PathValue("id")) и маппинг ошибок в статусы.
	}
}

func main() {
	fmt.Println("--- Проверка 05_repo_abstraction ---")

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
