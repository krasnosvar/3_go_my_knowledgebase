# encoding/json, csv, base64 — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Версии указаны там, где поведение зависит от них.

## JSON: API

Подробно: [01_json.md](./01_theory/01_json.md).

| Задача | Код |
| :--- | :--- |
| Структура → JSON | `data, err := json.Marshal(u)` |
| JSON → структура | `err := json.Unmarshal(data, &u)` — ⚠️ указатель |
| С отступами (CLI, отладка, файлы) | `json.MarshalIndent(u, "", "  ")` |
| Читать поток (файл, HTTP body) | `dec := json.NewDecoder(r); dec.Decode(&v)`; конец потока — `io.EOF` |
| Запретить неизвестные поля | `dec.DisallowUnknownFields()` |
| Писать поток | `enc := json.NewEncoder(w); enc.SetIndent("", "  "); enc.Encode(u)` |
| Отложить разбор части | поле `Data json.RawMessage`, потом `json.Unmarshal(env.Data, &x)` |
| Свой формат типа | `MarshalJSON() ([]byte, error)` / `(*T) UnmarshalJSON([]byte) error` |

## JSON: теги

| Тег | Эффект |
| :--- | :--- |
| `json:"name"` | имя поля в JSON |
| `json:"age,omitempty"` | пропустить `0`, `""`, `false`, `nil`, пустой slice/map |
| `json:"created,omitzero"` | пропустить нулевое значение, в т.ч. `time.Time{}` и структуры (Go 1.24+) |
| `json:"-"` | игнорировать поле |
| `json:"id,string"` | число как строка: `"id":"1152921504606846976"` (большие int64 для JavaScript) |
| поле с маленькой буквы | ⚠️ не сериализуется вообще (не экспортировано) |

## JSON → `any`

| JSON | Go-тип при декодировании в `any` |
| :--- | :--- |
| object | `map[string]any` |
| array | `[]any` |
| number | `float64` ⚠️ (не `int`) |
| string | `string` |
| `true`/`false` | `bool` |
| `null` | `nil` |

```go
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Envelope struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type User struct {
	Name string `json:"name"`
}

func main() {
	var env Envelope
	_ = json.Unmarshal([]byte(`{"type":"user","data":{"name":"bob"}}`), &env)
	var u User
	_ = json.Unmarshal(env.Data, &u) // разбираем по env.Type
	fmt.Println(env.Type, u.Name)    // user bob

	dec := json.NewDecoder(strings.NewReader(`{"name":"a","nmae":"typo"}`))
	dec.DisallowUnknownFields()
	fmt.Println(dec.Decode(&u)) // json: unknown field "nmae"
}
```

## CSV

Подробно: [02_csv.md](./01_theory/02_csv.md).

| Задача | Код |
| :--- | :--- |
| Читатель | `r := csv.NewReader(f)` |
| Разделитель / комментарии | `r.Comma = ';'`, `r.Comment = '#'` |
| Разное число полей в строках | `r.FieldsPerRecord = -1` (по умолчанию — как в первой строке, иначе ошибка) |
| Всё сразу | `records, err := r.ReadAll()` → `[][]string` |
| Построчно (большие файлы) | `rec, err := r.Read()` до `io.EOF` |
| Заголовок → индексы | `header, _ := r.Read(); idx[h] = i` |
| Писатель | `w := csv.NewWriter(f); defer w.Flush()` |
| Записать строку | `w.Write([]string{"id", "name"})` |
| Ошибка записи | `w.Flush(); err := w.Error()` |
| Кавычки | ставятся сами: `x,"say ""hi""","a,b"` (разделитель, `"`, перевод строки) |

## base64 и hex

Подробно: [03_base64.md](./01_theory/03_base64.md).

| Кодировка | Алфавит | Паддинг | Где | `[]byte{0xfb, 0xff}` |
| :--- | :--- | :--- | :--- | :--- |
| `base64.StdEncoding` | `+` `/` | `=` | бинарное в JSON, email | `+/8=` |
| `base64.URLEncoding` | `-` `_` | `=` | URL-параметры, токены | `-_8=` |
| `base64.RawStdEncoding` | `+` `/` | нет | строки без `=` | `+/8` |
| `base64.RawURLEncoding` | `-` `_` | нет | JWT | `-_8` |
| `hex.EncodeToString` | `0-9a-f` | — | ключи/хэши в логах | `fbff` |

| Задача | Код |
| :--- | :--- |
| Закодировать | `base64.StdEncoding.EncodeToString([]byte("hello"))` → `aGVsbG8=` |
| Декодировать | `b, err := base64.StdEncoding.DecodeString(s)` |
| Потоком | `enc := base64.NewEncoder(base64.StdEncoding, w); defer enc.Close()` |
| Размер | base64: `ceil(n/3)*4` (≈ +33 %); hex: ровно `2n` |
| Basic auth | `Authorization: Basic base64(user:pass)` |

## Ловушки

- ⚠️ `json.Unmarshal(data, u)` без `&` — ошибка `non-pointer`.
- ⚠️ Числа в `any` — `float64`: большие int64 теряют точность; тег `,string` или `json.Number` (`dec.UseNumber()`).
- ⚠️ `omitempty` не пропускает `time.Time{}` и пустые структуры — `omitzero` (1.24+).
- ⚠️ Без `DisallowUnknownFields` опечатки клиента в полях молча игнорируются.
- ⚠️ Не закрыт `resp.Body` — утечка соединений.
- ⚠️ Забыли `w.Flush()` у `csv.Writer` — хвост файла потерян; ошибки — только через `w.Error()`.
- ⚠️ CSV из Excel: BOM в начале (отрезать 3 байта), `;` вместо `,` в русской локали.
- ⚠️ `csv.Reader` не парсит типы — все поля строки, дальше `strconv`.
- ⚠️ Склейка CSV через `Sprintf` ломается на кавычках и переводах строк — только `csv.Writer`.
- ⚠️ `StdEncoding` и `URLEncoding` — разные алфавиты; строка без `=` — `Raw*Encoding`, иначе `illegal base64 data`.
- ⚠️ Забыли `Close()` у `base64.NewEncoder` — потерян хвост.
- ⚠️ base64 — кодирование, не шифрование и не обфускация.
