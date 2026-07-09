package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// printMessage должен вывести строку "Hello, Go!" в переданный Writer
func printMessage(w io.Writer) {
	// TODO: Выведите строку "Hello, Go!" с переносом строки в w
	// (Подсказка: используйте fmt.Fprintln)
}

func main() {
	fmt.Println("--- Проверка 01_print ---")

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
