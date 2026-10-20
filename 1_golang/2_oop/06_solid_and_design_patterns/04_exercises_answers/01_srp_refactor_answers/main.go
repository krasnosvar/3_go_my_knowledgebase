// Упражнение 01: разделение ответственности (SRP) через интерфейс (Решение).
package main

import "fmt"

type Formatter interface {
	Format(amount int) string
}

type Calculator struct{}

func (Calculator) Total(items []int) int {
	sum := 0
	for _, v := range items {
		sum += v
	}
	return sum
}

type PlainFormatter struct{}

func (PlainFormatter) Format(amount int) string {
	return fmt.Sprintf("итого: %d", amount)
}

func main() {
	fmt.Println("--- Проверка 01_srp_refactor (Решение) ---")

	calc := Calculator{}
	var f Formatter = PlainFormatter{}

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
