# Docker Compose: Postgres primary/replica + PgBouncer

## What this starts

`docker-compose.yaml` starts:

- **PostgreSQL primary** (writes) on `localhost:5432`
- **PostgreSQL replica** (read-only standby)
- **PgBouncer** on `localhost:6432` (aliases: `primary`, `replica`)
- **pgAdmin** on `localhost:5050` (pre-configured server pointing to PgBouncer)

Notes:

- `./init-db` contains only minimal init needed for replication user.
- If you want a totally clean state: `docker compose down -v` and then `docker compose up -d`.

## Start / stop

```bash
docker compose up -d
docker compose ps
```

Stop:

```bash
docker compose down
```

## Connect (docker exec / docker compose exec)

Primary:
```
docker compose exec -T db-primary psql -U postgres -d postgres
```

Replica:
```
docker compose exec -T db-replica psql -U postgres -d postgres
```

PgBouncer (connect to the *primary* alias):

```bash
docker compose exec -T pgbouncer psql "postgres://postgres:admin@localhost:6432/primary"
```

PgBouncer (connect to the *replica* alias):

```bash
docker compose exec -T pgbouncer psql "postgres://postgres:admin@localhost:6432/replica"
```

## Health checks

### Check container health

```bash
docker compose ps
```

### Check replication on primary (should show 1 row for the standby)

```bash
docker compose exec -T db-primary psql -U postgres -d postgres -c "
SELECT
  pid,
  application_name,
  client_addr,
  state,
  sync_state,
  write_lag,
  flush_lag,
  replay_lag
FROM pg_stat_replication;
"
```

### Check replica is really a standby + WAL receiver state

```bash
docker compose exec -T db-replica psql -U postgres -d postgres -c "
SELECT pg_is_in_recovery() AS is_replica;
SELECT status, receive_start_lsn, latest_end_lsn, latest_end_time
FROM pg_stat_wal_receiver;
"
```

### Check replication lag (LSN diff; smaller is better)

```bash
docker compose exec -T db-primary psql -U postgres -d postgres -c "
SELECT
  application_name,
  pg_size_pretty(pg_wal_lsn_diff(pg_current_wal_lsn(), replay_lsn)) AS replay_lag
FROM pg_stat_replication;
"
```

### PgBouncer admin (SHOW POOLS / SHOW STATS)

Your `./pgbouncer/userlist.txt` contains:

- `postgres` / password `admin`
- `admin` / password `admin`

Connect to the PgBouncer admin console DB (`pgbouncer`) and run `SHOW` commands:

```bash
docker compose exec -T pgbouncer psql "postgres://admin:admin@localhost:6432/pgbouncer" -c "SHOW POOLS;"
docker compose exec -T pgbouncer psql "postgres://admin:admin@localhost:6432/pgbouncer" -c "SHOW STATS;"
docker compose exec -T pgbouncer psql "postgres://admin:admin@localhost:6432/pgbouncer" -c "SHOW DATABASES;"
```

## Execute SQL (commands and scripts)

List databases:

```bash
docker compose exec -T db-primary psql -U postgres -d postgres -c "\\l"
```

Create a database:

```bash
docker compose exec -T db-primary psql -U postgres -d postgres -c "CREATE DATABASE shop;"
```

Run a single SQL command against a DB:

```bash
docker compose exec -T db-primary psql -U postgres -d shop -v ON_ERROR_STOP=1 -c "SELECT now();"
```

Run a SQL script file from your host:

```bash
docker compose exec -T db-primary psql -U postgres -d shop -v ON_ERROR_STOP=1 < ./schema.sql
```

Run a script that contains `CREATE DATABASE ...` (connect to `postgres` first):

```bash
docker compose exec -T db-primary psql -U postgres -d postgres -v ON_ERROR_STOP=1 < ./init.sql
```
