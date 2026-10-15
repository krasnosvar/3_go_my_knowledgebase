# Все об ООП в Go

Go не является классическим объектно-ориентированным языком, но поддерживает многие ООП-концепции через свои уникальные механизмы: структуры, методы, интерфейсы и встраивание.

## 1. Структуры (Structs) — аналог классов

В Go нет классов, вместо них используются структуры — коллекции именованных полей.

```go
// Определение структуры
type Person struct {
    Name string
    Age  int
    City string
}

// Создание экземпляра
person1 := Person{Name: "Алексей", Age: 30, City: "Москва"}
person2 := Person{"Мария", 25, "Санкт-Петербург"} // Без имен полей
person3 := Person{Name: "Иван"} // Остальные поля получат нулевые значения

// Доступ к полям
fmt.Println(person1.Name) // Алексей
person1.Age = 31
```

### Анонимные структуры

Полезны для одноразовых структур данных.

```go
user := struct {
    ID   int
    Name string
}{
    ID:   1,
    Name: "Анонимный",
}
```

## 2. Методы — функции, привязанные к типам

Методы добавляют поведение к структурам.

```go
type Rectangle struct {
    Width, Height float64
}

// Метод с получателем по значению (value receiver)
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Метод с получателем по указателю (pointer receiver)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Использование
rect := Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area()) // 50

rect.Scale(2)
fmt.Println(rect.Area()) // 200
```

### Получатель по значению vs по указателю

| Тип получателя | Когда использовать | Изменяет оригинал |
|----------------|-------------------|-------------------|
| По значению `(r Rectangle)` | Метод только читает данные, структура маленькая | Нет |
| По указателю `(r *Rectangle)` | Метод изменяет данные, структура большая | Да |

**Правило:** Если хотя бы один метод использует указатель, используй указатели для всех методов типа (для консистентности).

## 3. Конструкторы

В Go нет встроенных конструкторов, но принято создавать функции `New...()`.

```go
type User struct {
    ID       int
    Username string
    Email    string
}

// Конструктор
func NewUser(username, email string) *User {
    return &User{
        ID:       generateID(), // Какая-то логика генерации ID
        Username: username,
        Email:    email,
    }
}

// Использование
user := NewUser("john_doe", "john@example.com")
```

### Конструктор с валидацией

```go
func NewUser(username, email string) (*User, error) {
    if username == "" {
        return nil, fmt.Errorf("username не может быть пустым")
    }
    if !strings.Contains(email, "@") {
        return nil, fmt.Errorf("некорректный email")
    }
    
    return &User{
        ID:       generateID(),
        Username: username,
        Email:    email,
    }, nil
}
```

## 4. Инкапсуляция — видимость через регистр

В Go нет модификаторов доступа (public/private). Видимость определяется регистром первой буквы:

- **Заглавная буква** — экспортируется (public)
- **Строчная буква** — не экспортируется (private)

```go
package bank

type Account struct {
    Owner   string  // Экспортируется (доступно вне пакета)
    balance float64 // Не экспортируется (доступно только внутри пакета)
}

// Геттер для приватного поля
func (a *Account) Balance() float64 {
    return a.balance
}

// Сеттер с валидацией
func (a *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return fmt.Errorf("сумма должна быть положительной")
    }
    a.balance += amount
    return nil
}
```

## 5. Встраивание (Embedding) — композиция вместо наследования

Go **не поддерживает наследование**, но поддерживает композицию через встраивание.

```go
type Engine struct {
    Power int
}

func (e Engine) Start() {
    fmt.Println("Двигатель запущен")
}

type Car struct {
    Engine      // Встраивание (анонимное поле)
    Brand string
    Model string
}

// Использование
car := Car{
    Engine: Engine{Power: 150},
    Brand:  "Toyota",
    Model:  "Camry",
}

car.Start()              // Вызов метода встроенного типа
fmt.Println(car.Power)   // Прямой доступ к полю встроенного типа
fmt.Println(car.Engine.Power) // Альтернативный доступ
```

### Множественное встраивание

```go
type GPS struct {
    Latitude  float64
    Longitude float64
}

func (g GPS) Location() string {
    return fmt.Sprintf("Lat: %.2f, Lon: %.2f", g.Latitude, g.Longitude)
}

type SmartCar struct {
    Engine
    GPS
    Brand string
}

car := SmartCar{
    Engine: Engine{Power: 200},
    GPS:    GPS{Latitude: 55.75, Longitude: 37.61},
    Brand:  "Tesla",
}

car.Start()                    // Метод Engine
fmt.Println(car.Location())    // Метод GPS
```

### Переопределение методов

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println("Какой-то звук")
}

type Dog struct {
    Animal
}

// Переопределение метода
func (d Dog) Speak() {
    fmt.Println("Гав!")
}

dog := Dog{Animal: Animal{Name: "Бобик"}}
dog.Speak()        // Гав! (вызов метода Dog)
dog.Animal.Speak() // Какой-то звук (вызов метода Animal)
```

## 6. Интерфейсы — полиморфизм

Интерфейс — набор методов. Тип реализует интерфейс **неявно**, просто имея все его методы.

```go
type Speaker interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Гав!"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Мяу!"
}

// Функция принимает любой тип, реализующий Speaker
func MakeSound(s Speaker) {
    fmt.Println(s.Speak())
}

// Использование
dog := Dog{Name: "Бобик"}
cat := Cat{Name: "Мурка"}

MakeSound(dog) // Гав!
MakeSound(cat) // Мяу!
```

### Пустой интерфейс interface{} (any)

Может хранить значение любого типа.

```go
var i interface{}

i = 42
i = "hello"
i = []int{1, 2, 3}

// С Go 1.18+ можно использовать any
var x any
x = 3.14
```

### Приведение типов (Type Assertion)

```go
var i interface{} = "hello"

// Простое приведение (паникует, если тип не совпадает)
s := i.(string)
fmt.Println(s) // hello

// Безопасное приведение с проверкой
s, ok := i.(string)
if ok {
    fmt.Println("Это строка:", s)
} else {
    fmt.Println("Это не строка")
}

// Type switch
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Целое число: %d\n", v)
    case string:
        fmt.Printf("Строка: %s\n", v)
    case bool:
        fmt.Printf("Булево: %t\n", v)
    default:
        fmt.Printf("Неизвестный тип: %T\n", v)
    }
}
```

## 7. Стандартные интерфейсы

Go имеет множество полезных встроенных интерфейсов.

### fmt.Stringer

Позволяет определить строковое представление типа.

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d лет)", p.Name, p.Age)
}

person := Person{Name: "Иван", Age: 30}
fmt.Println(person) // Иван (30 лет)
```

### io.Reader и io.Writer

Основа для работы с потоками данных.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Пример использования
func CopyData(dst io.Writer, src io.Reader) error {
    _, err := io.Copy(dst, src)
    return err
}
```

### error

```go
type error interface {
    Error() string
}

// Кастомная ошибка
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Использование
func ValidateAge(age int) error {
    if age < 0 {
        return ValidationError{
            Field:   "Age",
            Message: "не может быть отрицательным",
        }
    }
    return nil
}
```

## 8. Композиция интерфейсов

Интерфейсы можно комбинировать.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// Композиция интерфейсов
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

## 9. Указатели и интерфейсы

**Важное правило:** Если метод определен для указателя, только указатель реализует интерфейс.

```go
type Incrementer interface {
    Increment()
}

type Counter struct {
    value int
}

// Метод с получателем-указателем
func (c *Counter) Increment() {
    c.value++
}

// Правильно
var i Incrementer
counter := &Counter{} // Указатель
i = counter

// Ошибка компиляции!
counter2 := Counter{}
// i = counter2 // Counter не реализует Incrementer, только *Counter
```

## 10. Шаблон проектирования: Dependency Injection

Использование интерфейсов для инъекции зависимостей.

```go
// Интерфейс для работы с базой данных
type UserRepository interface {
    GetByID(id int) (*User, error)
    Save(user *User) error
}

// Сервис зависит от интерфейса, а не от конкретной реализации
type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*User, error) {
    return s.repo.GetByID(id)
}

// Реализация для PostgreSQL
type PostgresUserRepo struct {
    // ...
}

func (p *PostgresUserRepo) GetByID(id int) (*User, error) {
    // Логика работы с PostgreSQL
}

func (p *PostgresUserRepo) Save(user *User) error {
    // Логика сохранения
}

// Использование
repo := &PostgresUserRepo{}
service := NewUserService(repo)
user, err := service.GetUser(1)
```

## 11. Generics (с Go 1.18+)

Обобщенное программирование позволяет писать функции и типы, работающие с разными типами данных.

```go
import "cmp"

// Обобщенная функция; cmp.Ordered — стандартный constraint (Go 1.21+),
// раньше использовали golang.org/x/exp/constraints.Ordered
func Min[T cmp.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

fmt.Println(Min(10, 20))       // 10
fmt.Println(Min(3.14, 2.71))   // 2.71
fmt.Println(Min("abc", "xyz")) // abc

// Обобщенный тип
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

// Использование
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)

stringStack := Stack[string]{}
stringStack.Push("hello")
stringStack.Push("world")
```

## 12. Важные практики ООП в Go

### 1. Предпочитай композицию наследованию

```go
// Плохо: попытка имитировать наследование
type Animal struct {
    Name string
}

type Dog struct {
    Animal
    Breed string
}

// Хорошо: явная композиция
type Dog struct {
    Name  string
    Breed string
}
```

### 2. Маленькие интерфейсы

"Чем меньше интерфейс, тем больше его мощь" — Rob Pike

```go
// Хорошо: маленькие, сфокусированные интерфейсы
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Комбинируй при необходимости
type ReadWriter interface {
    Reader
    Writer
}
```

### 3. Принимай интерфейсы, возвращай структуры

```go
// Хорошая практика
func ProcessData(r io.Reader) (*Result, error) {
    // r — интерфейс (гибкость)
    // *Result — конкретный тип (ясность)
}

// Плохая практика
func ProcessData(r *os.File) (io.Reader, error) {
    // Слишком конкретный вход, слишком абстрактный выход
}
```

### 4. Используй указатели для больших структур

```go
type LargeStruct struct {
    Data [1000000]int
}

// Хорошо: передача по указателю
func Process(ls *LargeStruct) {
    // Избегаем копирования
}

// Плохо: передача по значению
func Process(ls LargeStruct) {
    // Копирование огромного массива!
}
```

### 5. Нулевые значения должны быть полезны

```go
// Хорошо: структура работает с нулевыми значениями
type Buffer struct {
    data []byte
}

func (b *Buffer) Write(p []byte) {
    b.data = append(b.data, p...) // Работает, даже если data == nil
}

var buf Buffer
buf.Write([]byte("hello")) // OK!
```

## 13. Сравнение с классическим ООП

| Концепция | Классическое ООП | Go |
|-----------|-----------------|-----|
| Классы | class | struct |
| Наследование | extends | Встраивание (embedding) |
| Полиморфизм | Интерфейсы/наследование | Интерфейсы |
| Инкапсуляция | public/private/protected | Регистр первой буквы |
| Конструкторы | constructor() | Функции New...() |
| Деструкторы | destructor/finalize | Нет (defer для очистки) |
| Перегрузка методов | Поддерживается | Не поддерживается |
| Абстрактные классы | abstract class | Интерфейсы |

## Источники

- [SOLID и DRY в Go (Habr, OTUS)](https://habr.com/ru/companies/otus/articles/786314/)
