# Частые ошибки fmt

## 1. Неправильный verb для типа

```go
fmt.Printf("%d", "hello") // %!d(string=hello)
```

Go не падает, но в выводе появляется мусор с пометкой ошибки. Используй `%v`, если не уверен в типе, либо `%T` чтобы увидеть реальный тип.

## 2. Аргументов меньше, чем verb'ов

```go
fmt.Printf("%d %d", 1) // %!d(MISSING)
```

Тоже не падает, просто пишет `MISSING`. Проверяй визуально или `go vet`.

## 3. Забыть `\n`

```go
fmt.Printf("msg=%s", "ok")
fmt.Printf("msg=%s", "err")
// msg=okmsg=err
```

Добавь `\n` или используй `Println`.

## 4. Sprintf на горячем пути

Каждый `Sprintf` аллоцирует и делает reflection. В критичных местах — `strconv`, `strings.Builder`.

## 5. Errorf без %w там, где ждут wrapping

```go
return fmt.Errorf("save: %v", err) // теряет цепочку
```

`errors.Is/As` перестают работать. Используй `%w`, когда хочешь сохранить оригинальную ошибку.

## 6. %s с небезопасными данными

```go
fmt.Printf(userInput) // опасно
```

Если `userInput` содержит `%s`, `%d` — получишь мусор или `MISSING`. Всегда:

```go
fmt.Printf("%s", userInput)
// или
fmt.Print(userInput)
```

Это не SQL-инъекция, но «Format string vulnerability» — можно случайно зализать.

## 7. Использовать `Println` в хот пасе

Println тоже идёт через reflection. В горячем цикле на миллионы итераций — `strings.Builder` быстрее и предсказуемее.

## 8. Забыть о stdout/stderr

```go
fmt.Println("error happened")
```

Ушло в stdout, а пользователь ожидает от CLI: данные в stdout, ошибки в stderr.

```go
fmt.Fprintln(os.Stderr, "error happened")
```

## 9. Сравнение через Sprintf

```go
if fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b) { }
```

Медленно и нестабильно: порядок полей в map не гарантирован. Сравнивай значения напрямую или через `reflect.DeepEqual` (в тестах).
