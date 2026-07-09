package main

import (
	"fmt"
	"strings"
)

var _ = strings.Contains // Держим импорт активным для компиляции заготовки

// normalizeTag обрезает пробелы по краям строки s, переводит ее в нижний регистр,
// и заменяет любые последовательности пробельных символов на один дефис '-'.
// Пример: "  Go   Backend  " -> "go-backend"
func normalizeTag(s string) string {
	// TODO: Реализуйте нормализацию.
	return ""
}

func main() {
	fmt.Println("--- Проверка 07_normalization ---")

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
