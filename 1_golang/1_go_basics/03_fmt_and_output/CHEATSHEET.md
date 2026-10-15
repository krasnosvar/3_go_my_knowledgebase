# fmt и вывод — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.

## Функции

Подробно: [01_fmt_basics.md](./01_theory/01_fmt_basics.md), [03_sprintf_errorf.md](./01_theory/03_sprintf_errorf.md).

| Куда | Без формата | С пробелами и `\n` | По шаблону |
| :--- | :--- | :--- | :--- |
| `os.Stdout` | `fmt.Print(a, b)` | `fmt.Println(a, b)` | `fmt.Printf(format, a, b)` |
| Любой `io.Writer` (файл, `os.Stderr`, `http.ResponseWriter`) | `fmt.Fprint(w, ...)` | `fmt.Fprintln(w, ...)` | `fmt.Fprintf(w, format, ...)` |
| Вернуть `string` | `fmt.Sprint(...)` | `fmt.Sprintln(...)` | `fmt.Sprintf(format, ...)` |
| Вернуть `error` | — | — | `fmt.Errorf(format, ...)`, `%w` — обернуть |

| Поведение | Пример → вывод |
| :--- | :--- |
| `Print` ставит пробел только между двумя не-строками | `fmt.Print("a", "b", 1, 2)` → `ab1 2` |
| `Println` — пробел между всеми + `\n` | `fmt.Println("a", "b", 1, 2)` → `a b 1 2` |
| `Printf`/`Fprintf` не добавляют `\n` | `fmt.Printf("n=%d\n", 42)` |
| Данные → stdout, диагностика → stderr | `fmt.Fprintln(os.Stderr, "warn:", w)` |
| Тип с методом `String() string` (`fmt.Stringer`) | `%v`, `%s`, `Println` вызывают `String()` |
| Тип с `Error() string` | печатается через `Error()` |

## Глаголы

Подробно: [02_printf_verbs.md](./01_theory/02_printf_verbs.md).

| Verb | Что | Пример → вывод |
| :--- | :--- | :--- |
| `%v` | значение по умолчанию | `{1 den}` |
| `%+v` | struct с именами полей | `{ID:1 Name:den}` |
| `%#v` | Go-синтаксис | `main.User{ID:1, Name:"den"}` |
| `%T` | тип | `main.User` |
| `%%` | литерал `%` | `50%` |
| `%d` | целое десятичное | `42` |
| `%b` / `%o` / `%O` | двоичное / восьмеричное / с `0o` | `101` / `10` / `0o10` |
| `%x` / `%X` / `%#x` | hex | `ff` / `FF` / `0xff` |
| `%c` / `%U` | руна как символ / Unicode | `A` / `U+0041` |
| `%f` / `%F` | float | `3.141590` |
| `%e` / `%E` | научная нотация | `1.234568e+03` |
| `%g` / `%G` | компактный float | `1e-06` |
| `%s` | строка, `[]byte` | `hi` |
| `%q` | строка в Go-кавычках, экранированная | `"hi\n"` |
| `%x` для строки / `% x` | байты в hex / с пробелами | `6869` / `68 69` |
| `%t` | bool | `true` |
| `%p` | указатель | `0xc000012345` |

## Ширина, точность, индексы

Подробно: [02_printf_verbs.md](./01_theory/02_printf_verbs.md).

| Формат | Аргумент | Вывод |
| :--- | :--- | :--- |
| `%5d` | `42` | `   42` |
| `%-5d` | `42` | `42   ` |
| `%05d` | `42` | `00042` |
| `%.2f` | `3.14159` | `3.14` |
| `%8.2f` | `3.14` | `    3.14` |
| `%10.5s` | `"hello world"` | `     hello` (не больше 5 символов) |
| `%-10s` | `"go"` | `go        ` — колонки таблиц |
| `%[2]d %[1]d` | `1, 2` | `2 1` |
| `%[1]d + %[1]d = %[2]d` | `2, 4` | `2 + 2 = 4` |

## Errorf и %w

Подробно: [03_sprintf_errorf.md](./01_theory/03_sprintf_errorf.md).

| Задача | Код |
| :--- | :--- |
| Обернуть с контекстом | `fmt.Errorf("save user %d: %w", id, err)` |
| Проверить обёрнутую | `errors.Is(err, ErrNotFound)`, `errors.As(err, &target)` |
| Обернуть несколько (1.20+) | `fmt.Errorf("%w; %w", err1, err2)` |
| Только текст, без цепочки | `%v` — `errors.Is/As` уже не найдут исходную |
| Вместо `errors.New(fmt.Sprintf(...))` | `fmt.Errorf(...)` |

## Scan

Подробно: [04_scan.md](./01_theory/04_scan.md).

| Функция | Источник | Разделитель |
| :--- | :--- | :--- |
| `fmt.Scan(&a, &b)` | `os.Stdin` | любые пробелы, включая `\n` |
| `fmt.Scanln(&a)` | `os.Stdin` | до конца строки |
| `fmt.Sscan(s, &a)` / `fmt.Sscanln` | строка | пробелы / до `\n` |
| `fmt.Sscanf(s, "listen %d", &port)` | строка | по шаблону |
| `fmt.Fscan(r, &a)` / `fmt.Fscanln` | `io.Reader` | пробелы / до `\n` |
| Результат | `(n int, err error)` — проверять `err` | — |

| Вместо `fmt.Scan` | Когда |
| :--- | :--- |
| `bufio.Scanner` / `bufio.Reader` | большие объёмы, построчно |
| `bufio.NewReader(os.Stdin).ReadString('\n')` | строка с пробелами |
| `flag` / `cobra` | CLI-флаги |
| `encoding/json`, `encoding/csv` | структурированные форматы |

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

type User struct {
	ID   int
	Name string
}

type Status int

func (s Status) String() string { return [...]string{"pending", "active"}[s] }

var ErrNotFound = errors.New("not found")

func main() {
	u := User{ID: 1, Name: "den"}
	fmt.Printf("%v | %+v | %#v | %T\n", u, u, u, u)
	// {1 den} | {ID:1 Name:den} | main.User{ID:1, Name:"den"} | main.User
	fmt.Printf("[%5d] [%-5d] [%05d] [%.2f] [%8.2f] [%10.5s]\n", 42, 42, 42, 3.14159, 3.14, "hello world")
	// [   42] [42   ] [00042] [3.14] [    3.14] [     hello]
	fmt.Printf("%b %o %x %X %#x %c %U %q %e\n", 5, 8, 255, 255, 255, 'A', 'A', "hi\n", 1234.5678)
	// 101 10 ff FF 0xff A U+0041 "hi\n" 1.234568e+03
	fmt.Printf("%[2]d %[1]d %d\n", 1, 2)               // 2 1 2 — после %[1] счёт идёт с аргумента 2
	fmt.Printf("%v %d | 50%%\n", Status(1), Status(1)) // active 1 | 50%

	err := fmt.Errorf("load user %d: %w", u.ID, ErrNotFound)
	fmt.Println(err, errors.Is(err, ErrNotFound)) // load user 1: not found true

	var port int
	n, err := fmt.Sscanf("listen 8080", "listen %d", &port)
	fmt.Println(n, err, port) // 1 <nil> 8080

	fmt.Fprintln(os.Stderr, "warn: диагностика идёт в stderr")
}
```

## Когда не fmt

Подробно: [01_fmt_basics.md](./01_theory/01_fmt_basics.md#4-о-производительности), [06_use_cases.md](./01_theory/06_use_cases.md).

| Задача | Вместо `fmt` |
| :--- | :--- |
| Склейка двух строк | `a + b` |
| Сборка строки из многих частей / в цикле | `strings.Builder` |
| Число в строку на горячем пути | `strconv.Itoa`, `strconv.AppendInt` |
| Логи в проде | `log/slog` (1.21+), `zap`, `zerolog` |
| Отладка в тестах | `t.Logf("got %+v, want %+v", got, want)` |
| Сериализация | `encoding/json`, `encoding/xml` |

## Ловушки

- ⚠️ Неверный verb не паникует: `fmt.Printf("%d", "hello")` → `%!d(string=hello)`.
- ⚠️ Аргументов меньше, чем глаголов: `%!d(MISSING)`; лишние — `%!(EXTRA ...)`. Ловит `go vet`.
- ⚠️ Забытый `\n` в `Printf` склеивает строки логов.
- ⚠️ `fmt.Printf(userInput)` — format string vulnerability; писать `fmt.Print(userInput)` или `Printf("%s", userInput)`.
- ⚠️ `%v` вместо `%w` в `Errorf` рвёт цепочку для `errors.Is/As`.
- ⚠️ После `%[n]` следующие глаголы без индекса берут аргументы `n+1`, `n+2`, …
- ⚠️ `fmt.Scan` в `string` читает только первое слово.
- ⚠️ Ошибки CLI в stdout ломают пайпы — диагностика в `os.Stderr`.
- ⚠️ `Sprintf`/`Println` в горячем цикле: упаковка в `any` (аллокации) + рефлексия.
- ⚠️ Сравнивать значения через `Sprintf("%v")` медленно — сравнивать напрямую или `reflect.DeepEqual` в тестах.
