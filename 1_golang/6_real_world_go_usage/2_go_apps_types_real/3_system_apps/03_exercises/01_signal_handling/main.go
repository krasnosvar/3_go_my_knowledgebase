// Упражнение 01: обработка сигналов демоном.
//
// В настоящем демоне канал наполняет signal.Notify(sigs, syscall.SIGHUP,
// syscall.SIGTERM, os.Interrupt). Логику обработки выносим в функцию,
// которую можно проверить, подавая сигналы в канал вручную.
//
// Задача: реализуйте loop(sigs <-chan os.Signal, reload func()) string:
//   - SIGHUP → вызвать reload() и продолжать ждать;
//   - SIGTERM или os.Interrupt → вернуть "stopped by <sig>" (sig.String());
//   - канал закрыт → вернуть "channel closed".
//
// Запуск: go run ./03_exercises/01_signal_handling
package main

import (
	"fmt"
	"os"
	"syscall"
)

func loop(sigs <-chan os.Signal, reload func()) string {
	// TODO: for sig := range sigs { switch sig { ... } }.
	return ""
}

func main() {
	fmt.Println("--- Проверка 01_signal_handling ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	sigs := make(chan os.Signal, 4)
	reloads := 0
	sigs <- syscall.SIGHUP
	sigs <- syscall.SIGHUP
	sigs <- syscall.SIGTERM
	got := loop(sigs, func() { reloads++ })
	check(reloads == 2, "два SIGHUP → два reload")
	check(got == "stopped by "+syscall.SIGTERM.String(), "SIGTERM останавливает цикл: "+got)

	sigs2 := make(chan os.Signal)
	close(sigs2)
	check(loop(sigs2, func() {}) == "channel closed", "закрытый канал → channel closed")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
