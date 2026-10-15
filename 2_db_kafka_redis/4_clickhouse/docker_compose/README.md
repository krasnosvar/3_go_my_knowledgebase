# Docker Compose: ClickHouse Server

## Что поднимается

`docker-compose.yaml` запускает один контейнер `clickhouse` (образ `clickhouse/clickhouse-server:latest`):

- HTTP-интерфейс: `localhost:8123` (в том числе встроенный веб-UI `http://localhost:8123/play`);
- native-протокол (для Go-драйвера): `localhost:9000`;
- пользователь `default` / пароль `admin`, база `test_db`;
- данные — в `./clickhouse_data` (каталог в `.gitignore`).

## Запуск и остановка

```bash
docker compose up -d
docker compose ps
```

Остановка:

```bash
docker compose down
```

## Проверка

Ping (без авторизации):

```bash
curl -s http://localhost:8123/ping
```

Запрос через HTTP:

```bash
curl -s --user default:admin "http://localhost:8123/?query=SELECT%201"
```

Через `clickhouse-client` внутри контейнера:

```bash
docker compose exec -T clickhouse clickhouse-client --password admin --query "SELECT version()"
```

## UI

Встроенный веб-интерфейс ClickHouse: `http://localhost:8123/play` (пользователь `default`, пароль `admin`).

## Дальше

Упражнения — [../exercises/](../exercises/README.md).
