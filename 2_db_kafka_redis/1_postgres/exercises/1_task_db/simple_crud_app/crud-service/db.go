// Можно вынести работу с базой в отдельную папку db/.
// Тогда в файле db/db.go будет package db.
// В main.go ты сможешь импортировать его как import "crud-service/db"
package main

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// var connStr string = ""

func InitDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Устанавливаем таймаут на проверку соединения (например, 5 секунд)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Пытаемся достучаться до базы
	if err := db.PingContext(ctx); err != nil {
		db.Close() // Обязательно закрываем, если не смогли достучаться
		return nil, err
	}

	// Настройки пула (опционально, но полезно)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
