# encoding/json

Пакет `encoding/json` — работа с JSON. Один из самых используемых в backend Go.

## Marshal / Unmarshal

```go
type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age,omitempty"`
    Email string `json:"email,omitempty"`
}

data, err := json.Marshal(u)        // []byte, err
err = json.Unmarshal(data, &u)
```

`omitempty` — пропустить пустые значения (0, "", nil, empty slice/map).

`-` — полностью игнорировать поле: `json:"-"`.

## Indent

```go
data, _ := json.MarshalIndent(u, "", "  ")
```

Для CLI-вывода, отладки, сохранения файлов.

## Encoder / Decoder

Для работы с потоками (файл, HTTP body):

```go
dec := json.NewDecoder(r)
dec.DisallowUnknownFields() // строгая проверка
for {
    var v Item
    if err := dec.Decode(&v); err == io.EOF { break }
    // handle v
}

enc := json.NewEncoder(w)
enc.SetIndent("", "  ")
_ = enc.Encode(u)
```

Потоковый API дешевле по памяти на больших объёмах.

## Теги

- `json:"name"` — имя в JSON.
- `json:",omitempty"` — пропустить пустое.
- `json:"-"` — игнорировать.
- `json:",string"` — сериализовать как строку (полезно для больших int64 в JavaScript).

## Кастомизация

```go
func (t MyTime) MarshalJSON() ([]byte, error) { ... }
func (t *MyTime) UnmarshalJSON(data []byte) error { ... }
```

Реализуешь `json.Marshaler` / `json.Unmarshaler` — получаешь полный контроль над форматом.

## interface{} / any

При декодировании неизвестной структуры:

- JSON object → `map[string]any`;
- JSON array → `[]any`;
- JSON number → `float64` (внимание — не int!);
- JSON string → `string`;
- JSON null → `nil`;
- JSON bool → `bool`.

## json.RawMessage

Отложенный парсинг куска JSON:

```go
type Envelope struct {
    Type string          `json:"type"`
    Data json.RawMessage `json:"data"`
}
```

Сначала смотрим `Type`, потом парсим `Data` в подходящую структуру. Экономит аллокации, когда не всё нужно разбирать.

## Частые ошибки

- Забыть указатель в `Unmarshal(data, &u)`. Без `&` — panic или ошибка.
- `float64` как все числа — при работе с больших int нужен tag `,string` или отдельный тип.
- Забыть `omitempty` для полей с нулевыми значениями — исходящий JSON раздувается.
- Не закрывать `resp.Body` в HTTP, перед чтением — утечка соединений.
- Decode без `DisallowUnknownFields` — пропускает опечатки в клиентском JSON.

## Производительность

Для экстремально горячих путей (> 100k req/s) есть быстрые альтернативы:

- `goccy/go-json` — drop-in, быстрее.
- `bytedance/sonic` — ещё быстрее, amd64 only.

Но это крайность. В 99% случаев stdlib достаточно.
