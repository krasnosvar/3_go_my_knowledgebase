# Cache layer

Кеш — снижение нагрузки на медленный источник (БД, внешний API) за счёт хранения часто запрашиваемых данных в быстром (память, Redis).

## Уровни кеша

- **In-process** — `map`, `sync.Map`, или библиотеки типа `bigcache`, `freecache`. Самый быстрый, но не разделяется между инстансами.
- **Distributed** — Redis, Memcached. Один кеш на N инстансов.
- **CDN / edge** — для статики и часто запрашиваемых публичных endpoint'ов.

В реальном backend часто комбинируют: L1 in-process + L2 Redis.

## Базовая реализация in-memory

```go
type Cache[V any] struct {
    mu    sync.RWMutex
    items map[string]item[V]
    ttl   time.Duration
}

type item[V any] struct {
    val V
    exp time.Time
}

func (c *Cache[V]) Get(k string) (V, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    it, ok := c.items[k]
    if !ok || time.Now().After(it.exp) {
        var zero V
        return zero, false
    }
    return it.val, true
}

func (c *Cache[V]) Set(k string, v V) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.items[k] = item[V]{val: v, exp: time.Now().Add(c.ttl)}
}
```

Для прода — нужны: ограничение размера (LRU), background cleanup, метрики hit/miss.

## Cache-aside (lazy)

Самый распространённый паттерн:

```go
func GetUser(ctx context.Context, id int) (User, error) {
    if u, ok := cache.Get(key(id)); ok {
        return u, nil // hit
    }
    u, err := db.Get(ctx, id) // miss
    if err != nil { return User{}, err }
    cache.Set(key(id), u)
    return u, nil
}
```

Кеш заполняется по запросу. Простая реализация, хорошо подходит для read-heavy.

## Read-through / write-through

Кеш сам ходит в БД (read-through) или сам пишет в БД при write (write-through). Чище API, но требует библиотечной обёртки.

## Write-back

Запись в кеш, асинхронно в БД. Самая быстрая запись, но риск потери данных при падении.

## Invalidation

Самая сложная часть. Стратегии:

- **TTL** — данные живут N секунд. Просто, но stale-окно есть.
- **Explicit invalidate** — после write в БД, удаляем из кеша.
- **Write-through** — каждый write обновляет кеш сразу.
- **Versioning** — ключ кеша содержит версию (`user:42:v3`).

## Метрики

Минимум для мониторинга:

- hit_rate = hits / (hits + misses).
- avg latency на cache.Get.
- size кеша.
- evictions / sec.

Если hit_rate < 50% — кеш приносит мало пользы или размер мал.

## Когда НЕ нужен кеш

- Источник уже быстрый (in-memory структура).
- Запросы уникальные — нет hit'ов.
- Данные часто меняются — хранить устаревшие опасно.
- Требуется strong consistency — кеш расходится с БД.

## Правило

- Cache-aside как default.
- TTL обязателен — данные в кеше не должны жить вечно.
- Метрики hit/miss — без них непонятно, помогает ли кеш.
- Кешируй то, что **дорого** считать, не всё подряд.
