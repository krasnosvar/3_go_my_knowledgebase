# Use-cases: где гонки встречаются чаще всего

## 1. Счётчики и агрегаты

```go
var count int
go func() { count++ }()
go func() { count++ }()
```

Самая частая гонка. Fix: `atomic.Int64` или mutex.

## 2. Общий map

```go
go func() { m["a"] = 1 }()
go func() { _ = m["a"] }()
```

Не просто гонка — runtime, обнаружив одновременную запись/чтение map, завершает процесс с `fatal error: concurrent map writes` (не паника: `recover` не перехватит; обнаружение не гарантировано). Fix: mutex, `sync.Map` для специфических сценариев.

## 3. Флаг «работа закончилась»

```go
var done bool
go func() { done = true }()
for !done { /* ждём */ }
```

Без синхронизации цикл может крутиться вечно. Fix: `atomic.Bool`, `context`, канал.

## 4. Ленивая инициализация

```go
if client == nil {
    client = New()
}
```

Два запроса одновременно создадут двух клиентов. Fix: `sync.Once`, `atomic.Pointer` + CAS.

## 5. Переменная цикла в горутине

Главный баг «до Go 1.22», всё ещё встречается. Fix: параметр или локальная копия.

## 6. HTTP-handler с общим state

```go
var counter int // пакетная переменная

func handler(w http.ResponseWriter, r *http.Request) {
    counter++
}
```

Каждый handler вызывается в своей горутине. Fix: atomic или mutex.

## 7. Кэш без синхронизации

```go
func Get(key string) string {
    if v, ok := cache[key]; ok { return v }
    v := compute(key)
    cache[key] = v
    return v
}
```

Несколько Get'ов одновременно → гонка, двойные вычисления. Fix: mutex + double-check или `singleflight`.

## 8. Логгер/метрики с shared state

Глобальный слайс в логгере, в который пишут все. Часто встречается. Fix: передавать через канал в логгер-горутину или lock'и.

## 9. Общий buffer для ответа

```go
var buf bytes.Buffer
go func() { buf.Write(a) }()
go func() { buf.Write(b) }()
```

`bytes.Buffer` не потокобезопасен. Fix: `sync.Mutex` + Buffer, или отдельные buffer'ы → merge.

## 10. Тестовый хелпер с глобальной переменной

Тесты, которые делятся `var testData`. Если тесты параллельные (`t.Parallel()`) — гонка. Fix: локальные данные в каждом тесте.

## Правило сканирования

При ревью или при дебаге рано или поздно задавай вопрос:

> «Могут ли эти строки выполниться из двух горутин одновременно? Кто из них пишет?»

Если да → нужен primitive или redesign.
