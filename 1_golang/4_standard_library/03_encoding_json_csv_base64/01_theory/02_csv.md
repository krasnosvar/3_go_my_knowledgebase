# encoding/csv

Пакет `encoding/csv` — чтение и запись CSV-файлов.

## Чтение

```go
f, _ := os.Open("data.csv")
defer f.Close()

r := csv.NewReader(f)
r.FieldsPerRecord = -1   // разрешить переменное число полей в строке
r.Comma = ';'            // разделитель (по умолчанию запятая)
r.Comment = '#'          // строки-комментарии

records, err := r.ReadAll()
for _, rec := range records {
    fmt.Println(rec[0], rec[1])
}
```

На больших файлах — читай построчно:

```go
for {
    rec, err := r.Read()
    if err == io.EOF { break }
    if err != nil { return err }
    process(rec)
}
```

## Запись

```go
f, _ := os.Create("out.csv")
defer f.Close()

w := csv.NewWriter(f)
defer w.Flush()

_ = w.Write([]string{"id","name","age"})
for _, u := range users {
    _ = w.Write([]string{fmt.Sprint(u.ID), u.Name, fmt.Sprint(u.Age)})
}
```

**`Flush()` обязателен** — без него данные застревают в буфере. `w.Error()` вернёт последнюю ошибку записи.

## Кавычки и экранирование

`encoding/csv` сам ставит кавычки, если поле содержит:

- разделитель (`,` или твой `Comma`);
- кавычку (она удваивается);
- перенос строки.

Поэтому безопасно писать любые строки, включая многострочные.

## Особенности

- UTF-8 без BOM работает сразу. Если файл с BOM от Excel — прочитай первые 3 байта и отбрось вручную.
- Excel может требовать `;` как разделитель в русской локали — переключай `Comma`.
- `csv.Reader` не парсит типы: все поля — строки. Конвертируй отдельно через `strconv`.

## С заголовком

Если первая строка — заголовок, читай отдельно:

```go
header, err := r.Read()
// map[name]index для удобного доступа
idx := map[string]int{}
for i, h := range header {
    idx[h] = i
}
```

Так зависимость от порядка колонок исчезает.

## Частые ошибки

- Забыть `defer w.Flush()` — половина файла потеряется.
- Не закрыть файл — на Windows файл остаётся «заблокирован».
- Использовать fmt.Sprintf + writeString вручную вместо `csv.Writer` — сломается на полях с кавычками или переводами строк.
- Одинаковое количество полей в записи обязательно (по умолчанию). Для «рваных» файлов выставь `FieldsPerRecord = -1`.

## Альтернативы

- `github.com/jszwec/csvutil` — struct-to-CSV с тегами.
- Для больших датасетов — колоночные форматы (Parquet) будут сильно эффективнее.
