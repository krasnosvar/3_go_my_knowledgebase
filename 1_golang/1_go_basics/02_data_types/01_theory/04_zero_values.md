# Нулевые значения

В Go каждый тип имеет строго определенное нулевое значение. Это значение, которое переменная получает, если объявлена без инициализации.

## Таблица

| Тип       | Нулевое значение | Сравнимый | Изменяемый |
|-----------|------------------|-----------|------------|
| bool      | false            | да        | да         |
| int, uint, float | 0         | да        | да         |
| string    | ""               | да        | нет        |
| array     | массив нулей     | да*       | да         |
| struct    | нулевые поля     | да*       | да         |
| pointer   | nil              | да        | да         |
| slice     | nil              | нет**     | да         |
| map       | nil              | нет**     | да         |
| chan      | nil              | да        | да         |
| func      | nil              | нет**     | нет        |
| interface | nil              | да        | —          |

\* — если все элементы/поля сравнимы.

\*\* — сравнимы только с `nil`.

## Пример

```go
var (
    b   bool
    n   int
    s   string
    p   *int
    sl  []int
    m   map[string]int
    ch  chan int
    fn  func()
    any interface{}
)

fmt.Println(b, n, s, p, sl, m, ch, fn, any)
// false 0 "" <nil> [] map[] <nil> <nil> <nil>
```

## Важные нюансы

- `nil map` можно читать, но нельзя писать — будет panic:

```go
var m map[string]int
_ = m["x"]  // OK -> 0
m["x"] = 1  // panic
```

- `nil slice` безопасен для `len`, `cap`, `range`, `append`:

```go
var s []int
s = append(s, 1) // OK
```

- `nil chan` блокирует всегда — используется намеренно, чтобы отключить ветку `select`.

## Принцип "useful zero value"

Хорошо спроектированные типы в Go работают с нулевым значением без явной инициализации. Пример — `sync.Mutex`:

```go
var m sync.Mutex
m.Lock()
m.Unlock()
```

Это позволяет избегать ненужных конструкторов.
