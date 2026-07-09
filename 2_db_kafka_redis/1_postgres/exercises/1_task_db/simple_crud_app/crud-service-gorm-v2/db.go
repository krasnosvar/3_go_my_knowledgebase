package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB — создаём подключение к Postgres через GORM
func InitDB() *gorm.DB {
	// DSN — строка подключения к базе
	dsn := "host=localhost user=postgres password=admin dbname=shop port=5432 sslmode=disable"

	// gorm.Open — открываем соединение
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	return db
}
