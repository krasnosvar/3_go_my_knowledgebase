// Упражнение 04: запись CSV через csv.Writer (Решение).
package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

// writeCSV сериализует строки в CSV.
func writeCSV(rows [][]string) (string, error) {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := w.WriteAll(rows); err != nil { // WriteAll сам делает Flush
		return "", fmt.Errorf("write csv: %w", err)
	}
	return buf.String(), nil
}

func main() {
	fmt.Println("--- Проверка 04_csv_write (Решение) ---")

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
