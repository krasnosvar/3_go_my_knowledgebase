// Упражнение 06: композиция нескольких компонентов.
//
// Задача: API собран из двух компонентов — Auth (проверка доступа) и
// *AuditLog (журнал). Реализуйте Handle(user):
//   - если a.Check(user) — результат "200 OK", иначе "403 Forbidden";
//   - в журнал записать строку "<user> -> <результат>" через a.Record(...).
//
// Методы Check и Record вызываются напрямую у API благодаря встраиванию.
//
// Запуск: go run ./03_exercises/06_composition/main.go
package main

import (
	"fmt"
)

// Auth знает разрешённых пользователей.
type Auth struct {
	allowed map[string]bool
}

func (a Auth) Check(user string) bool { return a.allowed[user] }

// AuditLog — журнал событий.
type AuditLog struct {
	lines []string
}

func (l *AuditLog) Record(s string) { l.lines = append(l.lines, s) }

// API — композиция Auth и журнала.
type API struct {
	Auth
	*AuditLog
}

// Handle обрабатывает запрос пользователя.
func (a API) Handle(user string) string {
	// TODO: выберите статус по a.Check(user), запишите a.Record(user + " -> " + status).
	return ""
}

func main() {
	fmt.Println("--- Проверка 06_composition ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	api := API{
		Auth:     Auth{allowed: map[string]bool{"den": true}},
		AuditLog: &AuditLog{},
	}
	check(api.Handle("den") == "200 OK", "den должен получить 200 OK")
	check(api.Handle("eve") == "403 Forbidden", "eve должна получить 403 Forbidden")
	want := "[den -> 200 OK eve -> 403 Forbidden]"
	check(fmt.Sprint(api.lines) == want, fmt.Sprintf("журнал: ожидалось %s, получено %v", want, api.lines))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
