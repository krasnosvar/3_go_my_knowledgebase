# fmt

Пакет `fmt` — стандартный форматированный ввод/вывод.

## Основные функции

- `fmt.Print(a, b)` — вывод без разделителя (пробелы между non-string аргументами).
- `fmt.Println(a, b)` — то же + `\n` в конце, пробелы между аргументами всегда.
- `fmt.Printf(format, args...)` — с форматом. Куда чаще всего используется.
- `fmt.Sprint` / `Sprintln` / `Sprintf` — то же, но возвращают строку, не пишут.
- `fmt.Fprint` / `Fprintln` / `Fprintf` — пишут в `io.Writer` (файл, сеть, stderr).
- `fmt.Errorf(format, args...)` — возвращает `error`.

## Verbs (глаголы формата)

Самое частое:

```go
fmt.Printf("%d\n", 42)         // 42
fmt.Printf("%s\n", "hello")    // hello
fmt.Printf("%q\n", "hello")    // "hello" — в кавычках, экранированные
fmt.Printf("%v\n", struct{A,B int}{1,2})  // {1 2}
fmt.Printf("%+v\n", struct{A,B int}{1,2}) // {A:1 B:2}
fmt.Printf("%#v\n", struct{A,B int}{1,2}) // struct { A int; B int }{A:1, B:2}
fmt.Printf("%T\n", 42)         // int
fmt.Printf("%f / %.2f\n", 3.14, 3.14)     // 3.140000 / 3.14
fmt.Printf("%x\n", 255)        // ff
fmt.Printf("%b\n", 5)          // 101
fmt.Printf("%c\n", 'A')        // A
fmt.Printf("%p\n", &x)         // 0x...
```

Ширина и точность: `%5d` (ширина 5), `%-5d` (выравнивание влево), `%.3f` (3 знака после точки), `%8.2f` (ширина 8, точность 2).

## Errorf с wrapping (%w)

```go
return fmt.Errorf("save user %d: %w", id, err)
```

`%w` сохраняет цепочку ошибок. Используется с `errors.Is` / `errors.As`. См. `errors` раздел.

## Stringer

Тип с методом `String() string` автоматически форматируется через этот метод в `%v`, `%s`, `Println`:

```go
type Status int
func (s Status) String() string { return []string{"new","active","done"}[s] }
```

## Частые ошибки

- Забыть `\n` в конце `Printf` — строки склеиваются.
- Использовать `Sprintf` там, где хватит `strconv.Itoa` — быстрее и без reflection.
- Печатать user input напрямую через `Printf(s)` — любой `%d` в input ломает вывод. Всегда `Printf("%s", s)` или `Print(s)`.
- Смешивать stdout и stderr в CLI — ошибки должны идти в `os.Stderr` через `Fprintln(os.Stderr, ...)`.

## Когда НЕ использовать fmt

- Hot path с миллионами строк — `strings.Builder` быстрее.
- Структурированные логи — `log/slog`, не `fmt.Printf`.
- Сериализация в стабильный формат — `encoding/json`, не `Sprintf`.
