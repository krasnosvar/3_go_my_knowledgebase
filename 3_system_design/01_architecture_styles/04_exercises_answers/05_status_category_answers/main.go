// Упражнение 05: категория HTTP-статуса (Решение).
package main

import (
	"fmt"
	"net/http"
)

var _ = http.StatusOK

func statusCategory(code int) string {
	if code < 100 || code > 599 {
		return "unknown"
	}
	switch code / 100 {
	case 1:
		return "informational"
	case 2:
		return "success"
	case 3:
		return "redirection"
	case 4:
		return "client_error"
	case 5:
		return "server_error"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println("--- Проверка 05_status_category (Решение) ---")

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
