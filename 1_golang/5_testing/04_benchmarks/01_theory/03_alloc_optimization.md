# Оптимизация по аллокациям

В Go аллокации — главный источник лагов под нагрузкой: каждое выделение давит на GC, создаёт паузы, увеличивает RSS.

## Признаки проблемы

В бенче с `-benchmem`:

```
50000000     85.2 ns/op    32 B/op    2 allocs/op
```

В hot path даже 1 allocs/op на сотнях тысяч RPS = миллионы аллокаций в секунду.

## Частые источники

### 1. Преобразование string ↔ []byte

```go
b := []byte(s)
s := string(b)
```

Каждое — копия. В hot path → `strings.Builder` или работа в одной форме.

### 2. fmt.Sprintf

```go
s := fmt.Sprintf("user_%d", id) // аллокация + reflection
```

Альтернатива:

```go
s := "user_" + strconv.Itoa(id)
```

Или `strings.Builder`.

### 3. Слайсы без cap

```go
res := []int{}
for i := 0; i < 1000; i++ { res = append(res, i) }
```

`append` будет переаллоцировать массив несколько раз. Знаешь размер — выделяй сразу:

```go
res := make([]int, 0, 1000)
```

### 4. Замыкания в hot loop

Закрытие переменной в анонимной функции часто escape'ится в heap. `go build -gcflags="-m"` показывает escape analysis.

### 5. Интерфейсы

Передача значения через `interface{}` или `any` — аллокация box'а. Generics обычно решают.

## sync.Pool

Переиспользование короткоживущих объектов:

```go
var bufPool = sync.Pool{
    New: func() any { return new(bytes.Buffer) },
}

b := bufPool.Get().(*bytes.Buffer)
defer bufPool.Put(b)
b.Reset()
// ...
```

Полезно для буферов в JSON-сериализации, парсерах, сетевых обработчиках.

Pool не гарантия наличия объекта — GC может его очистить. Это «помощь», а не хранилище.

## escape analysis

```
go build -gcflags="-m" ./...
```

Вывод покажет, какие переменные ушли в heap:

```
./main.go:10:6: moved to heap: x
```

Если функция возвращает `*T`, переменная ушла в heap. Можно проверить, действительно ли это нужно.

## strings.Builder vs []byte vs +=

```go
// плохо: O(n²) аллокаций
s := ""
for _, w := range words { s += w }

// хорошо: 1-2 аллокации
var b strings.Builder
for _, w := range words { b.WriteString(w) }
return b.String()
```

## Правило

- Сначала `pprof` (см. 7_under_the_hood/04_profiling), потом оптимизация.
- 0 allocs/op в горячем пути — реальная цель.
- Не оптимизируй вне hot path: тратишь время, ломаешь читаемость.
