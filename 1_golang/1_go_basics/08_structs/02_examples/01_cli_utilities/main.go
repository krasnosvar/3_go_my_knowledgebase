package main

import (
	"fmt"
)

type BaseConfig struct {
	LogLevel string
	Version  string
}

func (c BaseConfig) PrintVersion() {
	fmt.Printf("[Config] Текущая версия: %s\n", c.Version)
}

// ServerConfig расширяет BaseConfig с помощью встраивания (Embedding)
type ServerConfig struct {
	BaseConfig // Встроенное анонимное поле
	Port       int
	Version    string // Переопределяет (затеняет) BaseConfig.Version
}

func main() {
	cfg := ServerConfig{
		BaseConfig: BaseConfig{
			LogLevel: "DEBUG",
			Version:  "v1.0.0-base",
		},
		Port:    8080,
		Version: "v1.2.0-server", // Затеняющая строка
	}

	// 1. Прямой доступ к полям встроенной структуры
	fmt.Println("LogLevel:", cfg.LogLevel) // Проксируется на cfg.BaseConfig.LogLevel

	// 2. Демонстрация затенения (shadowing)
	fmt.Println("cfg.Version (Server):", cfg.Version)
	fmt.Println("cfg.BaseConfig.Version:", cfg.BaseConfig.Version)

	// 3. Вызов унаследованного метода (будет использовать BaseConfig.Version!)
	cfg.PrintVersion()
}
