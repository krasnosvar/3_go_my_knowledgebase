# Упражнения: goroutines

Каждая задача — это `main.go` с функцией/методом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL` и не зависают — ожидания ограничены таймаутами) — ваша цель довести их до `PASS`. Только стандартная библиотека.

Запускайте с флагом `-race`: решение засчитывается, только если печатает `PASS` и детектор гонок молчит.

Запуск шаблона (из папки темы):

```bash
go run -race ./03_exercises/01_basic_go_keyword/main.go
```

Запуск эталонного решения:

```bash
go run -race ./04_exercises_answers/01_basic_go_keyword_answers/main.go
```

## Список задач (от простого к сложному)

1. [01_basic_go_keyword](./01_basic_go_keyword/main.go) — запуск n горутин через `go` без ожидания. Решение: [01_basic_go_keyword_answers](../04_exercises_answers/01_basic_go_keyword_answers/main.go).
2. [02_waitgroup](./02_waitgroup/main.go) — конкурентный запуск задач и ожидание через `WaitGroup`. Решение: [02_waitgroup_answers](../04_exercises_answers/02_waitgroup_answers/main.go).
3. [03_race_awareness](./03_race_awareness/main.go) — общий счётчик без data race (`Mutex`). Решение: [03_race_awareness_answers](../04_exercises_answers/03_race_awareness_answers/main.go).
4. [04_loop_capture](./04_loop_capture/main.go) — явная передача переменной цикла в горутину. Решение: [04_loop_capture_answers](../04_exercises_answers/04_loop_capture_answers/main.go).
5. [05_channel_signal](./05_channel_signal/main.go) — сигнал о завершении закрытием канала. Решение: [05_channel_signal_answers](../04_exercises_answers/05_channel_signal_answers/main.go).
6. [06_worker_basic](./06_worker_basic/main.go) — worker, читающий задачи из канала до закрытия. Решение: [06_worker_basic_answers](../04_exercises_answers/06_worker_basic_answers/main.go).
7. [07_leak_prevention](./07_leak_prevention/main.go) — генератор с путём остановки — без утечки горутины. Решение: [07_leak_prevention_answers](../04_exercises_answers/07_leak_prevention_answers/main.go).
8. [08_graceful_stop](./08_graceful_stop/main.go) — остановка нескольких горутин и ожидание через `WaitGroup`. Решение: [08_graceful_stop_answers](../04_exercises_answers/08_graceful_stop_answers/main.go).
