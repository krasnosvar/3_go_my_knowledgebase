package main

import (
	"fmt"
)

const (
	FlagActive  = 1 << 0
	FlagPremium = 1 << 1
	FlagBanned  = 1 << 2
)

// setFlag устанавливает переданный флаг
func setFlag(flags uint8, flag uint8) uint8 {
	return flags | flag
}

// clearFlag сбрасывает переданный флаг с помощью &^
func clearFlag(flags uint8, flag uint8) uint8 {
	return flags &^ flag
}

// hasFlag проверяет, установлен ли переданный флаг
func hasFlag(flags uint8, flag uint8) bool {
	return (flags & flag) != 0
}

func main() {
	fmt.Println("--- Проверка 07_flags (Решение) ---")

	var userFlags uint8 = 0

	userFlags = setFlag(userFlags, FlagActive)
	userFlags = setFlag(userFlags, FlagPremium)

	t1 := hasFlag(userFlags, FlagActive)
	t2 := hasFlag(userFlags, FlagPremium)
	t3 := hasFlag(userFlags, FlagBanned)

	userFlags = clearFlag(userFlags, FlagPremium)
	t4 := hasFlag(userFlags, FlagPremium)

	if t1 && t2 && !t3 && !t4 {
		fmt.Println("Тест управления флагами: PASS")
	} else {
		fmt.Printf("Тест управления флагами: FAIL (t1=%t, t2=%t, t3=%t, t4=%t)\n", t1, t2, t3, t4)
	}
}
