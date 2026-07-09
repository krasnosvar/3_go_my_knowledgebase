package main

import "time"

// User — структура для таблицы users
// GORM сам создаст таблицу и поля, если их нет
type User struct {
	ID         uint      `gorm:"primaryKey"` // primary key
	OrderCount int       // количество заказов
	CreatedAt  time.Time // автоматически заполняется
	UpdatedAt  time.Time // автоматически заполняется
}

// Order — структура для таблицы orders
type Order struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint // внешний ключ на users
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Item — структура для таблицы items
type Item struct {
	ID      uint    `gorm:"primaryKey"`
	OrderID uint    // внешний ключ на orders
	Price   float64 // цена товара
}
