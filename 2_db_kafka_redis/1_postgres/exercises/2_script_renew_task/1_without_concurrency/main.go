package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Standard library compatible driver
)

const batchSize = 100000

func main() {
	// 1. Подключение к базе данных
	connStr := os.Getenv("DATABASE_URL")
	if os.Getenv("DATABASE_URL") == "" {
		connStr = "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable"
	}
	log.Println("Using connection string:", connStr)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()
	// Check if the connection is actually alive
	err = db.Ping()
	if err != nil {
		log.Fatalf("Database unreachable: %v\n", err)
	}
	fmt.Println("Successfully connected!")

	// 2. SQL запрос для создания таблицы и наполнения ее данными
	// Используем обратные кавычки для многострочного текста
	query := `
	DROP TABLE IF EXISTS users;

	CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		order_count INT
	);

	INSERT INTO users (name, order_count)
	SELECT 
		(ARRAY[
			'Den', 'Max', 'John', 'Bob', 'Alice', 
			'Ann', 'Vasya', 'Petya', 'Masha', 'Olya'
		])[floor(random() * 10 + 1)], 
		0 
	FROM generate_series(1, 1000000);
	`
	// Выполнение SQL запроса
	start := time.Now()
	// Для тяжелых запросов лучше использовать контекст с таймаутом,
	// так как вставка 1 млн строк может занять несколько секунд.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		log.Fatalf("Ошибка при выполнении инициализации: %v", err)
	}
	fmt.Printf("Готово! 1 000 000 строк добавлено за %v\n", time.Since(start))

	// 3. Вытащить кол-во строк
	// Define a variable to hold the fetched count
	var totalRows int
	// Execute the query and "Scan" the result into the variable
	// QueryRow is used when you expect a single result
	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&totalRows)
	if err != nil {
		log.Fatalf("Failed to fetch count: %v", err)
	}
	// Use the fetched data
	fmt.Printf("Fetched count: %d\n", totalRows)

	// 4. Обновляем данные по частям, батчами по 100 000 записей
	//for i := 0; i < totalRows; i += batchSize
	//Что делает: Этот цикл шагает по всему объему данных (totalRows)
	//крупными шагами размером batchSize (например, по 100,000).
	//Переменная i: Это начало текущего батча (например: 0, 100000, 200000...).
	//Смысл: Разбить огромную задачу (1 млн записей) на управляемые куски по batchSize.
	// 4. Обновляем данные по частям, батчами по 100 000 записей
	// Важно: ID в базе начинаются с 1. Разбиваем на диапазоны.
	for i := 0; i < totalRows; i += batchSize {
		start := i + 1
		end := start + batchSize - 1
		if end > totalRows {
			end = totalRows
		}

		// Внутренний цикл по конкретным ID в диапазоне (батче)
		for userID := start; userID <= end; userID++ {
			updQuery := `
        UPDATE users 
        SET order_count = 1 
        WHERE id = $1 AND order_count = 0
    `
			result, err := db.Exec(updQuery, userID)
			if err != nil {
				log.Fatalf("Ошибка при выполнении запроса: %v", err)
			}
			// Проверяем, была ли обновлена строка
			rowsAffected, _ := result.RowsAffected()
			if rowsAffected == 0 {
				fmt.Println("Запись не обновлена (либо ID не найден, либо значение уже не 0)")
			} else {
				fmt.Printf("Поле успешно обновлено для ID: %d\n", userID)
			}
		}
	}
}
