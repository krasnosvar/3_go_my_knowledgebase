# fmt, strconv, strings — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Базовые глаголы `fmt` подробнее — в [1_go_basics/03_fmt_and_output](../../1_go_basics/03_fmt_and_output/).

## fmt: функции

Подробно: [01_fmt.md](./01_theory/01_fmt.md).

| Задача | Код |
| :--- | :--- |
| Печать без форматирования | `fmt.Print(a, b)` — пробел только между не-строками |
| Печать + `\n`, пробелы всегда | `fmt.Println(a, b)` |
| Печать по формату | `fmt.Printf("%d\n", n)` |
| Получить строку | `fmt.Sprint` / `Sprintln` / `Sprintf` |
| Писать в `io.Writer` (файл, сеть, stderr) | `fmt.Fprintf(os.Stderr, "err: %v\n", err)` |
| Ошибка с контекстом и цепочкой | `fmt.Errorf("save user %d: %w", id, err)` → `errors.Is/As` |
| Свой вид в `%v`/`%s`/`Println` | метод `String() string` (`fmt.Stringer`) |

## fmt: глаголы

| Глагол | Пример | Вывод |
| :--- | :--- | :--- |
| `%v` / `%+v` / `%#v` | `P{1, "x"}` | `{1 x}` / `{A:1 B:x}` / `main.P{A:1, B:"x"}` |
| `%T` | `P{}` | `main.P` |
| `%d` `%b` `%o` `%x` `%X` | `255` | `255` `11111111` `377` `ff` `FF` |
| `%s` / `%q` | `"hi"` | `hi` / `"hi"` (в кавычках, с экранированием) |
| `%c` / `%U` | `'A'` | `A` / `U+0041` |
| `%f` `%.2f` `%e` `%g` | `1234.5678` | `1234.567800` `1234.57` `1.234568e+03` `1234.5678` |
| `%t` | `true` | `true` |
| `%p` | `&x` | `0xc000...` |
| `%%` | — | `%` |
| Ширина / влево / нули | `%5d` `%-5d` `%05d` | `[   42]` `[42   ]` `[00042]` |
| Ширина + точность | `%8.2f` / `%.3s` | `[    3.14]` / `abc` |
| Индекс аргумента | `Printf("%[2]d %[1]d", 1, 2)` | `2 1` |
| Неверный глагол | `Printf("%d", "oops")` | `%!d(string=oops)` — ⚠️ без ошибки компиляции, ловит `go vet` |

## strconv

Подробно: [02_strconv.md](./01_theory/02_strconv.md).

| Задача | Код | Результат |
| :--- | :--- | :--- |
| string → int | `n, err := strconv.Atoi("42")` | `42` |
| int → string | `strconv.Itoa(42)` | `"42"` |
| С основанием и размером | `strconv.ParseInt("-10", 10, 64)` | `int64` |
| Автоопределение основания | `strconv.ParseInt("0x1f", 0, 64)` | `31` (`0x`, `0o`, `0b`, `0`) |
| Выход за `bitSize` | `strconv.ParseInt("300", 10, 8)` | `value out of range` |
| int → строка в базе | `strconv.FormatInt(255, 16)` | `"ff"` |
| string → float | `strconv.ParseFloat("3.14", 64)` | `3.14` |
| float → string | `strconv.FormatFloat(3.14159, 'f', 2, 64)` | `"3.14"` (`'e'`, `'g'`; точность `-1` = минимум) |
| bool | `strconv.ParseBool("t")` / `FormatBool(true)` | `"1" "t" "true"` → `true` |
| Кавычки | ``strconv.Quote(`a"b`)`` / `Unquote` | `"a\"b"` |
| Разбор ошибки | `var ne *strconv.NumError; errors.As(err, &ne)` | `ne.Num`, `ne.Err` (`strconv.ErrSyntax`, `ErrRange`) |

| Задача | Быстро | Медленно |
| :--- | :--- | :--- |
| int → string | `strconv.Itoa(n)` (без аллокации) | `fmt.Sprintf("%d", n)` (reflection + аллокация) |
| Строка с числом | `"user_" + strconv.Itoa(id)` | `fmt.Sprintf("user_%d", id)` |

## strings

Подробно: [03_strings.md](./01_theory/03_strings.md).

| Задача | Код |
| :--- | :--- |
| Подстрока есть? | `strings.Contains(s, sub)`; любой из символов — `ContainsAny(s, chars)` |
| Префикс / суффикс | `strings.HasPrefix(s, p)` / `HasSuffix(s, x)` |
| Позиция | `strings.Index(s, sub)` (`-1` — нет), `LastIndex` |
| Число вхождений | `strings.Count(s, sub)` (непересекающиеся) |
| Замена | `strings.Replace(s, old, new, n)` (`n=-1` — все), `ReplaceAll` |
| Регистр | `strings.ToLower` / `ToUpper` |
| Обрезка | `TrimSpace`, `Trim(s, cutset)`, `TrimLeft/Right` (набор символов), `TrimPrefix/Suffix` (строка) |
| Разбить | `strings.Split(s, ",")`, `SplitN(s, ",", 2)` |
| По пробелам (сжимая повторы) | `strings.Fields(s)` |
| Склеить | `strings.Join(parts, ", ")` |
| Собрать в цикле | `var b strings.Builder; b.WriteString(w); b.WriteByte(' '); b.String()` |
| Строка как `io.Reader` | `strings.NewReader(s)` (в тестах вместо файла) |
| То же для `[]byte` | пакет `bytes` (`bytes.Contains`, `bytes.Buffer`, …) |
| Проверка символов / обход рун | `unicode.IsLetter(r)`, `for _, r := range s` |

## Ловушки

- ⚠️ `Printf(userInput)` — `%` во вводе ломает вывод; только `Printf("%s", s)` или `Print(s)`.
- ⚠️ Забытый `\n` в `Printf` — строки склеиваются.
- ⚠️ Ошибки CLI — в `os.Stderr` (`Fprintln(os.Stderr, ...)`), данные — в stdout.
- ⚠️ `string(65)` — это `"A"` (руна), а не `"65"`; нужно `strconv.Itoa(65)`. `go vet` предупреждает.
- ⚠️ Игнорировать `err` из `strconv` — получите `0` как «валидное» значение.
- ⚠️ `Atoi` возвращает `int`: на 32-битной платформе это `int32`, для `int64` — `ParseInt(s, 10, 64)`.
- ⚠️ `s += w` в цикле — O(n²) по памяти; `strings.Builder`.
- ⚠️ `strings.Builder` нельзя копировать после использования (panic).
- ⚠️ `strings.Title` устарел — `golang.org/x/text/cases`; `ToUpper` для турецкой `i` — тоже `x/text`.
- ⚠️ `fmt` в hot path и для логов — медленно и неструктурно: `strings.Builder`, `log/slog` (Go 1.21+), `encoding/json`.
