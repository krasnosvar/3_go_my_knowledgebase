package main

import (
	"fmt"
	"strings"
)

func normalizeTag(s string) string {
	// 1. Приведение к нижнему регистру и trim пробелов
	trimmed := strings.TrimSpace(strings.ToLower(s))

	// 2. Разделение по пробельным символам и соединение через '-'
	fields := strings.Fields(trimmed)
	return strings.Join(fields, "-")
}

func main() {
	fmt.Println("--- Проверка 07_normalization (Решение) ---")

	res1 := normalizeTag("  Go   Backend  ")
	res2 := normalizeTag("\tCLI\nTool  ")

	t1 := res1 == "go-backend"
	t2 := res2 == "cli-tool"

	if t1 && t2 {
		fmt.Println("Тест normalizeTag: PASS")
	} else {
		fmt.Printf("Тест normalizeTag: FAIL (res1=%q, res2=%q)\n", res1, res2)
	}
}
