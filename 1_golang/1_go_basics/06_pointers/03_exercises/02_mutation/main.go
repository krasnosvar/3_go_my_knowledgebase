package main

import (
	"fmt"
)

type Employee struct {
	Name   string
	Salary int
}

// giveBonus увеличивает зарплату сотрудника на величину bonus
func giveBonus(e *Employee, bonus int) {
	// TODO: Увеличьте e.Salary на bonus.
}

func main() {
	fmt.Println("--- Проверка 02_mutation ---")

	emp := Employee{Name: "Bob", Salary: 5000}
	giveBonus(&emp, 1000)

	if emp.Salary == 6000 {
		fmt.Println("Тест giveBonus: PASS")
	} else {
		fmt.Printf("Тест giveBonus: FAIL (ожидалось 6000, получено %d)\n", emp.Salary)
	}
}
