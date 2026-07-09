package main

import (
	"fmt"
	"reflect"
)

// 1. Объявите ИМЕНОВАННЫЙ тип (Named Type) UserID на базе int64.
// TODO: type UserID ...

// 2. Объявите ПСЕВДОНИМ типа (Type Alias) с именем IDAlias для типа int64.
// TODO: type IDAlias = ...

var _ = reflect.TypeOf

// Функция проверки совместимости типов для автопроверщика
func checkTypes() (bool, bool) {
	// ВНИМАНИЕ: Для прохождения проверки раскомментируйте код ниже после объявления типов

	/*
		var uid UserID = 42
		var alias IDAlias = 42
		var raw int64 = 42

		// Проверяем, что IDAlias совместим с int64 без явного приведения
		alias = raw
		_ = alias

		// Проверяем, что UserID требует явного приведения (нельзя скомпилировать uid = raw без int64(uid) или UserID(raw))
		uidType := reflect.TypeOf(uid)
		aliasType := reflect.TypeOf(alias)

		isNamedType := uidType.Name() == "UserID" && uidType.Kind() == reflect.Int64
		isAlias := aliasType.Name() == "int64" && aliasType.Kind() == reflect.Int64

		return isNamedType, isAlias
	*/

	return false, false
}

func main() {
	fmt.Println("--- Проверка 07_alias_types ---")

	isNamed, isAlias := checkTypes()

	if isNamed {
		fmt.Println("Именованный тип UserID: PASS")
	} else {
		fmt.Println("Именованный тип UserID: FAIL (или код проверки закомментирован)")
	}

	if isAlias {
		fmt.Println("Псевдоним типа IDAlias: PASS")
	} else {
		fmt.Println("Псевдоним типа IDAlias: FAIL (или код проверки закомментирован)")
	}
}
