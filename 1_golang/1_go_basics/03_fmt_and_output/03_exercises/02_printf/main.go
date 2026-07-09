package main

import (
	"bytes"
	"fmt"
	"io"
)

// formatUser должен вывести в w отформатированную строку: "User: <name>, Age: <age>"
func formatUser(w io.Writer, name string, age int) {
	// TODO: Напишите форматированный вывод с помощью fmt.Fprintf
}

func main() {
	fmt.Println("--- Проверка 02_printf ---")

	var buf bytes.Buffer
	formatUser(&buf, "Alice", 30)

	output := buf.String()
	expected := "User: Alice, Age: 30"

	if output == expected {
		fmt.Println("Тест formatUser: PASS")
	} else {
		fmt.Printf("Тест formatUser: FAIL (ожидалось %q, получено %q)\n", expected, output)
	}
}
