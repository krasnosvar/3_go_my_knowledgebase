// Упражнение 04: разбор полей структурного лога (logfmt) (Решение).
package main

import (
	"fmt"
	"strings"
)

func parseLogFields(line string) map[string]string {
	out := make(map[string]string)
	for _, tok := range strings.Fields(line) {
		parts := strings.SplitN(tok, "=", 2)
		if len(parts) != 2 || parts[0] == "" {
			continue
		}
		out[parts[0]] = parts[1]
	}
	return out
}

func main() {
	fmt.Println("--- Проверка 04_parse_log_line (Решение) ---")

	line := "level=INFO   msg=started port=8080 request_id=req-42 broken"
	got := parseLogFields(line)

	want := map[string]string{
		"level":      "INFO",
		"msg":        "started",
		"port":       "8080",
		"request_id": "req-42",
	}

	allOK := true
	if len(got) != len(want) {
		allOK = false
		fmt.Printf("FAIL: получено %d полей, ожидалось %d: %v\n", len(got), len(want), got)
	}
	for k, v := range want {
		if got[k] != v {
			allOK = false
			fmt.Printf("FAIL: поле %q = %q, ожидалось %q\n", k, got[k], v)
		}
	}
	if _, ok := got["broken"]; ok {
		allOK = false
		fmt.Println("FAIL: токен без '=' не должен попадать в map")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
