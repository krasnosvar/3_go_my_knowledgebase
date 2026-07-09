# Cache: pitfalls

## 1. Cache stampede (thundering herd)

Тысячи запросов одновременно промахиваются мимо кеша → все идут в БД, БД падает.

Решения:

- **singleflight** (`golang.org/x/sync/singleflight`): дедуплицирует одновременные запросы за одним ключом, в БД идёт только один.

```go
val, err, _ := group.Do(key, func() (any, error) { return loadFromDB(key) })
```

- **probabilistic early refresh**: обновлять чуть раньше TTL по случайной вероятности.
- **lock per key** в кеше — только один writer обновляет при miss.

## 2. Stale data навсегда

Данные обновили в БД, кеш всё ещё показывает старое. Стратегии:

- TTL короткий → stale-окно ограничено.
- Invalidate сразу после update.
- Versioning ключа: при update меняй версию, старые ключи естественно протухают.

## 3. Кеш стал источником правды

Если в кеше что-то, чего нет в БД — у тебя уже не кеш, а primary store. Так делается специально (write-back), но осознанно.

Случайный «кеш как primary» — вечный source of bugs.

## 4. Memory leak

`map[string]item` без ограничения размера → растёт пока процесс жив.

Решения:

- LRU (например, `hashicorp/golang-lru/v2`).
- Background cleanup горутина, удаляющая просроченные.
- Лимит по count или по байтам.

## 5. Сериализация туда-обратно

Для Redis/memcached данные сериализуются (обычно JSON). Это аллокации.

```go
data, _ := json.Marshal(user)
redis.Set(key, data, ttl)

// при чтении
data, _ := redis.Get(key)
var u User
json.Unmarshal(data, &u)
```

В hot path — рассмотри MessagePack или Protobuf.

## 6. Cache stampede при инвалидации

Вычистил весь кеш «для чистоты» → все клиенты пошли в БД. Лучше — мягкая инвалидация (TTL), а не полный flush.

## 7. Игнорирование ошибок Redis

```go
redis.Get(ctx, key) // ошибка → fallback к БД?
```

Если Redis упал, а ты этого не учёл — приложение упадёт. Всегда: cache fail → continue with DB. Кеш — оптимизация, не критическая зависимость.

## 8. Кеш на write-heavy

Для данных, которые часто меняются, кеш создаёт два класса проблем:

- Постоянная инвалидация бьёт по производительности больше, чем помогает кеш.
- Stale data всегда близко.

Не кешируй то, что меняется чаще, чем читается.

## 9. Слишком сложные ключи

```go
key := fmt.Sprintf("user:%d:perms:role=%s:org=%d:scope=%v", id, role, orgID, scope)
```

Длинный ключ → больше памяти, медленнее lookup. Если ключ становится сложным — может, ты кешируешь не то?

## 10. Распределённый кеш без TTL

Redis без TTL → ключи копятся бесконечно. Всегда: либо `EXPIRE`, либо явная eviction-политика на стороне Redis.

## Правило

- TTL всегда.
- Cache failure = degrade gracefully к источнику.
- Лимит размера + cleanup.
- Метрики hit/miss/evictions.
- singleflight против stampede.
