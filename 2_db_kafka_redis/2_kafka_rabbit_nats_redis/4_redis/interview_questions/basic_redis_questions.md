
# ТОП-20 вопросов по Redis для Go-разработчика бэкенда

## Простые вопросы (Junior уровень)

### 1. Что такое Redis и для каких задач он используется?

**Что ожидают услышать:**

- In-memory key-value datastore (часто: cache/kv store), поддерживает persistence
- Очень низкая латентность, простые операции
- Типичные кейсы: кеш, сессии, rate limiting, distributed locks, pub/sub, очереди, leaderboard

### 2. Какие базовые типы данных Redis вы знаете?

**Что ожидают услышать:**

- String, Hash, List, Set, Sorted Set
- Также: Stream, Bitmap, HyperLogLog, Geo, JSON (через Redis Stack)
- Важно: операции атомарны на уровне команды

### 3. Чем отличается `SET` от `HSET`? Когда что использовать?

**Что ожидают услышать:**

- `SET` хранит строку/байты целиком
- `HSET` — “объект” (словарь) внутри ключа, удобно обновлять отдельные поля
- `HSET` экономит ключи и снижает overhead при большом числе полей

### 4. Как задавать TTL и как проверить время жизни ключа?

**Что ожидают услышать:**

- TTL через `EXPIRE`, `PEXPIRE`, либо `SET key value EX <sec>` / `PX <ms>`
- Проверка: `TTL key`, `PTTL key`
- Удаление TTL: `PERSIST key`

### 5. Что происходит при истечении TTL? Как Redis удаляет ключи?

**Что ожидают услышать:**

- Lazy expiration (при обращении ключ может быть удален)
- Active expiration (фоновые выборки и удаление)
- Важно понимать, что удаление не строго “в момент X”, но примерно

## Средние вопросы (Middle уровень)

### 6. Какие стратегии eviction существуют и чем они отличаются?

**Что ожидают услышать:**

- Настройка `maxmemory` и `maxmemory-policy`
- Политики: noeviction, allkeys-lru, volatile-lru, allkeys-lfu, volatile-lfu, allkeys-random, volatile-random, volatile-ttl
- LRU/LFU — приблизительные (approximate), важны для кешей

### 7. Как реализовать кеш с защитой от stampede (thundering herd)?

**Что ожидают услышать:**

- Lock around recompute (mutex key)
- “Soft TTL” (stale-while-revalidate) + фоноваая перестройка
- Randomized TTL (jitter)
- Singleflight в Go + ограничение конкуренции

### 8. Как реализовать distributed lock в Redis и какие есть подводные камни?

**Что ожидают услышать:**

- `SET lock value NX PX <ms>` как базовый примитив
- Unlock через Lua: сравнить value и удалить ключ атомарно
- Риски: clock drift, долгие паузы GC/stop-the-world, потеря lock при истечении TTL
- Redlock как спорный алгоритм; важно обсуждать требования к надежности

### 9. Какие есть паттерны rate limiting в Redis?

**Что ожидают услышать:**

- Fixed window: `INCR` + `EXPIRE`
- Sliding window log: ZSET (дорого по памяти)
- Sliding window counter: два окна + интерполяция
- Token bucket / leaky bucket (через Lua)

### 10. Чем Pub/Sub отличается от Streams? Когда выбирать Streams?

**Что ожидают услышать:**

- Pub/Sub: не хранит сообщения, нет ack/ретраев, подписчики должны быть онлайн
- Streams: persistent log, consumer groups, ack, pending, replay
- Streams подходят для очередей/фоновых задач и событий с обработкой “как минимум один раз”

### 11. Как работает Redis Streams consumer group (XREADGROUP, ACK, pending)?

**Что ожидают услышать:**

- Сообщение может быть доставлено consumer’у, пока не ACK — считается pending
- Есть PEL (pending entries list)
- Можно делать retry/claim через `XCLAIM`/`XAUTOCLAIM`
- `XGROUP CREATE` + starting id, управление ретеншном `XTRIM`/`MAXLEN`

### 12. Какие команды/подходы дают атомарность для read-modify-write?

**Что ожидают услышать:**

- Транзакции `MULTI/EXEC` (опционально с `WATCH` для optimistic locking)
- Lua-скрипты (`EVAL`) — атомарно на сервере
- Пайплайнинг не дает атомарности, только снижает RTT

## Продвинутые вопросы (Senior уровень)

### 13. Как устроена persistence в Redis: RDB vs AOF?

**Что ожидают услышать:**

- RDB snapshot: компактно, быстро восстанавливать, но возможна потеря данных между snapshot’ами
- AOF: журнал команд, лучше durability, но может быть больше размер и время восстановления
- AOF rewrite, настройки `appendfsync` (always/everysec/no)
- На практике часто используют RDB + AOF

### 14. Что такое replication в Redis и как происходит failover (Sentinel)?

**Что ожидают услышать:**

- Primary/replica репликация (асинхронная)
- Sentinel мониторит master, голосование/quorum, promotion replica
- Возможен split brain при сетевых проблемах, важны таймауты

### 15. Redis Cluster: как работает шардирование и что такое hash slots?

**Что ожидают услышать:**

- 16384 hash slots, ключи мапятся в слот по CRC16
- Слот назначен node’у, миграция слотов при ребалансе
- Ограничения multi-key команд: только в пределах одного слота
- Hash tags `{...}` чтобы “приклеить” связанные ключи к одному слоту

### 16. Что такое hot keys и как с ними бороться?

**Что ожидают услышать:**

- Ключ, по которому слишком много запросов → нагрузка на один shard/CPU
- Шардинг на уровне приложения (key sharding), реплика чтения, локальный кеш
- Для counters: распределение по нескольким ключам + суммирование

### 17. Как оптимизировать работу Redis-клиента в Go?

**Что ожидают услышать:**

- Connection pooling, keep-alive, лимиты на соединения
- Pipelining/Batching где уместно
- Таймауты (Dial/Read/Write), контекст
- Обработка ошибок сети, retry с backoff, circuit breaker

### 18. Как мониторить Redis и на какие метрики смотреть?

**Что ожидают услышать:**

- `INFO`: used_memory, connected_clients, instantaneous_ops_per_sec
- Latency percentiles (latency-monitor), slowlog
- Evicted keys, expired keys, keyspace hits/misses
- Репликация: lag, backlog, sync status

### 19. Какие антипаттерны использования Redis вы знаете?

**Что ожидают услышать:**

- Хранить “вечные” ключи без TTL в кеше
- Делать большие ключи/значения (мегабайты) без нужды
- Делать O(N) операции на больших структурах в проде без понимания (например LRANGE огромных list)
- Использовать Pub/Sub там, где нужна надежная доставка
- Отсутствие лимитов/квот на ключи и память

### 20. Опишите архитектуру highload системы с Redis (кеш + очереди). На что обращаешь внимание?

**Что ожидают услышать:**

- Разделение use-case: cache vs durable queue (Streams) vs locks
- Настройки `maxmemory`, eviction, TTL policy, jitter
- HA: Sentinel/Cluster, read replicas, failover тестирование
- Наблюдаемость, алерты на latency/evictions/memory/replication lag
- Стратегия деградации: fallback на БД, ограничение функциональности

---

## Дополнительные темы для обсуждения

- **Data modeling**: выбор структур (hash vs string vs zset), оценка памяти
- **Безопасность**: ACL, TLS, изоляция сети
- **Интеграционные паттерны**: cache-aside, write-through, write-behind
- **Тестирование**: integration тесты с Testcontainers, нагрузочное тестирование
- **Миграции**: смена key format, backfill, dual write
