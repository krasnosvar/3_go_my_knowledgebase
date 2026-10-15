package main

import (
	"fmt"
)

type Employee struct {
	Name   string
	Salary int
}

func updateSalary(db map[string]Employee, name string, delta int) {
	emp, ok := db[name]
	if !ok {
		return
	}
	emp.Salary += delta
	db[name] = emp
}

func main() {
	fmt.Println("--- Проверка 02_maps_advanced (Решение) ---")

	db := map[string]Employee{
		"Alice": {Name: "Alice", Salary: 5000},
		"Bob":   {Name: "Bob", Salary: 6000},
	}

	updateSalary(db, "Alice", 1500)
	updateSalary(db, "Charlie", 1000)

	emp, ok := db["Alice"]
	t1 := ok && emp.Salary == 6500

	if t1 {
		fmt.Println("Тест updateSalary: PASS")
	} else {
		fmt.Printf("Тест updateSalary: FAIL (Salary Alice = %d)\n", emp.Salary)
	}
}
