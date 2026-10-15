# 7.4 Profiling

Инструменты Go для анализа производительности и поведения приложения в рантайме.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#04-профилирование) — раздел «04. Профилирование» общей шпаргалки раздела 7.

## Зачем

Оптимизировать «на глаз» почти всегда бесполезно. `pprof` показывает,
где реально тратится CPU, кто аллоцирует память, какие горутины зависли
и за какой мьютекс идёт борьба.

## Пререквизиты

- [01_runtime](../01_runtime/) и [02_memory_gc](../02_memory_gc/);
- [5_testing/04_benchmarks](../../5_testing/04_benchmarks/) — бенчмарки, `-cpuprofile` / `-memprofile`;
- `net/http` из [4_standard_library/05_net_http_net](../../4_standard_library/05_net_http_net/).

## Что внутри

- [01_theory/01_pprof.md](./01_theory/01_pprof.md) — `net/http/pprof`, сбор профилей CPU / memory / goroutine / mutex, Web UI, как читать результат.
- [01_theory/02_profile_types.md](./01_theory/02_profile_types.md) — типы профилей и когда какой использовать.
- [02_examples/01_pprof_playground/](./02_examples/01_pprof_playground/main.go) — учебное приложение с четырьмя типичными проблемами (CPU hot loop, утечка памяти, утечка горутин, конкуренция за мьютекс). Рядом лежат готовые профили `cpu.prof` и `mem.prof`, снятые с него.

Упражнений с самопроверкой нет: практика — найти все четыре проблемы
в «полигоне» (см. раздел «Практика» в [01_pprof.md](./01_theory/01_pprof.md)).

## Как запускать

Из корня репозитория (сервер работает, пока не остановите его `Ctrl+C`):

```bash
go run ./1_golang/7_under_the_hood/04_profiling/02_examples/01_pprof_playground
# в другом терминале:
go tool pprof -http=:8081 "http://localhost:8080/debug/pprof/profile?seconds=5"
```

Готовый профиль можно открыть и без запуска сервера:

```bash
go tool pprof -top 1_golang/7_under_the_hood/04_profiling/02_examples/01_pprof_playground/cpu.prof
```

## Workflow

1. Воспроизвести проблему (нагрузка, бенч, прод).
2. Снять профиль (`go tool pprof`, `/debug/pprof/...`).
3. Открыть в `pprof -http=:8081` или `top` в интерактиве.
4. Найти hot path или утечку.
5. Исправить и подтвердить бенчем.

## Дальше

[05_reflection_unsafe](../05_reflection_unsafe/) — reflection и `unsafe`.
