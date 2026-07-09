package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// 1. Создание клиента
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// 2. Проверка соединения (Ping)
	status, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Ошибка подключения к Redis: %v", err)
	}
	fmt.Printf("Redis Ping: %s\n", status)

	// 3. Запись данных (SET)
	key := "user:100:name"
	value := "Antigravity"
	err = rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatalf("Ошибка при записи SET: %v", err)
	}
	fmt.Printf("Записано в Redis: %s = %s\n", key, value)

	// 4. Чтение данных (GET)
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Ошибка при чтении GET: %v", err)
	}
	fmt.Printf("Прочитано из Redis: %s = %s\n", key, val)

	// 5. Удаление (DEL) - опционально, для чистоты эксперимента
	// rdb.Del(ctx, key)
}
