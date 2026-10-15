# Caching — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Redis-команды не запускались — это иллюстрация из теории.

## Зачем и когда

Подробно: [01_why_cache.md](./01_theory/01_why_cache.md).

| Уровень | Латентность |
| :--- | :--- |
| L1 cache | ~1 нс |
| RAM | ~100 нс |
| in-memory `map` | ~микросекунды (хеш + локи) |
| SSD | ~0.1 мс |
| Redis по сети | ~0.2–1 мс |
| запрос к БД | ~1–10 мс |
| HDD seek | ~10 мс |
| внешний API | ~50–500 мс |

| Формула | Пример |
| :--- | :--- |
| `hit ratio = hits / (hits + misses)` | 90% hit → нагрузка на БД падает в ~10 раз |
| `avg = hit_ratio * t_cache + (1 - hit_ratio) * t_source` | 0.9·0.5 мс + 0.1·10 мс ≈ 1.45 мс вместо 10 мс |

| Кэш уместен | Кэш вреден |
| :--- | :--- |
| read-heavy: читают намного чаще, чем меняют | данные меняются почти на каждый запрос |
| горячие ключи (80/20, Ципф) | равномерный доступ, датасет не помещается |
| допустима устареваемость | строгая согласованность (баланс, остаток при оплате) |
| источник дорогой (JOIN, агрегат, внешний API) | ответ уникален на запрос (персональный, одноразовый токен) |

## Паттерны

Подробно: [02_patterns.md](./01_theory/02_patterns.md).

| Паттерн | Чтение при промахе | Запись | Согласованность | Потеря данных | Где |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Cache-aside | приложение: get → miss → DB → set с TTL | DB, затем `DEL` ключа | eventual | нет | универсально, read-heavy |
| Read-through | кэш сам через loader | отдельно | eventual | нет | спрятать БД за кэшем |
| Write-through | (обычно + read-through) | кэш → DB синхронно | сильная для записанных ключей | нет | чтение сразу после записи |
| Write-behind | кэш | кэш → DB асинхронно батчами | слабая (лаг) | ⚠️ да | write-heavy, лаг допустим |

## TTL и вытеснение

Подробно: [03_ttl_eviction.md](./01_theory/03_ttl_eviction.md).

| Механизм | Удаляет | Как |
| :--- | :--- | :--- |
| TTL / expiration | протухшее по времени | lazy (при обращении) + active (фоновая выборка) |
| Eviction | при нехватке памяти | по политике вытеснения |
| Jitter | — | `TTL ± random(0, jitter)` — ключи не протухают разом |

| Политика | Выселяет | Метаданные | Стойкость к скану | Стоимость |
| :--- | :--- | :--- | :--- | :--- |
| LRU | давно не использованный | порядок доступа | низкая | O(1) |
| LFU | редко используемый | счётчики (+ aging) | высокая | выше |
| FIFO | давно вставленный | порядок вставки | низкая | O(1) |
| Random | случайный | нет | средняя | минимум |

| Redis `maxmemory-policy` | Поведение |
| :--- | :--- |
| `noeviction` | ошибка на запись |
| `allkeys-lru` / `allkeys-lfu` | типичный выбор для кэша |
| `volatile-lru` / `volatile-lfu` / `volatile-ttl` | только среди ключей с TTL |
| `allkeys-random` | случайный ключ |

LRU за O(1): `map` (key → узел) + двусвязный список `container/list`.

```go
package main

import (
	"container/list"
	"fmt"
)

type entry struct {
	key string
	val int
}

type LRU struct {
	cap   int
	ll    *list.List               // голова — most recently used
	items map[string]*list.Element // key → узел списка
}

func NewLRU(capacity int) *LRU {
	return &LRU{cap: capacity, ll: list.New(), items: map[string]*list.Element{}}
}

func (c *LRU) Get(key string) (int, bool) {
	if el, ok := c.items[key]; ok {
		c.ll.MoveToFront(el) // O(1)
		return el.Value.(*entry).val, true
	}
	return 0, false
}

func (c *LRU) Put(key string, val int) {
	if el, ok := c.items[key]; ok {
		el.Value.(*entry).val = val
		c.ll.MoveToFront(el)
		return
	}
	c.items[key] = c.ll.PushFront(&entry{key, val})
	if c.ll.Len() > c.cap { // выселяем хвост
		tail := c.ll.Back()
		c.ll.Remove(tail)
		delete(c.items, tail.Value.(*entry).key)
	}
}

func main() {
	c := NewLRU(2)
	c.Put("a", 1)
	c.Put("b", 2)
	c.Get("a")    // a стал самым свежим
	c.Put("c", 3) // выселен b
	_, okB := c.Get("b")
	va, okA := c.Get("a")
	fmt.Println(okB, va, okA) // false 1 true
}
```

## Проблемы и решения

Подробно: [04_problems.md](./01_theory/04_problems.md).

| Проблема | Суть | Решение |
| :--- | :--- | :--- |
| Stampede / thundering herd / dogpile | ключ протух → N одновременных промахов грузят БД | single-flight (`sync.Mutex` + карта in-flight; `golang.org/x/sync/singleflight`), лок `SET NX`, jitter, раннее вероятностное обновление, прогрев |
| Инвалидация | когда убрать устаревшее | TTL; событие на write path; версия/тег в ключе (для широкой инвалидации) |
| Stale data | кэш отстаёт от источника | решить допустимый лаг; TTL ограничивает максимум; критичное читать из источника |
| Cache penetration | поток запросов к несуществующим ключам | negative caching (маркер «нет» с коротким TTL), Bloom-фильтр |
| Cache pollution | разовый скан вымывает горячее | LFU / сегментированный LRU |
| Hot key | один ключ перегружает узел | локальный L1-кэш, реплика значения, шардирование |
| L1 + L2 рассинхрон | у инстансов свои локальные копии | инвалидация через pub/sub |

## Redis как кэш

Подробно: [05_redis_as_cache.md](./01_theory/05_redis_as_cache.md).

| | Локальный (`map` + `sync.RWMutex`) | Redis |
| :--- | :--- | :--- |
| Скорость | быстрее, без сети | доля мс |
| Общий для инстансов / согласованность | нет | да |
| Переживает рестарт | нет | да (persistence) |
| TTL, eviction, атомарность | руками | из коробки |

| Задача | Команда / приём |
| :--- | :--- |
| Значение с TTL | `SET key val EX <seconds>` |
| TTL отдельно | `EXPIRE key <seconds>` |
| Лок на пересчёт | `SET lock:key token NX EX 5` |
| Имена ключей | `entity:id:field`, например `user:42:profile` |
| Объект по полям | Hash (`user:42 → {name, email}`) |
| Рейтинги / последние N / уникальные | Sorted Set / List / HyperLogLog, Bitmap |
| Двухуровневый кэш | L1 in-process (hot key, без сети) + L2 Redis (общий) |
| Отказ Redis | таймауты, fail-fast, промахи идут в БД |

## Где кэшируют и что выбрать

Подробно: [06_use_cases.md](./01_theory/06_use_cases.md).

| Слой | Кэш |
| :--- | :--- |
| Браузер | `Cache-Control`, `ETag` |
| CDN | статика, кэшируемые ответы на edge |
| Reverse proxy | nginx / Varnish |
| Приложение | L1 in-process |
| Распределённый | Redis / Memcached (L2) |
| БД | buffer pool |

| Сценарий | Паттерн | Особенности |
| :--- | :--- | :--- |
| Профиль, каталог | cache-aside + TTL | read-heavy |
| Сессии / токены | write-through / aside | TTL = жизнь сессии |
| Лента, агрегаты, топы | cache-aside + версия/тег | широкая инвалидация |
| Rate limiting | атомарные операции Redis | TTL = окно лимита |
| Ответы внешнего API | cache-aside + negative caching | короткий TTL, кэшировать ошибки |
| Дорогое вычисление | cache-aside + single-flight | защита от stampede |

## Ловушки

- ⚠️ На записи обновлять ключ вместо `DEL` → гонка двух записей оставит старое значение; гонку «чтение с промахом + запись» не убирает и `DEL` — TTL обязателен как предохранитель.
- ⚠️ Одинаковый TTL у множества ключей → синхронный шторм перезагрузок.
- ⚠️ Кэш — не источник истины: система обязана работать (медленнее) без него.
- ⚠️ Write-behind теряет несброшенные записи при падении кэша.
- ⚠️ Без negative caching запросы к отсутствующим ключам всегда бьют в БД.
- ⚠️ Без метрик hit ratio и p99 непонятно, помогает ли кэш.
- ⚠️ TTL и eviction — разные механизмы и работают одновременно.
