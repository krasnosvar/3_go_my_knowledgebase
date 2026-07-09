# Profiling

Раздел про инструменты Go для анализа производительности и поведения приложения в рантайме.

## Что внутри

- `01_theory/01_pprof.md` — `net/http/pprof`, сбор профилей CPU / memory / goroutine / mutex.
- `01_theory/02_profile_types.md` — типы профилей и как их интерпретировать.
- `02_examples/01_pprof_playground/` — учебное приложение, симулирующее 4 типичные проблемы (CPU hot loop, memory leak, goroutine leak, mutex contention).

## Связанные разделы

- `5_testing/04_benchmarks` — бенчмарки + `-cpuprofile` / `-memprofile`.
- `7_under_the_hood/02_memory_gc` — теория памяти и GC.

## Workflow

1. Воспроизвести проблему (нагрузка, бенч, прод).
2. Снять профиль (`go tool pprof`, `/debug/pprof/...`).
3. Открыть в `pprof -http=:8081` или `top` в интерактиве.
4. Найти hot path или утечку.
5. Исправить и подтвердить бенчем.
