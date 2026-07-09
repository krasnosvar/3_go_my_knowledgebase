package main

import (
	"database/sql"
	"fmt"

	_ "github.com/ClickHouse/clickhouse-go/v2" // драйвер для ClickHouse
)

func main() {
	// DSN — строка подключения
	dsn := "clickhouse://default:admin@localhost:9000/test_db"

	// sql.Open — создаём подключение через стандартный интерфейс database/sql
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Создаём таблицу (если её нет)
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id UInt64,
            name String,
            age UInt8
        ) ENGINE = MergeTree()
        ORDER BY id
    `)
	if err != nil {
		panic(err)
	}

	// Вставляем данные
	_, err = db.Exec("INSERT INTO users (id, name, age) VALUES (?, ?, ?)", 1, "Denis", 30)
	if err != nil {
		panic(err)
	}

	// Читаем данные
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id uint64
		var name string
		var age uint8
		if err := rows.Scan(&id, &name, &age); err != nil {
			panic(err)
		}
		fmt.Printf("User: id=%d, name=%s, age=%d\n", id, name, age)
	}
}
