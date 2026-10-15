# 04. Data Scaling (Масштабирование данных)

Как хранить и масштабировать состояние, когда одна БД перестаёт справляться: репликация для отказоустойчивости и чтений, шардирование для записи и объёма, компромиссы согласованности (CAP/PACELC, кворумы), распределённые транзакции и выбор хранилища SQL/NoSQL. Теория подкреплена runnable-примерами на стандартной библиотеке Go и упражнениями с самопроверкой.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория, разбитая по подтемам:
    *   [01_replication.md](./01_theory/01_replication.md) — репликация: leader–follower, синхронная vs асинхронная (и semi-sync), отставание реплик (replication lag), аномалии чтения и чтение с реплик.
    *   [02_sharding_partitioning.md](./01_theory/02_sharding_partitioning.md) — шардирование vs партиционирование, выбор ключа шардирования, стратегии hash/range/directory, ре-шардинг и consistent hashing.
    *   [03_cap_theorem.md](./01_theory/03_cap_theorem.md) — CAP-теорема (C/A/P), выбор CP vs AP при разделении сети на примерах, кратко PACELC.
    *   [04_consistency.md](./01_theory/04_consistency.md) — уровни согласованности: strong/линеаризуемость, eventual, read-your-writes, monotonic reads; кворумы и правило `W + R > N`.
    *   [05_distributed_tx.md](./01_theory/05_distributed_tx.md) — распределённые транзакции кратко: 2PC, Saga (+ компенсации), transactional outbox.
    *   [06_sql_vs_nosql.md](./01_theory/06_sql_vs_nosql.md) — SQL vs NoSQL: семейства NoSQL (KV/документные/wide-column/граф/OLAP), критерии выбора, polyglot persistence, NewSQL.
    *   [07_use_cases.md](./01_theory/07_use_cases.md) — практические применения: эволюция хранилища под нагрузкой, типовые связки, чек-лист рассуждения на собеседовании, частые ошибки.
*   **[02_examples/](./02_examples/)** — runnable-примеры (детерминированные, без сети и внешних зависимостей, завершаются сами):
    *   [01_consistent_hashing](./02_examples/01_consistent_hashing/main.go) — кольцо консистентного хеширования (`hash/fnv`) с виртуальными узлами: распределение ключей и минимальное перемещение (~1/N) при добавлении узла.
    *   [02_shard_router](./02_examples/02_shard_router/main.go) — маршрутизация ключа в шард двумя стратегиями: `hash(key) % N` и по диапазону; показывает, в какой шард попадает ключ.
    *   [03_quorum_rw](./02_examples/03_quorum_rw/main.go) — симуляция кворума `W + R > N` на in-memory «репликах»: запись в W, чтение из R; почему пересечение множеств даёт согласованность (горутины синхронизированы через `sync`).
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   `01_pick_shard` — выбор шарда по хеш-ключу `hash(key) % n`.
    *   `02_reshard_moved` — доля перемещённых ключей при ре-шардинге наивного `hash%N`.
    *   `03_classify_cap` — определить CP или AP по описанию поведения при разделении сети.
    *   `04_check_quorum` — **проверка кворума `W + R > N` с валидацией параметров (Хардкор)**.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач.

## 🚀 Как проходить тему

1. Прочитайте теорию по подтемам в [01_theory/](./01_theory/), начиная с [01_replication.md](./01_theory/01_replication.md).
2. Запустите и изучите примеры из папки `02_examples/` через `go run` (все офлайн, только stdlib; работает корневой `go.mod` репозитория):
   ```bash
   go run ./02_examples/01_consistent_hashing/main.go
   ```
3. Выполните задачи в `03_exercises/`. Запустите встроенный тест для самопроверки:
   ```bash
   go run ./03_exercises/04_check_quorum/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`.
