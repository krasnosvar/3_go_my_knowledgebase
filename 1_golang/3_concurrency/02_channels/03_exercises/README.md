# Упражнения: channels

Каждая задача — это `main.go` с функцией/методом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL` и не зависают — ожидания ограничены таймаутами) — ваша цель довести их до `PASS`. Только стандартная библиотека.

Запускайте с флагом `-race`: решение засчитывается, только если печатает `PASS` и детектор гонок молчит.

Запуск шаблона (из папки темы):

```bash
go run -race ./03_exercises/01_unbuffered/main.go
```

Запуск эталонного решения:

```bash
go run -race ./04_exercises_answers/01_unbuffered_answers/main.go
```

## Список задач (от простого к сложному)

1. [01_unbuffered](./01_unbuffered/main.go) — передача значения через небуферизованный канал. Решение: [01_unbuffered_answers](../04_exercises_answers/01_unbuffered_answers/main.go).
2. [02_buffered](./02_buffered/main.go) — заполнение буферизованного канала без читателя. Решение: [02_buffered_answers](../04_exercises_answers/02_buffered_answers/main.go).
3. [03_close_range](./03_close_range/main.go) — producer закрывает канал, consumer читает через `range`. Решение: [03_close_range_answers](../04_exercises_answers/03_close_range_answers/main.go).
4. [04_directional](./04_directional/main.go) — sender/receiver с направленными каналами. Решение: [04_directional_answers](../04_exercises_answers/04_directional_answers/main.go).
5. [05_fan_out](./05_fan_out/main.go) — несколько worker'ов читают один канал, порядок результатов сохранён. Решение: [05_fan_out_answers](../04_exercises_answers/05_fan_out_answers/main.go).
6. [06_fan_in](./06_fan_in/main.go) — слияние двух каналов с закрытием выхода. Решение: [06_fan_in_answers](../04_exercises_answers/06_fan_in_answers/main.go).
7. [07_timeout](./07_timeout/main.go) — чтение с таймаутом через `time.After`. Решение: [07_timeout_answers](../04_exercises_answers/07_timeout_answers/main.go).
8. [08_pipeline](./08_pipeline/main.go) — конвейер из стадий «умножить» → «отфильтровать». Решение: [08_pipeline_answers](../04_exercises_answers/08_pipeline_answers/main.go).
