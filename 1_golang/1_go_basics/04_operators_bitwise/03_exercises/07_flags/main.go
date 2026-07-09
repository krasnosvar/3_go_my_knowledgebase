package main

import (
	"fmt"
)

const (
	FlagActive  = 1 << 0 // 1
	FlagPremium = 1 << 1 // 2
	FlagBanned  = 1 << 2 // 4
)

// setFlag устанавливает переданный флаг в маске flags и возвращает новую маску
func setFlag(flags uint8, flag uint8) uint8 {
	// TODO: установите флаг с помощью оператора побитового ИЛИ (|)
	return flags
}

// clearFlag сбрасывает переданный флаг в маске flags и возвращает новую маску
func clearFlag(flags uint8, flag uint8) uint8 {
	// TODO: сбросьте флаг с помощью оператора сброса битов И НЕ (&^)
	return flags
}

// hasFlag проверяет, установлен ли переданный флаг в маске flags
func hasFlag(flags uint8, flag uint8) bool {
	// TODO: проверьте наличие флага с помощью побитового И (&)
	return false
}

func main() {
	fmt.Println("--- Проверка 07_flags ---")

	var userFlags uint8 = 0

	// 1. Устанавливаем флаги
	userFlags = setFlag(userFlags, FlagActive)
	userFlags = setFlag(userFlags, FlagPremium)

	t1 := hasFlag(userFlags, FlagActive)  // true
	t2 := hasFlag(userFlags, FlagPremium) // true
	t3 := hasFlag(userFlags, FlagBanned)  // false

	// 2. Сбрасываем флаг
	userFlags = clearFlag(userFlags, FlagPremium)
	t4 := hasFlag(userFlags, FlagPremium) // false

	if t1 && t2 && !t3 && !t4 {
		fmt.Println("Тест управления флагами: PASS")
	} else {
		fmt.Printf("Тест управления флагами: FAIL (t1=%t, t2=%t, t3=%t, t4=%t)\n", t1, t2, t3, t4)
	}
}
