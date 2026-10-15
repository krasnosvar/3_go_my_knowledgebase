# Упражнения: sync_package

Каждая задача — это `main.go` с функцией/методом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL` и не зависают — ожидания ограничены таймаутами) — ваша цель довести их до `PASS`. Только стандартная библиотека.

Запускайте с флагом `-race`: решение засчитывается, только если печатает `PASS` и детектор гонок молчит.

Запуск шаблона (из папки темы):

```bash
go run -race ./03_exercises/01_waitgroup/main.go
```

Запуск эталонного решения:

```bash
go run -race ./04_exercises_answers/01_waitgroup_answers/main.go
```

## Список задач (от простого к сложному)

1. [01_waitgroup](./01_waitgroup/main.go) — параллельная сумма частей. Решение: [01_waitgroup_answers](../04_exercises_answers/01_waitgroup_answers/main.go).
2. [02_mutex](./02_mutex/main.go) — счётчик под `Mutex`. Решение: [02_mutex_answers](../04_exercises_answers/02_mutex_answers/main.go).
3. [03_rwmutex](./03_rwmutex/main.go) — конфиг: много читателей, один писатель. Решение: [03_rwmutex_answers](../04_exercises_answers/03_rwmutex_answers/main.go).
4. [04_once](./04_once/main.go) — ленивая инициализация ровно один раз. Решение: [04_once_answers](../04_exercises_answers/04_once_answers/main.go).
5. [05_cond](./05_cond/main.go) — ожидание публикации через `Cond.Broadcast`. Решение: [05_cond_answers](../04_exercises_answers/05_cond_answers/main.go).
6. [06_map_with_mutex](./06_map_with_mutex/main.go) — конкурентный подсчёт слов в map. Решение: [06_map_with_mutex_answers](../04_exercises_answers/06_map_with_mutex_answers/main.go).
7. [07_pool](./07_pool/main.go) — переиспользование `bytes.Buffer` через `Pool`. Решение: [07_pool_answers](../04_exercises_answers/07_pool_answers/main.go).
8. [08_barrier](./08_barrier/main.go) — барьер между двумя фазами работы. Решение: [08_barrier_answers](../04_exercises_answers/08_barrier_answers/main.go).
