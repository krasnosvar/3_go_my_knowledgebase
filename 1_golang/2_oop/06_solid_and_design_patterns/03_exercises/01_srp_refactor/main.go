// Упражнение 01: разделение ответственности (SRP) через интерфейс.
//
// Задача: расчёт итоговой суммы и её форматирование должны быть РАЗНЫМИ
// ответственностями (у каждой своя причина для изменения):
//
//   - Calculator.Total(items) — считает сумму.
//   - Formatter.Format(amount) — превращает сумму в строку.
//
// Реализуйте:
//   - метод Calculator.Total (сумма элементов),
//   - метод PlainFormatter.Format так, чтобы Format(120) == "итого: 120".
//
// Заготовка компилируется и в стартовом состоянии печатает FAIL —
// доведите её до PASS.
//
// Запуск: go run ./03_exercises/01_srp_refactor/main.go
package main

import "fmt"

// Formatter — ответственность за представление суммы в виде строки.
type Formatter interface {
	Format(amount int) string
}

// Calculator — ответственность за подсчёт суммы.
type Calculator struct{}

// Total возвращает сумму всех элементов.
func (Calculator) Total(items []int) int {
	// TODO: просуммируйте items и верните результат.
	return 0
}

// PlainFormatter форматирует сумму как "итого: <amount>".
type PlainFormatter struct{}

// Format должен вернуть строку вида "итого: 120".
func (PlainFormatter) Format(amount int) string {
	// TODO: верните строку "итого: <amount>" (используйте fmt.Sprintf).
	return ""
}

func main() {
	fmt.Println("--- Проверка 01_srp_refactor ---")

	calc := Calculator{}
	var f Formatter = PlainFormatter{} // держит интерфейс "живым"

	allOK := true

	if got := calc.Total([]int{10, 20, 30}); got != 60 {
		allOK = false
		fmt.Printf("FAIL: Total = %d, ожидалось 60\n", got)
	}
	if got := calc.Total(nil); got != 0 {
		allOK = false
		fmt.Printf("FAIL: Total(nil) = %d, ожидалось 0\n", got)
	}
	if got := f.Format(120); got != "итого: 120" {
		allOK = false
		fmt.Printf("FAIL: Format(120) = %q, ожидалось %q\n", got, "итого: 120")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
