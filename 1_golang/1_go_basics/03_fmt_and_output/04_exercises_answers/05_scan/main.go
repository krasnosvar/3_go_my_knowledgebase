package main

import (
	"fmt"
	"strings"
)

// parseConnectionInfo парсит строку "IP:PORT" с помощью fmt.Sscanf и strings.Replace
func parseConnectionInfo(input string) (string, int, error) {
	var ip string
	var port int

	// Заменяем двоеточие на пробел, чтобы Sscanf корректно разделил строки
	inputWithSpace := strings.Replace(input, ":", " ", 1)
	_, err := fmt.Sscanf(inputWithSpace, "%s %d", &ip, &port)
	if err != nil {
		return "", 0, err
	}

	return ip, port, nil
}

func main() {
	fmt.Println("--- Проверка 05_scan (Решение) ---")

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
