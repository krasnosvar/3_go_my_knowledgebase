# Pipeline

Pipeline — цепочка стадий, соединённых каналами. Каждая стадия читает из входного канала и пишет в выходной.

## Структура

```
source -> [stage1] -> [stage2] -> [stage3] -> sink
```

Каждая стадия — отдельная горутина.

## Пример

```go
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for v := range in {
            out <- v * v
        }
    }()
    return out
}

func main() {
    for v := range square(generate(1, 2, 3, 4)) {
        fmt.Println(v) // 1 4 9 16
    }
}
```

## Правила

1. **Каждая стадия закрывает свой выходной канал.**
2. **Range по входу** — значит, выход закроется сам, когда вход закрылся.
3. **Directional channels** в сигнатурах (`<-chan`, `chan<-`) — защита от ошибок.

## Почему хорошо

- Естественный backpressure через блокировку каналов.
- Легко добавлять/убирать стадии.
- Каждая стадия тестируется отдельно.

## Обогащение

### Контекст для отмены

```go
func square(ctx context.Context, in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for v := range in {
            select {
            case <-ctx.Done():
                return
            case out <- v * v:
            }
        }
    }()
    return out
}
```

Пайплайн должен уметь сворачиваться по сигналу.

### Параллельная стадия

Если стадия тяжёлая — запусти N воркеров и merge (fan-out + fan-in).

### Буферизованные каналы для smoothing

Если стадии работают с разной скоростью, буферизованные каналы сглаживают это.

## Ошибки

- Стадия не закрывает выход → следующая стадия «зависает».
- Нет ctx.Done() → при отмене stages продолжают работать и блокируются.
- Ошибки на стадии не имеют единого канала для передачи. Обычно: отдельный `errCh` или `errgroup`.

## Когда pipeline

- ETL-обработка.
- Streaming / media-проcessing.
- Любая «преобразовательная» последовательность шагов.

## Когда НЕ pipeline

- Очень короткие операции — overhead каналов превысит профит.
- Операции с сильным разнобоем времени — лучше worker pool на каждую.
