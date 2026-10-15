# Тема 11. Хеш-таблицы (Maps)

Этот раздел посвящен детальному изучению хеш-таблиц (`map[K]V`) в Go: их внутреннему устройству (Swiss tables в Go 1.24+, `hmap` до Go 1.24), неадресуемости значений, утечкам памяти при удалении ключей и защите от race conditions.

## 📋 Пререквизиты

*   Пройдена предыдущая тема: [10_strings](../10_strings/).

> ⏩ **Забегание вперёд.** Некоторые примеры и упражнения используют конструкции из следующих тем:
> *   `06_maps_concurrency` и пример `03_product_backend` используют горутины и `sync` — подробно в [3_concurrency/04_sync_package](../../3_concurrency/04_sync_package/).
> *   `08_maps_lru_cache` использует двусвязный список `container/list` — см. [8_algorithms_dsa/01_linked_list](../../8_algorithms_dsa/01_linked_list/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория, разбитая по подтемам:
    *   [01_basics_operations.md](./01_theory/01_basics_operations.md) — основы: объявление `map[K]V`, инициализация через `make`/литерал, nil map, запись и чтение, comma-ok, `delete`, случайный порядок обхода.
    *   [02_internals_hmap.md](./01_theory/02_internals_hmap.md) — внутреннее устройство: Swiss tables (Go 1.24+) и историческая схема `hmap` (до Go 1.24) — бакеты `bmap`, `tophash`, overflow buckets, эвакуация; рандомизация порядка обхода.
    *   [03_unaddressability.md](./01_theory/03_unaddressability.md) — неадресуемость значений мапы: почему нельзя `&m[k]` и `m[k].Field = ...`, и как это обойти.
    *   [04_memory_delete.md](./01_theory/04_memory_delete.md) — почему `delete` не возвращает память ОС и как делать ручную усадку мапы пересозданием.
    *   [05_concurrency.md](./01_theory/05_concurrency.md) — небезопасность конкурентного доступа, защита через `sync.RWMutex` и когда уместен `sync.Map`.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — краткий обзор практических применений мап (registry, set, потокобезопасное хранилище, сортировка, LRU-кэш).
*   **[02_examples/](./02_examples/)** — практические примеры:
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — паттерн Command Registry для диспетчеризации консольных команд по имени за $O(1)$.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — симуляция утечки памяти при вызове `delete()` на большой мапе и её устранение через пересоздание мапы с замером памяти с помощью `runtime.MemStats`.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — **потокобезопасное хранилище сессий пользователей (SessionStore) с использованием sync.RWMutex для разделения прав чтения/записи**.
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи для самопроверки с автопроверками `PASS`/`FAIL`:
    *   `01_maps_basics` — безопасный поиск значения по ключу через comma-ok.
    *   `02_maps_advanced` — изменение полей структур, находящихся внутри мапы (решение проблемы неадресуемости).
    *   `03_maps_set` — реализация структуры данных Set (множество) с оптимизацией памяти через `map[string]struct{}`.
    *   `04_maps_sorting` — извлечение и сортировка элементов мапы по ключам.
    *   `05_maps_key_constraints` — использование структур в качестве сравниваемых (comparable) ключей мапы.
    *   `06_maps_concurrency` — **потокобезопасный счетчик посещений с использованием sync.Mutex**.
    *   `07_maps_memory` — **динамическая усадка (shrink) мапы для освобождения неиспользуемых бакетов из памяти**.
    *   `08_maps_lru_cache` — **реализация алгоритма вытеснения LRU Cache на базе map и двусвязного списка container/list (Хардкор)**.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения.

## 🚀 Как проходить тему

1. Прочитайте теорию по подтемам в [01_theory/](./01_theory/), начиная с [01_basics_operations.md](./01_theory/01_basics_operations.md).
2. Запустите и изучите примеры из папки `02_examples/` через `go run`:
   ```bash
   go run ./02_examples/02_system_infrastructure/main.go
   ```
3. Выполните задачи в `03_exercises/`. Запустите встроенный тест для самопроверки:
   ```bash
   go run ./03_exercises/08_maps_lru_cache/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`.

## ➡️ Что дальше

Когда все упражнения темы печатают `PASS`, переходите к теме [12_errors](../12_errors/).
