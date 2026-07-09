package main

import (
	"fmt"
)

// Глобальные константы на уровне пакета
const (
	AppName    = "config-loader"
	AppVersion = "1.0.4"
)

// Глобальная переменная для дефолтного хоста
var defaultHost = "localhost"

func main() {
	// 1. Использование var для явного типа при нулевой инициализации
	var port int // port == 0

	// 2. Использование := для локальной переменной с выводом типа
	verbose := true

	// Имитация чтения конфига: переопределяем порт
	port = 8080

	fmt.Printf("--- Запуск CLI %s (v%s) ---\n", AppName, AppVersion)
	fmt.Printf("Сервер запущен на %s:%d\n", defaultHost, port)
	fmt.Printf("Отладочный вывод активен: %t\n", verbose)

	// Пример области видимости блока
	if verbose {
		// Переменная tempConfig видна ТОЛЬКО внутри этого блока if
		tempConfig := "Loaded from /etc/app/config.json"
		fmt.Printf("Отладка: %s\n", tempConfig)
	}

	// fmt.Println(tempConfig) // Ошибка компиляции: undefined: tempConfig
}
