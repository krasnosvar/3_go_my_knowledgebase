package main

import (
	"fmt"
	"strings"
)

// parseConnectionInfo парсит строку вида "IP:PORT" (например "127.0.0.1:8080")
// и возвращает отдельно IP (string) и PORT (int).
var _ = strings.Replace

func parseConnectionInfo(input string) (string, int, error) {
	var ip string
	var port int

	// TODO: Используйте fmt.Sscanf для разбора строки input в переменные ip и port.
	// Подсказка: так как Sscanf не поддерживает character-classes вроде %[^:],
	// сначала замените двоеточие на пробел с помощью strings.Replace, а затем разберите шаблон "%s %d".

	return ip, port, nil
}

func main() {
	fmt.Println("--- Проверка 05_scan ---")

	ip, port, err := parseConnectionInfo("192.168.1.1:9000")
	if err != nil {
		fmt.Printf("Тест parseConnectionInfo: FAIL (ошибка: %v)\n", err)
		return
	}

	if ip == "192.168.1.1" && port == 9000 {
		fmt.Println("Тест parseConnectionInfo: PASS")
	} else {
		fmt.Printf("Тест parseConnectionInfo: FAIL (получено ip=%q, port=%d)\n", ip, port)
	}
}
