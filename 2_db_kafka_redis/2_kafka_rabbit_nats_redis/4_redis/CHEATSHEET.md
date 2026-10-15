# Redis — шпаргалка

Подробности и объяснения — в [interview_questions/theory.md](./interview_questions/theory.md)
и [basic_redis_questions.md](./interview_questions/basic_redis_questions.md). ⚠️ — частая ловушка.
Команды CLI не запускались — проверяйте на стенде из [docker_compose/](./docker_compose/README.md):
`docker compose exec -T redis redis-cli <команда>`. Go-клиент — `github.com/redis/go-redis/v9`.

## Какой тип данных выбрать

Подробно: [theory.md → Когда использовать какой тип данных?](./interview_questions/theory.md#когда-использовать-какой-тип-данных).

| Тип | Что это | Для чего |
| :--- | :--- | :--- |
| String | байты до 512 MB | кеш HTML/JSON, счётчики, флаги, токены, простые сессии, bitmap |
| Hash | поля → значения внутри ключа (≈ `map[string]string`) | объекты (user, product), частичные обновления, настройки |
| List | двусторонний список | очереди (FIFO), стеки (LIFO), activity feed, последние N элементов |
| Set | неупорядоченные уникальные строки | теги, уникальные посетители, друзья, пересечения множеств |
| Sorted Set (ZSET) | уникальные элементы со `score` | leaderboard, priority queue, time-series по timestamp, sliding-window rate limit |
| Stream | append-only лог с consumer groups и ACK | очередь с гарантиями, event sourcing, audit log |
| Ещё | Bitmap, HyperLogLog, Geo, JSON (Redis Stack) | — |

## Команды по типам

Подробно: [theory.md → Типы данных в Redis](./interview_questions/theory.md#типы-данных-в-redis).

| Задача | redis-cli | go-redis |
| :--- | :--- | :--- |
| Записать / прочитать | `SET k v` / `GET k` | `rdb.Set(ctx, k, v, 0)` / `rdb.Get(ctx, k).Result()` |
| Атомарный счётчик | `INCR k`, `INCRBY k 10` | `rdb.Incr(ctx, k)`, `rdb.IncrBy(ctx, k, 10)` |
| Несколько ключей | `MSET k1 v1 k2 v2` / `MGET k1 k2` | `rdb.MSet(...)` / `rdb.MGet(...)` |
| Бит | `SETBIT bm 7 1` / `GETBIT bm 7` | `rdb.SetBit(ctx, "bm", 7, 1)` |
| Поля hash | `HSET user:1 name Alice age 25` / `HGET user:1 name` / `HGETALL user:1` | `rdb.HSet(...)`, `rdb.HGet(...)`, `rdb.HGetAll(...)` |
| Очередь | `LPUSH q a` / `RPUSH q b` / `LPOP q` / `RPOP q` | `rdb.LPush`, `rdb.RPush`, `rdb.LPop`, `rdb.RPop` |
| Блокирующее чтение очереди | `BLPOP q 5` | `rdb.BLPop(ctx, 5*time.Second, "q")` |
| Срез / длина списка | `LRANGE q 0 -1` / `LLEN q` | `rdb.LRange(ctx, "q", 0, -1)`, `rdb.LLen` |
| Множество | `SADD s a b` / `SISMEMBER s a` / `SMEMBERS s` / `SCARD s` / `SREM s a` | `rdb.SAdd`, `rdb.SIsMember`, ... |
| Операции над множествами | `SINTER s1 s2` / `SUNION s1 s2` / `SDIFF s1 s2` | `rdb.SInter`, `rdb.SUnion`, `rdb.SDiff` |
| Случайный элемент | `SRANDMEMBER s` / `SPOP s` (с удалением) | `rdb.SRandMember`, `rdb.SPop` |
| Добавить в ZSET | `ZADD lb 1100 Alice` | `rdb.ZAdd(ctx, "lb", redis.Z{Score: 1100, Member: "Alice"})` |
| Топ-3 по убыванию | `ZREVRANGE lb 0 2 WITHSCORES` (6.2+: `ZRANGE lb 0 2 REV WITHSCORES`) | `rdb.ZRevRangeWithScores(ctx, "lb", 0, 2)` |
| Место / очки / +очки | `ZREVRANK lb Alice` (с 0) / `ZSCORE lb Alice` / `ZINCRBY lb 50 Bob` | `rdb.ZRevRank`, `rdb.ZScore`, `rdb.ZIncrBy` |
| Диапазон по score / удалить старые | `ZRANGEBYSCORE lb 1000 2000` / `ZREMRANGEBYSCORE k 0 <ts>` | `rdb.ZRangeByScore`, `rdb.ZRemRangeByScore` |
| Stream: добавить / прочитать | `XADD s * f v` / `XREAD COUNT 10 STREAMS s 0` / `XREVRANGE s + -` | `rdb.XAdd`, `rdb.XRead`, `rdb.XRevRange` |
| Stream: группа | `XGROUP CREATE s workers 0` / `XREADGROUP GROUP workers c1 STREAMS s >` / `XACK s workers <id>` | `rdb.XGroupCreate`, `rdb.XReadGroup`, `rdb.XAck` |
| Stream: зависшие | `XPENDING`, `XCLAIM` / `XAUTOCLAIM`, обрезка `XTRIM s MAXLEN 1000` | — |
| Ключа нет | `GET` → `(nil)` | `err == redis.Nil` |

## TTL и удаление ключей

Подробно: [theory.md → TTL и Expiration](./interview_questions/theory.md#ttl-и-expiration).

| Задача | Команда |
| :--- | :--- |
| Записать с TTL | `SET k v EX 1800` (сек) / `PX 500` (мс); Go: `rdb.Set(ctx, k, v, 30*time.Minute)` |
| Назначить TTL существующему | `EXPIRE k 60` / `PEXPIRE` / `EXPIREAT k <unix>` |
| Сколько осталось | `TTL k` / `PTTL k` (`-1` — без TTL, `-2` — ключа нет) |
| Снять TTL | `PERSIST k` |
| Существует ли | `EXISTS k` → `1` / `0` |

| Удаление истёкших | Как | Минус |
| :--- | :--- | :--- |
| Lazy | проверка TTL при обращении | память не освобождается сразу |
| Active | каждые 100 мс: 20 случайных ключей с TTL, удалить истёкшие, если > 25% истекло — повтор | удаление не строго в момент истечения |

## Eviction (`maxmemory-policy`)

Подробно: [theory.md → Cache Eviction Policies](./interview_questions/theory.md#cache-eviction-policies-политики-вытеснения).

| Политика | Кого вытесняет | Когда |
| :--- | :--- | :--- |
| `noeviction` (по умолчанию) | никого: ошибка записи при нехватке памяти | хранилище, где терять данные нельзя; не для кеша |
| `allkeys-lru` | давно не использованные среди всех ключей | универсальный кеш |
| `allkeys-lfu` | редко используемые | есть устойчиво «горячие» данные |
| `volatile-lru` / `volatile-lfu` | только ключи с TTL | кеш и постоянные данные в одном инстансе |
| `allkeys-random` / `volatile-random` | случайные | редко; непредсказуемо |
| `volatile-ttl` | ключи с ближайшим истечением | — |

Конфиг: `maxmemory 2gb` + `maxmemory-policy allkeys-lru`. LRU/LFU в Redis приблизительные (сэмплирование).

## Кеширование

Подробно: [theory.md → Redis как кеш](./interview_questions/theory.md#redis-как-кеш).

| Стратегия | Запись | Плюс / минус |
| :--- | :--- | :--- |
| TTL-based | `SET k v EX 3600` | просто / данные могут устареть |
| Cache-aside | UPDATE БД → `DEL` кеш; чтение: miss → БД → `SET` | просто / первый запрос после инвалидации медленный |
| Write-through | UPDATE БД → UPDATE кеш | всегда свежо / медленнее запись |
| Event-based | UPDATE БД → событие → UPDATE кеш | гибко / сложная архитектура |

| Stampede (thundering herd) | Защита |
| :--- | :--- |
| много промахов на один ключ одновременно | mutex-ключ на пересчёт, `singleflight` в Go |
| массовое истечение | jitter к TTL |
| долгий пересчёт | soft TTL / stale-while-revalidate + фоновое обновление |

## Атомарность: транзакции, Lua, pipeline

Подробно: [theory.md → Transactions и Pipelining](./interview_questions/theory.md#transactions-и-pipelining).

| Механизм | Атомарность | Когда |
| :--- | :--- | :--- |
| Одна команда (`INCR`, `SETNX`, `ZADD`) | да | всегда, когда хватает одной команды |
| `MULTI` / `EXEC` | команды выполняются подряд без вмешательства других клиентов; **отката нет** | несколько записей пачкой |
| `WATCH` + `MULTI`/`EXEC` | оптимистичная блокировка: `EXEC` вернёт nil, если ключ изменился → повторить | read-modify-write; Go: `rdb.Watch(ctx, fn, keys...)` + `tx.TxPipelined` |
| Lua (`EVAL`) | весь скрипт атомарно на сервере | сравнить-и-удалить, token bucket |
| Pipeline | **нет**, только меньше RTT | пачки независимых команд; Go: `pipe := rdb.Pipeline(); ...; pipe.Exec(ctx)` |

## Практические паттерны

Подробно: [theory.md → Практические паттерны](./interview_questions/theory.md#практические-паттерны).

| Задача | Решение |
| :--- | :--- |
| Rate limit: fixed window | `INCR rl:<ip>:<window>` + `EXPIRE` |
| Rate limit: sliding window log | ZSET: `ZREMRANGEBYSCORE` старых → `ZADD now` → `ZCARD` → `EXPIRE` в одном pipeline |
| Rate limit: token / leaky bucket | Lua-скрипт |
| Leaderboard | `ZADD` / `ZINCRBY`, топ — `ZREVRANGE 0 9 WITHSCORES`, место — `ZREVRANK + 1` |
| Distributed lock: захват | `SET lock:<name> <token> NX PX 30000` (Go: `rdb.SetNX(ctx, key, token, ttl)`) |
| Distributed lock: освобождение | Lua: `if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("del", KEYS[1]) else return 0 end` |
| Session store | `SET session:<id> <json> EX 86400` |
| Hot key / счётчик под нагрузкой | разбить на N ключей + суммировать, локальный кеш, реплики чтения |

## Pub/Sub vs Streams

Подробно: [theory.md → Pub/Sub vs Streams](./interview_questions/theory.md#pubsub-vs-streams).

| | Pub/Sub | Streams |
| :--- | :--- | :--- |
| Хранение | нет: подписчик оффлайн → сообщение потеряно | да, читается история |
| Подтверждение | нет | `XACK`, pending entries list, `XCLAIM` |
| Consumer groups | нет (fan-out всем подписчикам) | да |
| Команды | `PUBLISH ch msg`, `SUBSCRIBE ch`, `PSUBSCRIBE news:*` | `XADD`, `XREADGROUP`, `XACK` |
| Когда | live-уведомления, чат, real-time дашборды | очереди задач, event sourcing, аудит |

## Persistence

Подробно: [theory.md → Persistence](./interview_questions/theory.md#persistence-сохранение-данных).

| | RDB | AOF |
| :--- | :--- | :--- |
| Что | снимок памяти по расписанию (`save 900 1`, `save 300 10`, `save 60 10000`) | журнал всех записей + периодический rewrite |
| Потеря при сбое | с последнего снимка (минуты) | `appendfsync everysec` — до ~1 c; `always` — минимум; `no` — решает ОС |
| Файл / восстановление | компактный, быстро | больше, медленнее |
| Когда | кеш, бэкапы | сессии, критичные данные |

Рекомендация теории: включать оба (`appendonly yes` + `save ...`).

## Репликация, Sentinel, Cluster

Подробно: [theory.md → Replication](./interview_questions/theory.md#replication),
[theory.md → Redis Cluster](./interview_questions/theory.md#redis-cluster).

| Вариант | Что даёт | Ключевое |
| :--- | :--- | :--- |
| Primary / replica | масштаб чтения, копия данных | асинхронная; `replicaof <ip> 6379`, `masterauth`; первичная синхронизация через RDB snapshot |
| Sentinel | автоматический failover | кворум Sentinel-ов, повышение реплики (~30–60 c); риск split brain |
| Cluster | шардирование по 16384 hash slots | `slot = CRC16(key) % 16384`; обычно 3 primary + 3 replica; Go: `redis.NewClusterClient` |

| Ограничение Cluster | Обход |
| :--- | :--- |
| multi-key команды и `MULTI` только в одном слоте | hash tags: `user:{100}:profile`, `user:{100}:settings` |
| только база `0` (`SELECT` нет) | префиксы ключей |
| Pub/Sub рассылается по всему кластеру | sharded Pub/Sub `SPUBLISH` / `SSUBSCRIBE` (Redis 7.0+) |

## Эксплуатация

Подробно: [basic_redis_questions.md](./interview_questions/basic_redis_questions.md).

| Задача | Команда / настройка |
| :--- | :--- |
| Проверить связь | `redis-cli ping` → `PONG` |
| Память, клиенты, ops/s, hits/misses, evicted | `INFO` (`used_memory`, `connected_clients`, `instantaneous_ops_per_sec`, `keyspace_hits/misses`, `evicted_keys`) |
| Медленные команды | `SLOWLOG GET 10`, latency monitor |
| Безопасность | `requirepass`, ACL, TLS, `rename-command FLUSHALL ""` |
| Go-клиент | один клиент на приложение (пул), таймауты Dial/Read/Write, `ctx`, retry с backoff |
| UI стенда | Redis Commander `http://localhost:8081`, RedisInsight `http://localhost:5540` |

## Ловушки

- ⚠️ Кеш без TTL или с `noeviction` → память кончается, записи падают с ошибкой.
- ⚠️ `volatile-*` политика, а ключей с TTL нет → ведёт себя как `noeviction`.
- ⚠️ `MULTI`/`EXEC` не откатывает: ошибка одной команды во время выполнения не отменяет остальные.
- ⚠️ Pipeline не атомарен — для read-modify-write нужен `WATCH` или Lua.
- ⚠️ Lock через `DEL` без проверки токена → снимете чужой lock после истечения своего TTL.
- ⚠️ Lock с TTL короче работы (GC-пауза, долгий запрос) → два владельца; Redlock спорен.
- ⚠️ `KEYS *`, `SMEMBERS`/`LRANGE 0 -1`/`HGETALL` на огромных структурах → O(N) блокирует однопоточный сервер; используйте `SCAN`/`SSCAN`/`HSCAN`.
- ⚠️ Pub/Sub для надёжной доставки → сообщения теряются у оффлайн-подписчиков; берите Streams.
- ⚠️ В sliding-window rate limit `member` = timestamp в секундах → запросы в одну секунду схлопнутся; добавляйте уникальный суффикс.
- ⚠️ Чтение с реплики сразу после записи может вернуть старое значение (асинхронная репликация).
- ⚠️ `ZREVRANK` считает с 0 — для «места» прибавьте 1.
