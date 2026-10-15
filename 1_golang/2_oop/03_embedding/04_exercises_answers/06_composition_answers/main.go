// Упражнение 06: композиция нескольких компонентов (Решение).
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

func (a API) Handle(user string) string {
	status := "403 Forbidden"
	if a.Check(user) {
		status = "200 OK"
	}
	a.Record(user + " -> " + status)
	return status
}

func main() {
	fmt.Println("--- Проверка 06_composition (Решение) ---")

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
