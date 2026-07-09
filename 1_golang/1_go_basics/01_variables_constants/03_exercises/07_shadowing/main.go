package main

import (
	"errors"
	"fmt"
)

type Database struct {
	Connected bool
}

func connectDB() (*Database, error) {
	return &Database{Connected: true}, nil
}

func initSystem() (*Database, error) {
	var db *Database
	var err error
	_ = err

	hasConfig := true

	if hasConfig {
		// TODO: Исправьте эту строчку. Здесь допущена ошибка затенения (shadowing)
		// из-за использования :=, в результате чего локальный db затеняет внешний.
		// Внешний db остается nil!
		db, err := connectDB()
		if err != nil {
			return nil, err
		}
		_ = db
	}

	// Проверяем внешнюю переменную db
	if db == nil {
		return nil, errors.New("база данных осталась nil (ошибка shadowing!)")
	}

	return db, nil
}

func main() {
	fmt.Println("--- Проверка 07_shadowing ---")

	db, err := initSystem()
	if err == nil && db != nil && db.Connected {
		fmt.Println("Тест инициализации DB: PASS")
	} else {
		fmt.Printf("Тест инициализации DB: FAIL (ошибка: %v)\n", err)
	}
}
