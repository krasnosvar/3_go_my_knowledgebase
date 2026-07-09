package main

import (
	"fmt"
)

// makeLogMessage возвращает отформатированную строку вида "[INFO] host=<host> status=<status>"
func makeLogMessage(host string, status int) string {
	// TODO: Сгенерируйте и верните отформатированную строку с помощью fmt.Sprintf
	return ""
}

func main() {
	fmt.Println("--- Проверка 03_sprintf ---")

	msg := makeLogMessage("127.0.0.1", 200)
	expected := "[INFO] host=127.0.0.1 status=200"

	if msg == expected {
		fmt.Println("Тест makeLogMessage: PASS")
	} else {
		fmt.Printf("Тест makeLogMessage: FAIL (ожидалось %q, получено %q)\n", expected, msg)
	}
}
