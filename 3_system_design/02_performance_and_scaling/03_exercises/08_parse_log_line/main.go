// Упражнение 08: разбор полей структурного лога (logfmt).
//
// Задача: реализовать parseLogFields — распарсить строку в формате
// logfmt (пары key=value через пробел) в map[string]string.
//
// Пример входа:
//
//	level=INFO msg=started port=8080 request_id=req-42
//
// Правила (упрощённые, значения без пробелов и без кавычек):
//   - разделитель пар — один или несколько пробелов;
//   - в каждой паре ровно один '=', слева ключ, справа значение;
//   - токен без '=' игнорируется;
//   - пустой ключ игнорируется.
//
// Реализуйте parseLogFields.
//
// Запуск: go run ./03_exercises/08_parse_log_line/main.go
package main

import (
	"fmt"
	"strings"
)

// Держим импорт strings живым (в решении удобны Fields/SplitN).
var _ = strings.Fields

// parseLogFields разбирает строку logfmt в map ключ->значение.
func parseLogFields(line string) map[string]string {
	// TODO:
	//  1) разбить строку на токены по пробелам (strings.Fields);
	//  2) для каждого токена: strings.SplitN(tok, "=", 2);
	//  3) если частей 2 и ключ непустой -> положить в map;
	//  4) вернуть map.
	return nil
}

func main() {
	fmt.Println("--- Проверка 08_parse_log_line ---")

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
	// Токен "broken" без '=' должен быть проигнорирован.
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
