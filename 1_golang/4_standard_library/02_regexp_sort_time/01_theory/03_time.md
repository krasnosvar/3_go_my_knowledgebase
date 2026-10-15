# time

Пакет `time` — работа с датами, временем, длительностями, таймерами.

## Типы

- `time.Time` — конкретный момент.
- `time.Duration` — длительность (например, `5 * time.Second`).
- `time.Location` — часовой пояс.
- `time.Ticker`, `time.Timer` — примитивы тайминга.

## Основные операции

```go
now := time.Now()
later := now.Add(10 * time.Second)
diff := later.Sub(now)              // 10s

if now.Before(later) { /* ... */ }

ts := now.Unix()                    // Unix seconds
ms := now.UnixMilli()
```

## Форматирование и парсинг

Go использует референсный момент `2006-01-02 15:04:05 MST` — запомнить его обязательно.

```go
s := now.Format(time.RFC3339)              // "2026-04-17T15:04:05+03:00"
t, err := time.Parse("2006-01-02", "2026-04-17")
t, err = time.ParseInLocation("2006-01-02", "2026-04-17", time.Local)
```

Популярные константы: `time.RFC3339`, `time.RFC1123`, `time.DateTime` (Go 1.20+), `time.DateOnly`, `time.TimeOnly`.

## Duration

```go
d := 2 * time.Second + 500 * time.Millisecond
d.Milliseconds()  // 2500
time.Since(t)     // сколько прошло с t
time.Until(t)     // сколько осталось до t
```

## Sleep, Timer, Ticker

```go
time.Sleep(d) // блокирующая пауза

t := time.NewTimer(d)
defer t.Stop()
<-t.C          // одноразовое срабатывание

tk := time.NewTicker(d)
defer tk.Stop()
for {
    select {
    case <-tk.C:  // периодическое
    case <-ctx.Done():
        return
    }
}
```

**Всегда останавливай `Timer`/`Ticker` через `Stop()`** (в defer). До Go 1.23 не остановленные таймеры не собирались GC — утечка. С Go 1.23 GC собирает таймеры без ссылок, но не остановленный `Ticker`, на который есть ссылка, продолжает тикать.

## time.After

`time.After(d)` возвращает канал — одноразовый таймер. Удобно в select:

```go
select {
case v := <-ch:
case <-time.After(2*time.Second):
    return ErrTimeout
}
```

В цикле создаёт новый таймер каждый раз — в hot path предпочти `time.NewTimer`.

## Monotonic vs Wall

`time.Now()` возвращает и монотонное, и «настенное» время. Для измерения длительности — `time.Since` / `Until` — они правильно используют монотонную часть.

Изменение системного времени (NTP, летнее время) не ломает `time.Since`.

## Часто ошибки

- Форматировать через свои strings.Sprintf — ошибки преобразования. Используй `Format` / `Parse`.
- Layout с неправильными цифрами (`2002-01-02` вместо `2006-01-02`) — типовой баг при первой попытке.
- Забыть `Stop()` на Timer/Ticker в долгоживущем коде.
- Сравнивать через `==`: есть monotonic clock, лучше `t1.Equal(t2)`.
- `time.Time{}` — это не «нулевое время Unix», а год 1. Проверка «не инициализирован» — `t.IsZero()`.
