# Docker Compose: Redis + RedisInsight

## What is this

`docker-compose.yaml` starts:

- **Redis** on `localhost:6379` (AOF enabled)
- **RedisInsight** on `localhost:5540`
- **Redis Commander** on `localhost:8081` (пре-настроен на подключение к redis)

## Start / stop

```bash
docker compose up -d
docker compose ps
```

Stop:

```bash
docker compose down
```

## Connect

Ping:

```bash
docker compose exec -T redis redis-cli ping
```

Get/set:

```bash
docker compose exec -T redis redis-cli SET mykey 123
docker compose exec -T redis redis-cli GET mykey
```

## UI

- **Redis Commander**: `http://localhost:8081` (самый простой, работает сразу)
- **RedisInsight**: `http://localhost:5540` (более мощный, нужно добавить БД)

В RedisInsight добавьте базу:

- Host: `redis`
- Port: `6379`

