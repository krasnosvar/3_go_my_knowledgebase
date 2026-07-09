package main

import (
	"fmt"
)

type Cfg struct {
	Enabled bool
	Count   int
	Name    string
}

// isZeroValue проверяет, является ли переданная структура Cfg полностью нулевой
// (то есть все её поля имеют дефолтные нулевые значения своих типов).
func isZeroValue(cfg Cfg) bool {
	// TODO: реализуйте функцию
	return false
}

func main() {
	fmt.Println("--- Проверка 06_zero_values ---")

	tests := []struct {
		input Cfg
		want  bool
	}{
		{Cfg{Enabled: false, Count: 0, Name: ""}, true},
		{Cfg{Enabled: true, Count: 0, Name: ""}, false},
		{Cfg{Enabled: false, Count: 5, Name: ""}, false},
		{Cfg{Enabled: false, Count: 0, Name: "Admin"}, false},
	}

	allPassed := true
	for i, tt := range tests {
		got := isZeroValue(tt.input)
		if got == tt.want {
			fmt.Printf("Тест %d (Enabled=%t, Count=%d, Name=%q): PASS\n", i+1, tt.input.Enabled, tt.input.Count, tt.input.Name)
		} else {
			fmt.Printf("Тест %d (Enabled=%t, Count=%d, Name=%q): FAIL (ожидалось %t, получено %t)\n",
				i+1, tt.input.Enabled, tt.input.Count, tt.input.Name, tt.want, got)
			allPassed = false
		}
	}

	if allPassed {
		fmt.Println("Все тесты пройдены успешно!")
	}
}
