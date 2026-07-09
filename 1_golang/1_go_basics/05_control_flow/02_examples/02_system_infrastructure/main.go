package main

import (
	"fmt"
	"strings"
)

func main() {
	// База логов серверов: слайс файлов логов, содержащих строки лога
	logs := [][]string{
		{"[INFO] Server started", "[DEBUG] Connecting to db", "[ERROR] Connection failed"},
		{"[INFO] Healthy node", "[WARNING] Slow disk read"},
		{"[CRITICAL] Out of memory", "[INFO] Shutting down"},
	}

	fmt.Println("--- Системный поиск: Использование меток (Labels) ---")

	// Цель: Найти первую критическую ошибку (ERROR или CRITICAL), но пропускать логи серверов,
	// которые содержат пометку "[WARNING] Slow disk read" (имитация пропуска нестабильных узлов).

OuterLoop:
	for serverIdx, serverLogs := range logs {
		fmt.Printf("Анализируем сервер #%d...\n", serverIdx)

		for _, logLine := range serverLogs {
			// Если сервер сообщает об аппаратной проблеме диска, пропускаем этот сервер целиком
			if strings.Contains(logLine, "Slow disk read") {
				fmt.Println("  -> Пропускаем сервер из-за проблем с диском")
				continue OuterLoop
			}

			// Если нашли серьезную ошибку, выходим из всех циклов
			if strings.Contains(logLine, "ERROR") || strings.Contains(logLine, "CRITICAL") {
				fmt.Printf("  [НАЙДЕНО] Ошибка: %s\n", logLine)
				break OuterLoop
			}

			fmt.Printf("  Проверен лог: %s\n", logLine)
		}
	}
	fmt.Println("Поиск завершен.")
}
