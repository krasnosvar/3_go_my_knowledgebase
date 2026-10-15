# Docker Compose: MongoDB

## Что поднимается

`docker-compose.yaml` запускает один контейнер `mongodb` (образ `mongo:latest`):

- порт `localhost:27017`;
- root-пользователь `admin` / пароль `admin` (база аутентификации — `admin`);
- база по умолчанию `test_db`;
- данные — в `./mongo_data` (каталог в `.gitignore`).

## Запуск и остановка

```bash
docker compose up -d
docker compose ps
```

Остановка:

```bash
docker compose down
```

Данные остаются в `./mongo_data`; чтобы начать с чистой базы, удалите этот каталог после остановки.

## Подключение

```bash
docker exec -it mongodb mongosh -u admin -p admin --authenticationDatabase admin
```

Строка подключения для приложений: `mongodb://admin:admin@localhost:27017`.

## Дальше

Упражнения — [../exercises/](../exercises/README.md).
