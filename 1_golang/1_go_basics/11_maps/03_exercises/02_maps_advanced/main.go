package main

import (
	"fmt"
)

type Employee struct {
	Name   string
	Salary int
}

// updateSalary увеличивает зарплату сотрудника на указанную величину delta.
// Если сотрудника нет в базе, ничего делать не нужно.
// Помните: значения в мапе неадресуемы!
func updateSalary(db map[string]Employee, name string, delta int) {
	// TODO: Обновите зарплату сотрудника в мапе.
}

func main() {
	fmt.Println("--- Проверка 02_maps_advanced ---")

	db := map[string]Employee{
		"Alice": {Name: "Alice", Salary: 5000},
		"Bob":   {Name: "Bob", Salary: 6000},
	}

	updateSalary(db, "Alice", 1500)
	updateSalary(db, "Charlie", 1000) // Не должен вызывать панику

	emp, ok := db["Alice"]
	t1 := ok && emp.Salary == 6500

	if t1 {
		fmt.Println("Тест updateSalary: PASS")
	} else {
		fmt.Printf("Тест updateSalary: FAIL (Salary Alice = %d)\n", emp.Salary)
	}
}
