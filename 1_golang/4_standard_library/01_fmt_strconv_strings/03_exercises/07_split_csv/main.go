// Упражнение 07: разбор простой строки с разделителем-запятой.
//
// Задача: реализуйте splitFields(line):
//   - делит строку по ',' (strings.Split);
//   - обрезает пробелы у каждого поля;
//   - пустая строка (или строка из пробелов) → пустой срез длины 0
//     (ловушка: strings.Split("", ",") возвращает [""] — срез из ОДНОГО элемента);
//   - пустые поля внутри строки сохраняются: "a,,b" → ["a" "" "b"].
//
// Для настоящего CSV (кавычки, экранирование) используйте encoding/csv —
// он разобран в теме 03_encoding_json_csv_base64.
//
// Запуск: go run ./03_exercises/07_split_csv/main.go
package main

import (
	"fmt"
	"strings"
)

var _ = strings.Split // держим импорт живым в заготовке

// splitFields делит строку по запятым и обрезает пробелы у полей.
func splitFields(line string) []string {
	// TODO: обработайте пустую строку, затем Split + TrimSpace для каждого поля.
	return nil
}

func main() {
	fmt.Println("--- Проверка 07_split_csv ---")

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
