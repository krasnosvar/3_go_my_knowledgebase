# Sprintf и Errorf

## Sprintf — форматирование в строку

Та же семантика, что у `Printf`, но возвращает строку:

```go
s := fmt.Sprintf("user=%s status=%d", name, code)
```

Куда применять:

- сборка сообщений пользователю (когда формат сложный);
- готовим ключ/метку для логов, метрик;
- тесты: ожидаемое значение.

Для горячего пути и большого объёма сборки — используй `strings.Builder`: `Sprintf` аллоцирует.

## Fprintf — в io.Writer

```go
fmt.Fprintf(os.Stderr, "error: %v\n", err)
```

Работает с любым `io.Writer`: файл, сетевой сокет, HTTP `w http.ResponseWriter`. В бою чаще встретишь именно `Fprintf`, а не `Printf`.

## Println vs Printf

```go
fmt.Println("a", 1, true) // "a 1 true\n"
```

`Println` сам добавляет пробелы и `\n`. Удобно для простых отладочных строк и tests.

## Errorf — форматированная ошибка

```go
return fmt.Errorf("save user %d: %w", id, err)
```

Возвращает `error`. С `%w` поддерживает wrapping (см. `errors/02_wrapping_is_as.md`).

Правила:

- одна `%w` на вызов;
- текст — технический контекст, а не «копия err».

## Зачем различать Sprintf / Errorf

`Errorf` возвращает уже `error`, без ручного `errors.New(fmt.Sprintf(...))`. Это короче и поддерживает `%w`.

## Частая ошибка: забыть `\n`

У `Printf`/`Fprintf` перевод строки не добавляется автоматически. В логах это даёт склеенные строки. Либо `\n` в конце формата, либо используй `Println`.

## Бенч

- `Sprintf` — ~100 нс и аллокация.
- `strconv.Itoa` / `strings.Builder` для конкретных случаев — быстрее и без reflection.

В обычном CRUD это неважно, но в hot path иногда принципиально.
