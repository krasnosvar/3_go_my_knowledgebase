# MongoDB: упражнения

## Зачем

Базовый CRUD над коллекцией из Go на официальном драйвере `go.mongodb.org/mongo-driver` и то же самое вручную в `mongosh`.

## Пререквизиты

- Go 1.22+ и Docker (с `docker compose`).
- Запущенный стенд из [../docker_compose/](../docker_compose/README.md): MongoDB на `localhost:27017`, пользователь `admin` / пароль `admin`.
- Теория: [../interview_questions/theory.md](../interview_questions/theory.md).

## Файлы

- [mongodb-example/mongodb.go](./mongodb-example/mongodb.go) — отдельный Go-модуль (`go.mod` внутри): вставка, чтение, обновление и удаление документа в коллекции `test_db.users`.

## Как запустить

```bash
cd ../docker_compose && docker compose up -d && cd ../exercises/mongodb-example
go run .
```

Если зависимости ещё не скачаны, `go run` загрузит их сам (нужен доступ в интернет); при ошибках про `go.sum` выполните `go mod tidy`.

## Проверка через mongosh

У MongoDB нет встроенного HTTP-API, поэтому вручную проверяем через `mongosh` внутри контейнера:

```bash
docker exec -it mongodb mongosh -u admin -p admin --authenticationDatabase admin
```

Команды внутри `mongosh`:

```javascript
use test_db

// Вставка документа
db.users.insertOne({name: "Denis", order_count: 5})

// Чтение
db.users.find({name: "Denis"})

// Обновление
db.users.updateOne({name: "Denis"}, {$set: {order_count: 10}})

// Удаление
db.users.deleteOne({name: "Denis"})
```

## Когда дальше

MongoDB — последний подраздел. Вернитесь к [../../README.md](../../README.md) или переходите к [system design](../../../3_system_design/README.md).
