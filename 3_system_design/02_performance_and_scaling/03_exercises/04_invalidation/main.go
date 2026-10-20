// Упражнение 04: инвалидация кэша по ключу (cache-aside на запись).
//
// Задача: реализуйте кэш поверх «БД» (map) с инвалидацией на записи.
//   - Get(key): при попадании вернуть значение из кэша, счётчик обращений
//     к «БД» НЕ увеличивать; при промахе — прочитать из db, положить в кэш,
//     увеличить dbReads, вернуть значение;
//   - Update(key, val): записать в db и УДАЛИТЬ ключ из кэша
//     (инвалидация), чтобы следующее чтение взяло свежее значение.
//
// Смысл: после Update следующий Get обязан снова сходить в «БД»
// (dbReads увеличится), а не отдать устаревшее из кэша.
//
// Запуск: go run ./03_exercises/04_invalidation/main.go
package main

import "fmt"

type store struct {
	db      map[string]string // источник истины
	cache   map[string]string // кэш
	dbReads int               // сколько раз реально читали из db
}

func newStore() *store {
	return &store{
		db:    map[string]string{"k": "v1"},
		cache: make(map[string]string),
	}
}

// Get реализует cache-aside чтение.
func (s *store) Get(key string) (string, bool) {
	// TODO:
	//  - если key есть в cache -> вернуть его (без обращения к db);
	//  - иначе прочитать из db; если нет в db -> (\"\", false);
	//    если есть -> dbReads++, положить в cache, вернуть значение.
	return "", false
}

// Update пишет в db и инвалидирует ключ в кэше (удаляет его).
func (s *store) Update(key, val string) {
	// TODO: записать в db и удалить key из cache.
}

func main() {
	fmt.Println("--- Проверка 04_invalidation ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	s := newStore()

	if v, ok := s.Get("k"); !ok || v != "v1" { // miss -> из db
		check(false, "первый Get должен вернуть v1 из db")
	}
	check(s.dbReads == 1, fmt.Sprintf("после первого Get dbReads=1, получено %d", s.dbReads))

	if v, ok := s.Get("k"); !ok || v != "v1" { // hit -> из кэша
		check(false, "повторный Get должен вернуть v1 из кэша")
	}
	check(s.dbReads == 1, "повторный Get не должен читать db (должно быть попадание)")

	s.Update("k", "v2") // запись + инвалидация

	if v, ok := s.Get("k"); !ok || v != "v2" { // снова miss -> свежее из db
		check(false, "после Update Get должен вернуть свежее v2")
	}
	check(s.dbReads == 2, fmt.Sprintf("после инвалидации Get должен снова читать db, dbReads=2, получено %d", s.dbReads))

	if _, ok := s.Get("nope"); ok {
		check(false, "несуществующий ключ не должен находиться")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
