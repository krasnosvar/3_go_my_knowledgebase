# strings

Пакет `strings` — утилиты для работы со строками. Нужно знать хорошо: используется постоянно.

## Поиск и проверка

```go
strings.Contains(s, sub)       // есть ли подстрока
strings.ContainsAny(s, chars)  // есть ли любой из символов
strings.HasPrefix(s, pre)
strings.HasSuffix(s, suf)
strings.Index(s, sub)          // -1 если не найдено
strings.LastIndex(s, sub)
strings.Count(s, sub)          // количество непересекающихся вхождений
```

## Модификация (возвращают новую строку)

```go
strings.Replace(s, old, new, n)       // n=-1 заменить все
strings.ReplaceAll(s, old, new)
strings.ToLower(s) / strings.ToUpper(s)
strings.Title(s)                      // DEPRECATED: golang.org/x/text/cases
strings.TrimSpace(s)
strings.Trim(s, cutset)
strings.TrimLeft / TrimRight / TrimPrefix / TrimSuffix
```

## Разделение / склейка

```go
parts := strings.Split(s, ",")        // []string
parts = strings.SplitN(s, ",", 2)     // максимум 2 части
fields := strings.Fields(s)           // по whitespace, множественные сжимаются

joined := strings.Join(parts, ", ")
```

`Fields` удобнее `Split(s, " ")` для обработки пользовательского ввода с лишними пробелами.

## strings.Builder

Эффективная сборка строки без лишних аллокаций:

```go
var b strings.Builder
for _, w := range words {
    b.WriteString(w)
    b.WriteByte(' ')
}
result := b.String()
```

Намного быстрее `+` или `fmt.Sprintf` в цикле. Правила:

- Не копируй Builder (там unsafe-магия).
- После `String()` Builder ещё можно использовать; `Reset()` обнулит.
- Для одноразовых операций `Sprintf` проще, в цикле — Builder.

## Reader

```go
r := strings.NewReader("hello")
// r реализует io.Reader, io.Seeker — удобно подсунуть в функции, ждущие io.Reader
```

Частый паттерн в тестах: `bytes.NewBuffer`/`strings.NewReader` вместо реального файла.

## Unicode и руны

Все функции работают с байтами UTF-8. Для corrupt-safe обработки, проверок символов — пакет `unicode` (`unicode.IsLetter`, `IsDigit`), для обхода рун — `for _, r := range s`. См. раздел `1_go_basics/strings`.

## Частые ошибки

- `strings.Title` — deprecated (плохо ведёт себя с Unicode). Используй `golang.org/x/text/cases`.
- `strings.ToUpper` на локализованных данных (турецкий i) может давать неожиданные результаты — для глобальных приложений смотри `x/text`.
- Собирать длинную строку через `+=` в цикле — O(n²) по памяти, Builder обязателен.

## Bytes для работы с []byte

Пакет `bytes` — зеркало `strings` для `[]byte`. Тот же API, но для mutable-байтовых буферов. Актуально в сетевом коде.
