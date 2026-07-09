# 10. Design Interview (System Design интервью)

Сборка всех предыдущих разделов базы в единый навык: как структурированно проходить system design интервью и тренироваться на типовых задачах. Раздел преимущественно **текстовый** — это методология и разборы задач; runnable-кода немного (base62-код и калькулятор оценок), потому что суть здесь — процесс мышления, а не строки на Go.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория и разборы, разбитые по подтемам:
    *   [01_approach.md](./01_theory/01_approach.md) — как проходить интервью: сбор функциональных/нефункциональных требований, уточняющие вопросы, оценка нагрузки, high-level дизайн, детализация, поиск узких мест, trade-offs; тайминг на 45 минут.
    *   [02_estimations.md](./01_theory/02_estimations.md) — прикидки «на салфетке»: степени 10, latency numbers every programmer should know, расчёт QPS / хранилища / bandwidth / кэша с примерами.
    *   [03_building_blocks.md](./01_theory/03_building_blocks.md) — строительные блоки (LB, кэш, реплики, шардинг, очередь, CDN, blob storage) и когда их доставать; со ссылками на разделы 01–09 базы.
    *   [04_example_url_shortener.md](./01_theory/04_example_url_shortener.md) — полный разбор «сокращателя ссылок»: требования → оценки → API → модель данных → генерация ключа → масштабирование → узкие места.
    *   [05_example_news_feed.md](./01_theory/05_example_news_feed.md) — разбор «ленты новостей»: fanout on write vs on read (гибрид), кэш ленты, хранение.
    *   [06_more_problems.md](./01_theory/06_more_problems.md) — краткие наброски по остальным задачам: чат/мессенджер, rate limiter, платёжная система (идемпотентность/консистентность), файловое хранилище.
*   **[02_examples/](./02_examples/)** — runnable-примеры на чистой stdlib (детерминированные, завершаются):
    *   [01_base62_shortener](./02_examples/01_base62_shortener/main.go) — base62-кодирование инкрементного ID в короткий код и обратно (для url shortener); печатает пары id → код → id.
    *   [02_capacity_estimation](./02_examples/02_capacity_estimation/main.go) — калькулятор оценок: QPS из DAU, объём хранилища и bandwidth из размера записи/ответа; печатает расчёт.
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   `01_base62` — реализовать base62 encode/decode с round-trip.
    *   `02_storage_estimation` — посчитать объём хранилища по входным параметрам.
    *   `03_qps_estimation` — посчитать средний и пиковый QPS из DAU.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач.

## 🚀 Как проходить тему

1. Изучите методологию: начните с [01_approach.md](./01_theory/01_approach.md), затем [02_estimations.md](./01_theory/02_estimations.md) (научитесь считать) и [03_building_blocks.md](./01_theory/03_building_blocks.md) (из чего собирать).
2. Разберите две полные задачи-эталона: [04_example_url_shortener.md](./01_theory/04_example_url_shortener.md) и [05_example_news_feed.md](./01_theory/05_example_news_feed.md) — они показывают методологию в действии.
3. Пробегите наброски остальных задач в [06_more_problems.md](./01_theory/06_more_problems.md) и найдите «центральное решение» каждой.
4. Запустите примеры через `go run` (офлайн, без внешних зависимостей и `go.mod`):
   ```bash
   go run ./02_examples/01_base62_shortener/main.go
   go run ./02_examples/02_capacity_estimation/main.go
   ```
5. Выполните задачи в `03_exercises/` до `PASS`, затем сверьтесь с `04_exercises_answers/`:
   ```bash
   go run ./03_exercises/01_base62/main.go
   ```
6. Тренируйтесь вслух: возьмите задачу из [06_more_problems.md](./01_theory/06_more_problems.md) и пройдите её целиком по чеклисту из [01_approach.md](./01_theory/01_approach.md) — требования → оценки → API → дизайн → узкие места.
