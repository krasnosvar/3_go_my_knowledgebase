// Упражнение 04: запись CSV через csv.Writer.
//
// Задача: реализуйте writeCSV(rows), которая записывает строки в CSV
// и возвращает результат строкой.
//   - пишите в bytes.Buffer через csv.NewWriter;
//   - не забудьте w.Flush() и проверку w.Error() — без Flush данные
//     останутся в буфере writer'а;
//   - экранирование кавычек/запятых делает csv.Writer, вручную не нужно.
//
// Запуск: go run ./03_exercises/04_csv_write/main.go
package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

var _ = csv.NewWriter // держим импорты живыми в заготовке
var _ bytes.Buffer

// writeCSV сериализует строки в CSV.
func writeCSV(rows [][]string) (string, error) {
	// TODO: bytes.Buffer + csv.NewWriter + WriteAll (или Write + Flush + Error).
	return "", nil
}

func main() {
	fmt.Println("--- Проверка 04_csv_write ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got, err := writeCSV([][]string{
		{"name", "comment"},
		{"Ann", "hello, world"},
		{"Bob", `say "hi"`},
	})
	want := "name,comment\nAnn,\"hello, world\"\nBob,\"say \"\"hi\"\"\"\n"
	check(err == nil && got == want, fmt.Sprintf("получено %q, ожидалось %q (err=%v)", got, want, err))

	got, err = writeCSV(nil)
	check(err == nil && got == "", fmt.Sprintf("пустой ввод: получено %q, %v", got, err))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
