package main

import "time"

// Эти структуры нужны для того, чтобы маппить строки из базы в Go‑объекты
// User соответствует таблице users
type User struct {
	ID, OrderCount int
}

// Order соответствует таблице orders
type Order struct {
	ID, UserID int
	CreatedAt  time.Time
}

// Item соответствует таблице items
type Item struct {
	ID, OrderID int
	Price       float64
}
