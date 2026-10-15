# Частые ошибки методов

## 1. Смешение value и pointer receiver

```go
func (u User)  Rename(n string) {}
func (u *User) Save() error { return nil }
```

Разные receiver в одном типе вводят в заблуждение и ломают соответствие интерфейсам. Выбери один стиль.

## 2. Метод с pointer receiver не вызывается у map-элемента

```go
m := map[string]User{"a": {}}
// m["a"].Save() // ошибка: нельзя взять адрес
```

Либо храни `*User`, либо получай через временную переменную и записывай обратно.

## 3. Копирование mutex

```go
type S struct{ mu sync.Mutex }
func (s S) Do() { s.mu.Lock() } // vet ругается: copies lock
```

`sync.Mutex` и `sync.WaitGroup` нельзя копировать. Используй `*S` receiver.

## 4. Nil-указатель

```go
func (u *User) Name() string { return u.name }

var u *User
u.Name() // panic: nil pointer dereference
```

Либо защищайся `if u == nil { return "" }`, либо гарантируй ненулевой указатель контрактом.

## 5. Receiver-имя конфликтует со стандартной переменной

```go
func (t *time.Time) ... // не разрешено на чужом типе
```

На встроенных и чужих типах методы определять нельзя — нужен свой именованный тип (`type MyTime time.Time`); alias (`type T = time.Time`) не поможет — это тот же чужой тип.

## 6. Перегрузка

```go
func (s S) Do(a int) {}
func (s S) Do(a string) {} // ошибка компиляции
```

Перегрузки по типам нет. Используй разные имена (`DoInt`, `DoString`) или `any`.

## 7. Method value захватывает receiver

```go
type Counter struct{ N int }

func (c *Counter) Inc()     { c.N++ }
func (c Counter) Show() int { return c.N }

var c Counter
inc := c.Inc   // pointer receiver: запомнен &c
show := c.Show // value receiver: запомнена копия c
c = Counter{N: 10}
inc()
fmt.Println(c.N, show()) // 11 0
```

Receiver вычисляется в момент создания method value: value receiver фиксирует **копию** (`show` не видит новых значений `c`), pointer receiver — **адрес** `&c` (`inc` меняет текущее содержимое `c`). Это редко мешает, но знать нужно.

## 8. String() с бесконечной рекурсией

```go
func (e Err) String() string {
    return fmt.Sprintf("%v", e) // бесконечный цикл: %v вызовет String()
}
```

Внутри `String()` не пользуйся `%v`/`%s` на себе — строй строку вручную.

## 9. Receiver на интерфейсе

Нельзя. Методы определяются на конкретном типе, не на интерфейсе.

## 10. Смешение методов и экспортированных полей

Если тип имеет и метод `Name()`, и поле `Name` — получится коллизия. Выбирай одно.
