# Use-cases atomic

## 1. Счётчики метрик

```go
var requests atomic.Int64

func handle() {
    requests.Add(1)
}
```

Дешевле mutex'а, подходит для очень частых инкрементов.

## 2. Флаги состояния

```go
var ready atomic.Bool

func markReady() { ready.Store(true) }
func isReady() bool { return ready.Load() }
```

Простое «включено/выключено» без синхронизации.

## 3. Лениво инициализируемое состояние

```go
var current atomic.Pointer[Snapshot]

func Get() *Snapshot {
    if s := current.Load(); s != nil {
        return s
    }
    s := build()
    if !current.CompareAndSwap(nil, s) {
        return current.Load()
    }
    return s
}
```

Типичная lock-free-инициализация без mutex.

## 4. Copy-on-write snapshot

Конфиг/кэш обновляется целиком, читатели всегда работают с согласованным snapshot'ом:

```go
var cfg atomic.Pointer[Config]

// на обновлении
new := loadCopy()
cfg.Store(new)

// на чтении
c := cfg.Load()
use(c)
```

Читателям не нужен mutex, они не блокируют друг друга.

## 5. Версионирование и генерации

```go
var gen atomic.Uint64
gen.Add(1) // при каждом обновлении
```

Используется как «эпоха» для invalidation: кеш сравнивает `gen` с сохранённым значением.

## 6. Graceful shutdown флаг

```go
var shuttingDown atomic.Bool

func acceptNew() {
    if shuttingDown.Load() { return }
    // ...
}
```

Писатель и читатели работают через atomic, без блокировок.

## 7. Статистика «последнее значение»

```go
var lastLatency atomic.Int64 // ns

func record(d time.Duration) { lastLatency.Store(d.Nanoseconds()) }
```

Не нужна история, только последнее замеренное значение — atomic подходит.

## Когда НЕ atomic

- Нужно обновить несколько связанных полей → mutex или один `atomic.Pointer` на структуру.
- Нужны условные обновления с побочными эффектами — mutex.
- Нужно мониторить полную карту → mutex + map или `sync.Map`.
- Нужна синхронизация по событиям (ждать результат) → каналы.
