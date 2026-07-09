# 03. Caching (Кэширование)

Кэш — один из главных инструментов ускорения и разгрузки системы: он держит быстрые копии дорогих в получении данных. Здесь разобрано, зачем нужен кэш, какими паттернами с ним работают, как устроены TTL и вытеснение, какими проблемами кэш расплачивается за скорость (stampede, инвалидация, stale data) и как всё это ложится на Redis. Теория подкреплена runnable-примерами на стандартной библиотеке Go и упражнениями с самопроверкой.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория, разбитая по подтемам:
    *   [01_why_cache.md](./01_theory/01_why_cache.md) — зачем кэш: латентность и порядки величин, разгрузка БД, стоимость; hit ratio; где кэш уместен, а где вреден.
    *   [02_patterns.md](./01_theory/02_patterns.md) — паттерны cache-aside, read-through, write-through, write-behind: схемы, trade-offs, гонки согласованности и почему на записи ключ удаляют.
    *   [03_ttl_eviction.md](./01_theory/03_ttl_eviction.md) — TTL и способы удаления протухшего, jitter; размер кэша и стратегии вытеснения LRU/LFU/FIFO/random; политики Redis.
    *   [04_problems.md](./01_theory/04_problems.md) — типичные проблемы: cache stampede/thundering herd, инвалидация, stale data, negative caching, hot key, cache pollution.
    *   [05_redis_as_cache.md](./01_theory/05_redis_as_cache.md) — Redis как кэш: структуры данных, TTL, паттерны в Go-сервисе (фрагменты кода как иллюстрация, без реального подключения), эксплуатация.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практические применения: уровни кэширования в стеке, типовые сценарии и выбор паттерна, чеклист «ставить ли кэш», итоговые принципы.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, детерминированные, завершаются сами):
    *   [01_cache_aside](./02_examples/01_cache_aside/main.go) — cache-aside поверх «медленной БД» (`time.Sleep`): промах → загрузка → запись в кэш → попадание; наглядная разница во времени и счётчик обращений к «БД».
    *   [02_lru_eviction](./02_examples/02_lru_eviction/main.go) — LRU-кэш на `map` + `container/list` с ограничением ёмкости; демонстрация вытеснения least recently used.
    *   [03_stampede_singleflight](./02_examples/03_stampede_singleflight/main.go) — защита от cache stampede: 50 параллельных горутин на один ключ вызывают загрузку ровно один раз (single-flight через `sync.Mutex` + карта in-flight; счётчик через `atomic`).
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   `01_ttl_cache` — кэш с TTL: `Set`/`Get`, протухший ключ считается отсутствующим (детерминированные часы вместо сна).
    *   `02_lru_cache` — LRU-кэш `Get`/`Put` с вытеснением при переполнении.
    *   `03_hit_ratio` — учёт hits/misses и подсчёт `hit ratio` с корректной обработкой нуля обращений.
    *   `04_invalidation` — cache-aside с инвалидацией на записи: `Update` удаляет ключ, следующий `Get` перечитывает свежее из «БД».
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач.

## 🚀 Как проходить тему

1. Прочитайте теорию по подтемам в [01_theory/](./01_theory/), начиная с [01_why_cache.md](./01_theory/01_why_cache.md).
2. Запустите и изучите примеры из папки `02_examples/` через `go run` (всё офлайн, без внешних зависимостей и без `go.mod`):
   ```bash
   go run ./02_examples/01_cache_aside/main.go
   ```
3. Выполните задачи в `03_exercises/`. Запустите встроенный тест для самопроверки:
   ```bash
   go run ./03_exercises/02_lru_cache/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`.
