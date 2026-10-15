// Упражнение 07: разбор простой строки с разделителем-запятой (Решение).
package main

import (
	"fmt"
	"strings"
)

// splitFields делит строку по запятым и обрезает пробелы у полей.
func splitFields(line string) []string {
	if strings.TrimSpace(line) == "" {
		return []string{}
	}
	parts := strings.Split(line, ",")
	for i, p := range parts {
		parts[i] = strings.TrimSpace(p)
	}
	return parts
}

func main() {
	fmt.Println("--- Проверка 07_split_csv (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := splitFields(" id, name ,  email ")
	check(len(got) == 3 && got[0] == "id" && got[1] == "name" && got[2] == "email",
		fmt.Sprintf("ожидалось [id name email], получено %q", got))

	got = splitFields("a,,b")
	check(len(got) == 3 && got[1] == "", fmt.Sprintf("пустые поля внутри сохраняются: получено %q", got))

	got = splitFields("   ")
	check(got != nil && len(got) == 0, fmt.Sprintf("пустая строка → пустой (не nil) срез, получено %q (nil=%t)", got, got == nil))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
