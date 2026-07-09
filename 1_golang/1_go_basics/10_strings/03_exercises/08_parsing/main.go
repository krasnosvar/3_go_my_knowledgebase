package main

import (
	"fmt"
	"strings"
)

var _ = strings.Contains // Держим импорт активным для компиляции заготовки

// parseKeyValue разбирает строку формата "key=value".
// Возвращает очищенные от пробелов ключ и значение, и флаг успешности разбора (bool).
// Если символ '=' отсутствует или ключ пустой после очистки, верните "", "", false.
func parseKeyValue(line string) (string, string, bool) {
	// TODO: Реализуйте разбор.
	return "", "", false
}

func main() {
	fmt.Println("--- Проверка 08_parsing ---")

	k1, v1, ok1 := parseKeyValue("env = prod")
	k2, v2, ok2 := parseKeyValue("invalid-line")
	k3, v3, ok3 := parseKeyValue(" = value")

	t1 := ok1 && k1 == "env" && v1 == "prod"
	t2 := !ok2 && k2 == "" && v2 == ""
	t3 := !ok3 && k3 == "" && v3 == ""

	if t1 && t2 && t3 {
		fmt.Println("Тест parseKeyValue: PASS")
	} else {
		fmt.Printf("Тест parseKeyValue: FAIL (ok1=%t, k1=%q, v1=%q, ok2=%t, ok3=%t)\n", ok1, k1, v1, ok2, ok3)
	}
}
