package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")
	if os.Getenv("DATABASE_URL") == "" {
		connStr = "postgres://postgres:admin@localhost:5432/shop?sslmode=disable"
	}
	log.Println("Using connection string:", connStr)
	db, err := InitDB(connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected successfully")
	defer db.Close()

	// вызов хендлеров на один энжпоинт - switch смотрит по типу запроса и перенаправляет на нужный хендлер
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			CreateUserHandler(db)(w, r)
		case http.MethodGet:
			GetUserHandler(db)(w, r)
		case http.MethodPut:
			UpdateUserHandler(db)(w, r)
		case http.MethodDelete:
			DeleteUserHandler(db)(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Блок для работы с транзакциями
	http.HandleFunc("/users_working_with_transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			CreateUserTxHandler(db)(w, r)
		case http.MethodPut:
			UpdateUserTxHandler(db)(w, r)
		case http.MethodDelete:
			DeleteUserTxHandler(db)(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// старт сервера
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
