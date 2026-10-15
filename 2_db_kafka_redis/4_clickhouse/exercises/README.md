# ClickHouse: упражнения

## Зачем

Подключиться к ClickHouse из Go через `database/sql`, создать таблицу на движке MergeTree, вставить и прочитать строки; затем повторить то же через HTTP-интерфейс.

## Пререквизиты

- Go 1.22+ и Docker (с `docker compose`).
- Запущенный стенд из [../docker_compose/](../docker_compose/README.md): HTTP на `localhost:8123`, native-протокол на `localhost:9000`, пользователь `default` / пароль `admin`, база `test_db`.
- Теория: [../interview_questions/theory.md](../interview_questions/theory.md).

## Файлы

- [click-example/go_click_connect_example.go](./click-example/go_click_connect_example.go) — отдельный Go-модуль (`go.mod` внутри): создаёт таблицу `users`, вставляет строку и читает её через драйвер `clickhouse-go`.

## Как запустить

```bash
cd ../docker_compose && docker compose up -d && cd ../exercises/click-example
go run .
```

## Проверка через HTTP-интерфейс

После запуска Go-примера таблица `users` в базе `test_db` уже существует:

```bash
# Простой запрос
curl --user default:admin "http://localhost:8123/?query=SELECT+1"

# Вставка данных
curl --user default:admin -X POST "http://localhost:8123/?database=test_db&query=INSERT+INTO+users+(id,name,age)+VALUES+(2,'Alex',25)"

# Чтение данных
curl --user default:admin "http://localhost:8123/?database=test_db&query=SELECT+*+FROM+users"
```

## Когда дальше

Дальше — документная база: [../../5_mongodb/](../../5_mongodb/README.md).
