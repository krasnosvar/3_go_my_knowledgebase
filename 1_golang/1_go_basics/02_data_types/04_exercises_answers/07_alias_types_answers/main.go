package main

import (
	"fmt"
)

// 1. Объявите ИМЕНОВАННЫЙ тип (Named Type) UserID на базе int64.
type UserID int64

// 2. Объявите ПСЕВДОНИМ типа (Type Alias) с именем IDAlias для типа int64.
type IDAlias = int64

func main() {
	fmt.Println("--- Проверка 07_alias_types (Решение) ---")

	var uid UserID = 42
	var alias IDAlias = 42

	// %T печатает имя типа значения.
	// Именованный тип — это НОВЫЙ тип: печатается как main.UserID.
	// Псевдоним — это то же самое имя для int64: печатается как int64.
	uidType := fmt.Sprintf("%T", uid)
	aliasType := fmt.Sprintf("%T", alias)

	if uidType == "main.UserID" {
		fmt.Println("Именованный тип UserID: PASS")
	} else {
		fmt.Printf("Именованный тип UserID: FAIL (тип %s, ожидался main.UserID)\n", uidType)
	}

	if aliasType == "int64" {
		fmt.Println("Псевдоним типа IDAlias: PASS")
	} else {
		fmt.Printf("Псевдоним типа IDAlias: FAIL (тип %s, ожидался int64)\n", aliasType)
	}
}
