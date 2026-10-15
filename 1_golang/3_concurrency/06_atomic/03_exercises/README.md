# Упражнения: atomic

Каждая задача — это `main.go` с функцией/методом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL` и не зависают — ожидания ограничены таймаутами) — ваша цель довести их до `PASS`. Только стандартная библиотека.

Запускайте с флагом `-race`: решение засчитывается, только если печатает `PASS` и детектор гонок молчит.

Запуск шаблона (из папки темы):

```bash
go run -race ./03_exercises/01_add/main.go
```

Запуск эталонного решения:

```bash
go run -race ./04_exercises_answers/01_add_answers/main.go
```

## Список задач (от простого к сложному)

1. [01_add](./01_add/main.go) — атомарный инкремент из многих горутин. Решение: [01_add_answers](../04_exercises_answers/01_add_answers/main.go).
2. [02_load_store](./02_load_store/main.go) — состояние сервиса на `atomic.Int32`. Решение: [02_load_store_answers](../04_exercises_answers/02_load_store_answers/main.go).
3. [03_compare_swap](./03_compare_swap/main.go) — CAS-цикл: атомарный максимум. Решение: [03_compare_swap_answers](../04_exercises_answers/03_compare_swap_answers/main.go).
4. [04_bool_flag](./04_bool_flag/main.go) — одноразовый флаг закрытия на `atomic.Bool`. Решение: [04_bool_flag_answers](../04_exercises_answers/04_bool_flag_answers/main.go).
5. [05_spin_wait](./05_spin_wait/main.go) — ожидание флага с `runtime.Gosched` и таймаутом. Решение: [05_spin_wait_answers](../04_exercises_answers/05_spin_wait_answers/main.go).
6. [06_counter](./06_counter/main.go) — статистика запросов на атомиках. Решение: [06_counter_answers](../04_exercises_answers/06_counter_answers/main.go).
7. [07_once_like](./07_once_like/main.go) — однократный запуск на CAS и отличие от `sync.Once`. Решение: [07_once_like_answers](../04_exercises_answers/07_once_like_answers/main.go).
8. [08_hybrid](./08_hybrid/main.go) — map под мьютексом + метрики на атомиках. Решение: [08_hybrid_answers](../04_exercises_answers/08_hybrid_answers/main.go).
