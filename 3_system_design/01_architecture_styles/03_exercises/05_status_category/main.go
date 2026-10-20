// Упражнение 05: категория HTTP-статуса.
//
// Задача: по числовому коду HTTP-статуса вернуть его категорию:
//
//	1xx -> "informational"
//	2xx -> "success"
//	3xx -> "redirection"
//	4xx -> "client_error"
//	5xx -> "server_error"
//
// Для кода вне диапазона [100, 599] вернуть "unknown".
//
// Запуск: go run ./03_exercises/05_status_category/main.go
package main

import (
	"fmt"
	"net/http"
)

// Держим импорт net/http живым (в заготовке он ещё не используется).
// В решении удобно сверяться с константами вроде http.StatusOK.
var _ = http.StatusOK

// statusCategory возвращает категорию по коду HTTP-статуса.
func statusCategory(code int) string {
	// TODO: реализуйте классификацию кода по диапазонам сотен.
	return ""
}

func main() {
	fmt.Println("--- Проверка 05_status_category ---")

	cases := []struct {
		code int
		want string
	}{
		{100, "informational"},
		{200, "success"},
		{204, "success"},
		{301, "redirection"},
		{404, "client_error"},
		{429, "client_error"},
		{500, "server_error"},
		{503, "server_error"},
		{99, "unknown"},
		{600, "unknown"},
	}

	allOK := true
	for _, c := range cases {
		got := statusCategory(c.code)
		if got != c.want {
			allOK = false
			fmt.Printf("FAIL: statusCategory(%d) = %q, ожидалось %q\n", c.code, got, c.want)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
