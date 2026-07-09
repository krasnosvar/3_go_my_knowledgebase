# regexp

Пакет `regexp` — регулярные выражения в Go.

## Подготовка

```go
re := regexp.MustCompile(`^\d{3}-\d{4}$`)   // паникует при ошибке компиляции
re, err := regexp.Compile(`...`)            // если паттерн из пользовательского ввода
```

Компилируй **один раз**, храни результат. Компиляция в горячем цикле — дорого.

## Основные методы

```go
re.MatchString("123-4567")           // true/false
re.FindString("abc 123-4567")        // "123-4567" или ""
re.FindAllString(s, -1)              // все вхождения, -1 = без лимита
re.FindStringIndex(s)                // [start, end]
re.FindStringSubmatch(s)             // основная + группы
re.FindAllStringSubmatch(s, -1)
re.ReplaceAllString(s, "X")
re.ReplaceAllStringFunc(s, strings.ToUpper)
re.Split(s, -1)                      // split по regexp
```

## Группы

```go
re := regexp.MustCompile(`(\w+)@(\w+\.\w+)`)
m := re.FindStringSubmatch("user@example.com")
// m[0] = "user@example.com" (full match)
// m[1] = "user"
// m[2] = "example.com"
```

Именованные группы:

```go
re := regexp.MustCompile(`(?P<user>\w+)@(?P<domain>\w+\.\w+)`)
m := re.FindStringSubmatch("a@b.c")
idx := re.SubexpIndex("user")
u := m[idx]
```

## Синтаксис

Go использует RE2 — без look-behind, без backreferences. Зато без catastrophic backtracking: всегда O(n).

- `.` — любой символ, кроме `\n` (флаг `(?s)` делает многострочно).
- `\d \w \s` — цифра, word, whitespace.
- `^ $` — якоря (флаг `(?m)` меняет поведение для многострок).
- `*? +? ??` — lazy-квантификаторы.
- `(?i)` — case-insensitive.

Проверяй регулярки на [regex101.com](https://regex101.com/) (режим Go/PCRE-lite).

## Производительность

- Компиляция — дорого, результат кэшируй.
- Регулярки медленнее ручных `strings.Contains` / `strings.Index`. Для простых случаев обойдись без regexp.
- RE2 гарантирует линейное время от длины строки, но константы заметные.

## Частые ошибки

- `MustCompile` на пользовательском вводе — panic при кривой регулярке. Используй `Compile` + обработка err.
- Забыть экранировать `.` / `*` / `(` когда нужен литеральный символ.
- Ожидать look-behind / backrefs из PCRE — в Go их нет.
- Использовать regexp для JSON или HTML parsing. Это всегда плохая идея.

## Альтернативы

- Точное совпадение — `strings.Contains/HasPrefix`.
- Сложный парсинг — свой scanner (`bufio.Scanner`, `strings.Fields`).
- Structured data — соответствующий decoder (`encoding/json`, `encoding/xml`).
