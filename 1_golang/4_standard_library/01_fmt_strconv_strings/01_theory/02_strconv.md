# strconv

Пакет `strconv` — конвертация между строками и примитивами (int, float, bool).

## Целые

```go
n, err := strconv.Atoi("42")            // int, удобная обёртка над ParseInt
s := strconv.Itoa(42)                   // "42"

n64, err := strconv.ParseInt("-10", 10, 64) // base=10, bitSize=64
u, err := strconv.ParseUint("0xff", 0, 64)  // base=0 — авторазбор: 0x, 0, 0b
s = strconv.FormatInt(n64, 16)          // "ff"
```

`base=0` — самый удобный: поддерживает `0x`, `0b`, `0o` префиксы. `bitSize` ограничивает диапазон результата.

## Float

```go
f, err := strconv.ParseFloat("3.14", 64)
s := strconv.FormatFloat(3.14159, 'f', 2, 64) // "3.14"
```

Формат: `'f'` — обычный, `'e'` — научный, `'g'` — авто, `-1` как точность — минимум нужных знаков.

## Bool

```go
b, err := strconv.ParseBool("true")   // "1","t","true" → true; "0","f","false" → false
s := strconv.FormatBool(true)         // "true"
```

## Quoting

```go
strconv.Quote(`line with "quote"`) // `"line with \"quote\""`
strconv.Unquote(`"hello\n"`)       // `hello\n` (раскодированное)
```

Полезно для логов, CLI-вывода, безопасного представления строки.

## Возвращаемая ошибка

`strconv` возвращает `*strconv.NumError`:

```go
_, err := strconv.Atoi("abc")
var ne *strconv.NumError
if errors.As(err, &ne) {
    fmt.Println(ne.Num, ne.Err) // abc invalid syntax
}
```

Всегда проверяй err: на кривом входе получишь 0 без понимания того, что парсинг провалился.

## Почему не через fmt.Sprintf/Sscan

- `strconv.Itoa(n)` — ~2 нс, без аллокации.
- `fmt.Sprintf("%d", n)` — ~60 нс, аллокация.
- Для hot path — прямой `strconv` предпочтителен.

## Частые ошибки

- Использовать `Atoi` для `int64` на 32-битной платформе — `int == int32` там, может переполниться.
- Конвертировать `int -> string` через `string(65)`. Это `"A"` (rune), а не `"65"`. Всегда `strconv.Itoa`.
- Игнорировать `err`. Тогда нулевое значение проглотится как валидное.
