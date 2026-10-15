# Redis: упражнения на Go

## Зачем

Подключиться к Redis из Go через [go-redis v9](https://github.com/redis/go-redis), проверить соединение и выполнить базовые `SET`/`GET`.

## Пререквизиты

- Go 1.22+ и Docker (с `docker compose`).
- Запущенный стенд из [../docker_compose/](../docker_compose/README.md): Redis на `localhost:6379` без пароля.
- Теория: [../interview_questions/theory.md](../interview_questions/theory.md).

## Файлы

- [ex1/main.go](./ex1/main.go) — отдельный Go-модуль (`go.mod` внутри): `PING`, `SET user:100:name`, `GET` того же ключа.

## Как запустить

```bash
cd ../docker_compose && docker compose up -d && cd ../exercises/ex1
go run .
```

Проверить значение из консоли:

```bash
docker exec redis redis-cli GET user:100:name
```

## Что попробовать самостоятельно

1. Задайте ключу TTL (`rdb.Set(ctx, key, value, 10*time.Second)`) и проверьте `TTL user:100:name` в `redis-cli`.
2. Замените строку на хеш: `HSet`/`HGetAll` для объекта пользователя.
3. Реализуйте cache-aside поверх «медленной» функции (паттерн из [3_system_design/03_caching](../../../../3_system_design/03_caching/)).

## Когда дальше

Следующий подраздел — [../../../4_clickhouse/](../../../4_clickhouse/README.md).
