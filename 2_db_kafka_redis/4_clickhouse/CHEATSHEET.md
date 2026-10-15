# ClickHouse — шпаргалка

Подробности и объяснения — в [interview_questions/theory.md](./interview_questions/theory.md) и
[basic_click_questions.md](./interview_questions/basic_click_questions.md). ⚠️ — частая ловушка.
Команды CLI/SQL не запускались — проверяйте на стенде из [docker_compose/](./docker_compose/README.md).

## Стенд и клиенты

Подробно: [docker_compose/README.md](./docker_compose/README.md).

| Задача | Команда |
| :--- | :--- |
| Поднять / статус / остановить | `docker compose up -d` / `docker compose ps` / `docker compose down` |
| Порты | HTTP `localhost:8123` (веб-UI `/play`), native `localhost:9000` (Go-драйвер) |
| Учётка стенда | `default` / `admin`, база `test_db` |
| Ping | `curl -s http://localhost:8123/ping` |
| Запрос по HTTP | `curl -s --user default:admin "http://localhost:8123/?query=SELECT%201"` |
| `clickhouse-client` в контейнере | `docker compose exec -T clickhouse clickhouse-client --password admin --query "SELECT version()"` |
| Go-драйвер | `github.com/ClickHouse/clickhouse-go/v2`: `clickhouse.Open(&clickhouse.Options{Addr: []string{"127.0.0.1:9000"}, Auth: clickhouse.Auth{...}})` |
| Пул соединений в драйвере | `MaxOpenConns`, `MaxIdleConns`, `ConnMaxLifetime` в `clickhouse.Options` |

## OLAP vs OLTP

Подробно: [theory.md → Что такое ClickHouse](./interview_questions/theory.md#что-такое-clickhouse), [колоночное хранение](./interview_questions/theory.md#колоночное-хранение-данных).

| Задача | ClickHouse? |
| :--- | :--- |
| Аналитика, логи, метрики, time-series, `GROUP BY` по миллиардам строк | да |
| Real-time дашборды, предагрегаты | да (+ materialized views) |
| OLTP, транзакции, частые `UPDATE`/`DELETE` | нет → PostgreSQL |
| Документы | нет → MongoDB |
| Маленькие таблицы (< 1 млн строк), построчная вставка | нет |

| Колоночное хранение | Эффект |
| :--- | :--- |
| Значения одного столбца лежат вместе | сильное сжатие однотипных данных |
| Читаются только нужные столбцы | `SELECT a, b` по широкой таблице быстрый, `SELECT *` — дорогой |
| Агрегат по столбцу | векторная обработка, мало I/O |
| Вставка одной строки | дорого — пишется во все столбцы, рождается part |

## Движки таблиц

Подробно: [theory.md → MergeTree](./interview_questions/theory.md#mergetree-и-партиционирование), [репликация и шардирование](./interview_questions/theory.md#репликация-и-шардирование).

| Нужно | Движок |
| :--- | :--- |
| Основная аналитическая таблица | `MergeTree()` |
| Копии на нескольких серверах (через ZooKeeper / ClickHouse Keeper) | `ReplicatedMergeTree('/clickhouse/tables/{shard}/t', '{replica}')` |
| Суммировать строки с одинаковым ключом при merge (счётчики, MV) | `SummingMergeTree()` |
| «Обновления» через перезапись версии строки | `ReplacingMergeTree` |
| Запрос/вставка по всем шардам (сама данных не хранит) | `Distributed('cluster', 'db', 'local_table', rand())` |
| Вспомогательные: в памяти / простой лог / буфер вставок | `Memory`, `Log`, `Buffer` |

## Ключи: ORDER BY, PRIMARY KEY, PARTITION BY

Подробно: [theory.md → PRIMARY KEY](./interview_questions/theory.md#primary-key-в-clickhouse).

```sql
CREATE TABLE events (
    date       Date,
    user_id    UInt32,
    event_type String,
    url        String
)
ENGINE = MergeTree()
PARTITION BY toYYYYMM(date)      -- партиция = месяц
ORDER BY (user_id, date)         -- он же PRIMARY KEY (разреженный индекс)
TTL date + INTERVAL 90 DAY DELETE;
```

| Понятие | Суть |
| :--- | :--- |
| `ORDER BY` | порядок сортировки внутри part; по умолчанию и PRIMARY KEY |
| PRIMARY KEY | **не** уникальность, а разреженный индекс: одна метка на гранулу (8192 строки) |
| Запрос по ключу | читаются только гранулы, чей диапазон может содержать значение |
| `PARTITION BY` | физическое разбиение для удаления/перемещения целиком, не для поиска |

| Выбор ключа | Хорошо | Плохо |
| :--- | :--- | :--- |
| `PARTITION BY` | `toYYYYMM(date)`, `toMonday(date)`, `toDate(ts)`, `(region, toYYYYMM(date))` | `user_id` (тысячи партиций), `rand()` |
| `ORDER BY` | самые фильтруемые поля слева: `(region, user_id, timestamp)` | `(timestamp, user_id, region)` при фильтре по `region`, `user_id` |

## Вставка, parts, merge

Подробно: [theory.md → Батчевые вставки](./interview_questions/theory.md#батчевые-вставки), [Parts и Merge](./interview_questions/theory.md#parts-и-процесс-merge).

| Задача | Как |
| :--- | :--- |
| Вставить из Go | `batch, _ := conn.PrepareBatch(ctx, "INSERT INTO events")` → `batch.Append(...)` в цикле → `batch.Send()` |
| Размер батча | 10 000–100 000 строк, ~10–100 MB |
| Много мелких продюсеров | `async_insert`, `Buffer`-таблица, Kafka → батч-писатель |
| Ручное слияние parts | `OPTIMIZE TABLE events;` |
| Удалить старые данные | `ALTER TABLE events DROP PARTITION '202401';` (мгновенно) вместо `DELETE` |
| Автоочистка / холодное хранение | `TTL created_at + INTERVAL 90 DAY DELETE`, `TTL created_at + INTERVAL 30 DAY TO DISK 'cold'` |
| Сжатие столбца | `CODEC(Delta, ZSTD)`; по умолчанию LZ4; `DoubleDelta` — время, `Gorilla` — float, `T64` — целые |

Каждый `INSERT` создаёт иммутабельный part; фоновый merge сливает parts; `SELECT` читает все parts →
много мелких parts = медленно и ошибка `Too many parts`.

## Materialized views

Подробно: [theory.md → Materialized Views](./interview_questions/theory.md#materialized-views).

```sql
CREATE MATERIALIZED VIEW events_hourly
ENGINE = SummingMergeTree()
ORDER BY (hour, event_type)
AS SELECT toStartOfHour(timestamp) AS hour, event_type, count() AS event_count
FROM events GROUP BY hour, event_type;

-- читать всегда с доагрегацией: merge мог ещё не схлопнуть строки
SELECT hour, event_type, sum(event_count) FROM events_hourly GROUP BY hour, event_type;
```

| MV — да | MV — нет |
| :--- | :--- |
| одни и те же агрегаты запрашиваются постоянно | агрегаты каждый раз разные |
| исходная таблица огромная | данные редко читают / таблица маленькая |

## Типы и диагностика

Подробно: [basic_click_questions.md](./interview_questions/basic_click_questions.md#средние-вопросы-middle-уровень).

| Задача | Как |
| :--- | :--- |
| Дата / время | `Date`, `Date32`, `DateTime`, `DateTime64`, пояс: `DateTime('Europe/Moscow')`; в Go → `time.Time` |
| План запроса | `EXPLAIN SELECT ...` |
| История запросов | `system.query_log`: `read_rows`, `read_bytes`, `memory_usage`, длительность |
| Метрики сервера | `system.metrics`, `system.events` |
| Ранний фильтр | `PREWHERE` (обычно ClickHouse переносит условие из `WHERE` сам) |
| Трассировка запроса | свой Query ID на каждый запрос |

## Ловушки

- ⚠️ Одиночные `INSERT` в цикле → тысячи parts и `Too many parts`; только батчи.
- ⚠️ PRIMARY KEY не гарантирует уникальность — дубликаты допустимы.
- ⚠️ `ReplacingMergeTree`/`SummingMergeTree` схлопывают строки только при merge, момент не гарантирован — читайте с `sum()`/`FINAL`/`argMax`.
- ⚠️ `PARTITION BY user_id` или другим высококардинальным полем → слишком много партиций и parts.
- ⚠️ `DELETE`/`UPDATE` — тяжёлые мутации, а не OLTP-операции; старые данные — через `DROP PARTITION` или TTL.
- ⚠️ JOIN больших таблиц без денормализации и `SELECT *` по широкой таблице.
- ⚠️ Нет timeout в `context` у запросов из Go; ошибки `Too many simultaneous queries` и `Memory limit exceeded` — повторять с backoff.
- ⚠️ Порог `Too many parts` настраивается (`parts_to_throw_insert`) и зависит от версии; не полагайтесь на конкретное число.
