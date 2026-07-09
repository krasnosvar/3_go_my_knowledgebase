package main

import (
	"fmt"
)

type ResourceNode struct {
	Name string
	IP   string
}

func main() {
	node := &ResourceNode{
		Name: "node-k8s-01",
		IP:   "10.0.0.12",
	}

	cpuUsage := 0.785432          // 78.54%
	memoryPercentage := 42.129384 // 42.13%

	fmt.Println("--- Системные метрики инфраструктуры ---")

	// Форматирование вещественных чисел с заданной точностью %.2f%%
	fmt.Printf("Узел: %s (%s)\n", node.Name, node.IP)
	fmt.Printf("Загрузка CPU: %.2f%%\n", cpuUsage*100)
	fmt.Printf("Загрузка RAM: %.3f%%\n", memoryPercentage)

	// Форматирование адресов памяти с помощью %p
	fmt.Printf("Адрес объекта node в памяти (указатель): %p\n", node)
}
