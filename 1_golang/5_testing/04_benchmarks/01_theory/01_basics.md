# Benchmarks: основы

Бенчмарки в Go — встроенный инструмент для измерения производительности. Запускаются `go test -bench`.

## Сигнатура

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = Add(2, 3)
    }
}
```

`b.N` рантайм подбирает сам, начиная с малого и увеличивая, пока тест не выполнится за ~1 секунду.

Результат `_` критичен: иначе компилятор может выкинуть «бесполезный» вызов.

## Запуск

```
go test -bench=. ./...                 # все бенчмарки в пакете
go test -bench=BenchmarkAdd            # один
go test -bench=. -benchmem ./...       # + аллокации
go test -bench=. -benchtime=10s ./...  # лимит времени
go test -bench=. -count=5 ./...        # повторить, чтобы получить ст. отклонение
```

`-benchmem` выводит:

```
BenchmarkAdd-8    100000000    10.5 ns/op    0 B/op    0 allocs/op
```

- `ns/op` — наносекунды на операцию.
- `B/op` — байт на операцию.
- `allocs/op` — аллокаций на операцию.

## ResetTimer / StartTimer / StopTimer

```go
func BenchmarkProcess(b *testing.B) {
    data := loadBigData()  // не учитывать в бенче
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        process(data)
    }
}
```

`ResetTimer` обнуляет счётчик после setup. Альтернативы — `StopTimer/StartTimer` для пропуска отдельных секций.

## SubBenchmarks

```go
func BenchmarkParse(b *testing.B) {
    cases := []string{"small.json", "large.json"}
    for _, c := range cases {
        b.Run(c, func(b *testing.B) {
            data := load(c)
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                _ = Parse(data)
            }
        })
    }
}
```

Каждый случай — отдельная строка в выводе, можно сравнивать.

## RunParallel

```go
func BenchmarkAddParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _ = Add(2, 3)
        }
    })
}
```

Полезно для оценки масштабирования по горутинам.

## Что бенчить

- Hot path — функции, которые вызываются миллионы раз.
- Альтернативы — два варианта реализации, чтобы выбрать.
- До и после оптимизации — измерить эффект.

## Что НЕ бенчить

- Случайный код «на всякий случай» — потерянное время.
- Алгоритм без понимания, что хочешь измерить.
- Бенчмарки на ноутбуке с фоном (Slack, тяжёлый браузер) — шум.

## Правило

- Всегда бенч с `-benchmem` — память часто важнее CPU.
- Сначала **профайл** (`pprof`), потом бенч под подозрительный участок. Бенч без понимания, что узкое место — это микро-оптимизация ради цифр.
