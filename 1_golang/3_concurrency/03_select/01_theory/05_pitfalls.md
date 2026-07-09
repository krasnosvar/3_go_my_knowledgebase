# Частые ошибки select

## 1. Busy loop с default

```go
for {
    select {
    case v := <-ch:
        use(v)
    default:
        // крутится постоянно, 100% CPU
    }
}
```

Либо уберите `default`, либо добавьте паузу / ticker / `ctx.Done()` ветку.

## 2. Забытое условие выхода

```go
for {
    select {
    case job := <-jobs:
        process(job)
    }
}
```

Нет `ctx.Done()` — горутина живёт вечно. При остановке сервиса — утечка.

## 3. break внутри select

```go
for {
    select {
    case <-done:
        break // выход ИЗ SELECT, не из for
    case v := <-ch:
        use(v)
    }
}
```

Для выхода из внешнего цикла нужно либо `return`, либо label:

```go
Loop:
for {
    select {
    case <-done:
        break Loop
    }
}
```

## 4. Чтение из nil-канала

```go
var ch chan int
select {
case v := <-ch: // блок навсегда
    use(v)
}
```

Полезный паттерн, когда хотим «отключить» ветку select, подставляя nil-канал. Но без понимания — это deadlock.

## 5. Одна и та же ветка готова много раз

Если в select несколько веток выбираются «одинаково часто», но одна всегда стоит в очереди — порядок случайный. Не рассчитывай на приоритет.

## 6. Потеря данных через default send

```go
select {
case ch <- v:
default: // тихо дропнули
}
```

Для метрик — ок, для бизнес-данных — нет.

## 7. time.After утечка

В hot цикле `time.After` создаёт таймер на каждой итерации, даже если не нужен:

```go
for {
    select {
    case v := <-ch:
        use(v)
    case <-time.After(5 * time.Second):
        // таймер создаётся заново каждый раз
    }
}
```

Лучше один `time.NewTimer` + `Reset`.

## 8. select без веток

```go
select {}
```

Блокирует навсегда. Изредка применяют, чтобы «удержать main» — но почти всегда есть более явная альтернатива (`<-ctx.Done()`).

## 9. Отправка с таймером без cancel

Проверяй, не висит ли writer, если reader ушёл по таймауту. Buffered-канал + ctx — обычно решают.

## 10. Забыть `cancel()` у контекста

После `context.WithTimeout` / `WithCancel` обязательно `defer cancel()`. Без этого контекст висит в куче до срабатывания таймера.
