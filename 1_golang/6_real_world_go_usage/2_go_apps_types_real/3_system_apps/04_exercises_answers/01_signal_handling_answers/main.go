// Упражнение 01: обработка сигналов демоном (Решение).
package main

import (
	"fmt"
	"os"
	"syscall"
)

func loop(sigs <-chan os.Signal, reload func()) string {
	for sig := range sigs {
		switch sig {
		case syscall.SIGHUP:
			reload()
		case syscall.SIGTERM, os.Interrupt:
			return "stopped by " + sig.String()
		}
	}
	return "channel closed"
}

func main() {
	fmt.Println("--- Проверка 01_signal_handling (Решение) ---")

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
