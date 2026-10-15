package main

import (
	"fmt"
	"strings"
)

func parseKeyValue(line string) (string, string, bool) {
	idx := strings.IndexByte(line, '=')
	if idx == -1 {
		return "", "", false
	}
	key := strings.TrimSpace(line[:idx])
	val := strings.TrimSpace(line[idx+1:])
	if key == "" {
		return "", "", false
	}
	return key, val, true
}

func main() {
	fmt.Println("--- Проверка 08_parsing (Решение) ---")

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
