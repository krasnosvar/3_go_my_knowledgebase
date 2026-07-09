package main

import (
	"fmt"
)

type Metrics struct {
	CPUUsage float64
	RAMUsage float64
}

type Server struct {
	Name    string
	Metrics *Metrics // Может быть nil, если сервер недоступен
}

// Печать метрик сервера с безопасной проверкой на nil
func printServerMetrics(s *Server) {
	if s == nil {
		fmt.Println("[ERROR] Указатель на сервер равен nil!")
		return
	}

	fmt.Printf("Сервер: %s\n", s.Name)

	// ВНИМАНИЕ: Если s.Metrics равен nil, разыменование s.Metrics.CPUUsage приведет к panic!
	// Обязательно проверяем внутренний указатель на nil:
	if s.Metrics == nil {
		fmt.Println("  Метрики недоступны (узел отключен)")
		return
	}

	fmt.Printf("  CPU: %.2f%%, RAM: %.2f%%\n", s.Metrics.CPUUsage, s.Metrics.RAMUsage)
}

func main() {
	fmt.Println("--- Мониторинг серверов: nil-указатели ---")

	// 1. Активный сервер с метриками
	server1 := &Server{
		Name: "node-prod-01",
		Metrics: &Metrics{
			CPUUsage: 45.2,
			RAMUsage: 73.1,
		},
	}

	// 2. Отключенный сервер (метрики отсутствуют, т.е. nil)
	server2 := &Server{
		Name:    "node-offline-02",
		Metrics: nil,
	}

	fmt.Println("Вывод для server1:")
	printServerMetrics(server1)

	fmt.Println("\nВывод для server2 (безопасный обход nil):")
	printServerMetrics(server2)
}
