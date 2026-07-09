# Docker Compose: ClickHouse Server + Tabix

## What is this

`docker-compose.yaml` starts:

- **ClickHouse Server**
  - HTTP: `localhost:8123`
  - Native: `localhost:9000`
- **Tabix** (web UI) on `localhost:8124`

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
curl -s http://localhost:8123/ping
```

Run a query:

```bash
curl -s "http://localhost:8123/?query=SELECT%201"
```

ClickHouse client:

```bash
docker compose exec -T clickhouse clickhouse-client --query "SELECT version()"
```

## UI

Open: `http://localhost:8124`

Connection:

- URL: `http://clickhouse:8123`
