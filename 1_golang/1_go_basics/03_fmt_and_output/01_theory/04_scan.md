# Scan и парсинг ввода

Пакет `fmt` умеет не только выводить, но и читать форматированный ввод.

## Scan, Scanln, Sscan, Fscan

| Функция | Откуда читает |
|---|---|
| `fmt.Scan` | `os.Stdin`, разделитель — пробелы |
| `fmt.Scanln` | `os.Stdin`, до конца строки |
| `fmt.Sscan` | из строки |
| `fmt.Sscanln` | из строки до конца строки |
| `fmt.Fscan` | из `io.Reader` |
| `fmt.Fscanln` | из `io.Reader` до конца строки |

```go
var name string
var age int
fmt.Scan(&name, &age) // из stdin
```

Аргументы — указатели (значения будут записаны туда).

## Scanf с форматом

```go
var port int
fmt.Sscanf("listen 8080", "listen %d", &port)
```

Полезно для простых форматов, но для чего-то сложного лучше `strconv.Atoi`, `strings.Fields`, `regexp` или полноценные парсеры.

## Возвращаемое значение

Все `Scan*` возвращают `(n int, err error)` — сколько элементов успешно прочитано и ошибку. Проверяй:

```go
if _, err := fmt.Scan(&x); err != nil {
    return fmt.Errorf("read x: %w", err)
}
```

## Когда fmt.Scan НЕ подходит

- Большие объёмы/поточное чтение: бери `bufio.Scanner` или `bufio.Reader`.
- CLI-флаги: `flag` / `cobra`.
- Структурированные форматы: `encoding/json`, `encoding/csv`.
- Пользовательский ввод со сложными правилами: самописный парсер или regexp.

## Баг, который ловят часто

`fmt.Scan` разделяет по whitespace, включая перевод строки. Если ожидаешь строку с пробелами — получишь только первое слово. Для строк с пробелами — `bufio.NewReader(os.Stdin).ReadString('\n')`.

## Рекомендация

`fmt.Scan*` удобны для учебных задач и быстрых CLI-прототипов. В боевом коде почти всегда лучше `bufio.Scanner` + свой парсер или готовая библиотека.
