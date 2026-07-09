package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// printMessage выводит строку "Hello, Go!" с переносом строки в w
func printMessage(w io.Writer) {
	fmt.Fprintln(w, "Hello, Go!")
}

func main() {
	fmt.Println("--- Проверка 01_print (Решение) ---")

	var buf bytes.Buffer
	printMessage(&buf)

	output := buf.String()
	expected := "Hello, Go!\n"

	if output == expected {
		fmt.Println("Тест printMessage: PASS")
	} else {
		fmt.Printf("Тест printMessage: FAIL (ожидалось %q, получено %q)\n", expected, output)
	}

	// Также выводим в stdout
	printMessage(os.Stdout)
}
