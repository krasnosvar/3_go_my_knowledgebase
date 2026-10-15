package main

import (
	"fmt"
)

// ServerStatus представляет состояние сервера
type ServerStatus int

// Объявление статусов с помощью iota.
// Значение 0 зарезервировано под StatusUnknown: так нулевое значение
// (zero value) переменной типа ServerStatus означает «статус неизвестен».
const (
	StatusUnknown ServerStatus = iota // 0
	StatusActive                      // 1
	StatusPending                     // 2
	StatusOffline                     // 3
)

// Человекочитаемые имена статусов — тоже константы.
// (Позже, в теме методов, такую печать делают через метод String() — см. ../../../../2_oop/01_methods/.)
const (
	NameUnknown = "UNKNOWN"
	NameActive  = "ACTIVE"
	NamePending = "PENDING"
	NameOffline = "OFFLINE"
)

func main() {
	webStatus := StatusActive
	dbStatus := StatusPending
	testStatus := StatusOffline
	var newStatus ServerStatus // нулевое значение = StatusUnknown

	fmt.Println("--- Статус инфраструктуры ---")
	fmt.Printf("Сервер srv-web-01: %s (код: %d)\n", NameActive, webStatus)
	fmt.Printf("Сервер srv-db-01:  %s (код: %d)\n", NamePending, dbStatus)
	fmt.Printf("Сервер srv-test-0: %s (код: %d)\n", NameOffline, testStatus)
	fmt.Printf("Новый сервер:      %s (код: %d)\n", NameUnknown, newStatus)
}
