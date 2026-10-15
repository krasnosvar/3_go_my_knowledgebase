# Оптимизация по памяти

## С чего начать

Не оптимизировать без данных. Сначала:

1. `pprof heap` — где аллоцируется.
2. `pprof allocs` — что было выделено за период.
3. `GODEBUG=gctrace=1` — частота и длительность GC.
4. Бенч с `-benchmem` — `B/op`, `allocs/op`.

Если ничего из этого не указывает на проблему — оптимизация память не нужна.

## Частые источники аллокаций

### 1. Конверсии string ↔ []byte

```go
data := []byte(s)  // copy
s := string(data)  // copy back
```

В hot path → используй `bytes.Buffer` / `strings.Builder`, или unsafe-конверсии (с большим вниманием).

### 2. fmt в hot loop

```go
key := fmt.Sprintf("user_%d", id)
```

Замени на `"user_" + strconv.Itoa(id)`. Быстрее в десятки раз.

### 3. Слайсы без cap

```go
res := []int{}
for i := 0; i < n; i++ { res = append(res, i) }
```

Делай `make([]int, 0, n)`. Одна аллокация вместо O(log n) переаллокаций.

### 4. interface{} / any

Любая упаковка значения в interface — аллокация box'а. Это часто скрытый источник.

### 5. Замыкания в hot path

```go
data.Process(func(x int) int { return x * mult })
```

Переменные, захваченные замыканием, могут escape'ить в heap. Профайл покажет.

## sync.Pool

Переиспользование коротко-живущих объектов:

```go
var bufPool = sync.Pool{
    New: func() any { return new(bytes.Buffer) },
}

b := bufPool.Get().(*bytes.Buffer)
defer bufPool.Put(b)
b.Reset()
```

Эффективен для:

- буферов сериализации;
- временных слайсов в парсерах;
- per-request объектов в HTTP-сервисе.

Не панацея: GC может опустошить пул в любой момент.

## Объекты с явным lifetime

Если знаешь, что объект жив весь запуск приложения — выделяй один раз глобально, не в каждой функции. Это снижает давление на GC и помогает кэшу.

## Уменьшение размера структур

```go
type Bad struct {
    a bool   // 1 байт + 7 padding
    b int64  // 8
    c bool   // 1 + 7
}             // итого 24

type Good struct {
    b int64  // 8
    a bool   // 1
    c bool   // 1 + 6 padding
}             // итого 16
```

Для миллионов структур разница в layout даёт реальные мегабайты.

## GOMEMLIMIT в Kubernetes

```
# cgroup v2: лимит контейнера в байтах ("max" — лимита нет)
GOMEMLIMIT=$(awk '{printf "%dB", $1*0.85}' /sys/fs/cgroup/memory.max)
```

Помогает GC не превысить pod-лимит и избежать OOM kill. `/proc/meminfo` здесь не годится: внутри контейнера `MemTotal` — память всего узла, а не лимит pod'а. Альтернатива — передать `limits.memory` в env через Downward API (`resourceFieldRef`).

## Правило

- Без pprof не оптимизируй — будешь тратить время впустую.
- 0 alloc/op в hot path — реальная и достижимая цель.
- Памятная экономия даёт более стабильный latency, не только меньшее потребление.
