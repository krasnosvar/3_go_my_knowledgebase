package main

import (
	"fmt"
)

// ServerStatus представляет состояние сервера
type ServerStatus int

// Объявление статусов с помощью iota.
// Начинаем с 1, пропуская 0 (Unknown) с помощью заглушки `_`.
const (
	StatusUnknown ServerStatus = iota // 0
	StatusActive                      // 1
	StatusPending                     // 2
	StatusOffline                     // 3
)

// Метод String позволяет красиво печатать статус
func (s ServerStatus) String() string {
	switch s {
	case StatusActive:
		return "ACTIVE"
	case StatusPending:
		return "PENDING"
	case StatusOffline:
		return "OFFLINE"
	default:
		return "UNKNOWN"
	}
}

func main() {
	// Создаем карту серверов и их состояний
	servers := map[string]ServerStatus{
		"srv-web-01": StatusActive,
		"srv-db-01":  StatusPending,
		"srv-test-0": StatusOffline,
	}

	fmt.Println("--- Статус инфраструктуры ---")
	for name, status := range servers {
		fmt.Printf("Сервер %s: %s (код: %d)\n", name, status, status)
	}
}
