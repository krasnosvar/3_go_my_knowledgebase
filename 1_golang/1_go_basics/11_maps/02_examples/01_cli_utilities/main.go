package main

import (
	"fmt"
	"os"
)

/*
ПРИМЕР 1: CLI утилиты (Паттерн Command Registry)

В терминальных приложениях (CLI) часто используется структура мапы,
которая связывает введенное пользователем слово (команду) с функцией-обработчиком.
Это заменяет огромные "простыни" if-else блоков на элегантный поиск за O(1).
*/

// Тип принимаемой функции
type CommandHandler func(args []string)

func main() {
	// Реестр команд
	registry := map[string]CommandHandler{
		"version": handleVersion,
		"help":    handleHelp,
		"run":     handleRun,
	}

	// Эмулируем запуск из консоли (обычно берется os.Args[1:])
	inputArgs := []string{"run", "nginx", "alpine"}
	if len(os.Args) > 1 {
		inputArgs = os.Args[1:]
	}

	cmdName := inputArgs[0]
	cmdArgs := inputArgs[1:]

	// Ищем команду в реестре
	handler, exists := registry[cmdName]
	if !exists {
		fmt.Printf("Error: unknown command '%s'\n", cmdName)
		registry["help"](nil)
		return
	}

	// Делегируем выполнение нужной функции
	handler(cmdArgs)
}

func handleVersion(args []string) {
	fmt.Println("App Version: v1.0.0")
}

func handleHelp(args []string) {
	fmt.Println("Available commands: version, help, run")
}

func handleRun(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: 'run' requires an image name")
		return
	}
	fmt.Printf("Starting container from image: %s\n", args[0])
}
