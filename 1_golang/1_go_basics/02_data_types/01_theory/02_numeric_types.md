# Числовые типы

## Целые числа со знаком

```go
int8    // -128 до 127
int16   // -32768 до 32767
int32   // -2147483648 до 2147483647
int64   // -9223372036854775808 до 9223372036854775807
int     // зависит от платформы (32 или 64 бита)
```

## Целые числа без знака

```go
uint8   // 0 до 255
uint16  // 0 до 65535
uint32  // 0 до 4294967295
uint64  // 0 до 18446744073709551615
uint    // зависит от платформы
```

## Специальные целочисленные типы

```go
byte    // alias для uint8
rune    // alias для int32 (Unicode code point)
uintptr // беззнаковое целое для хранения указателя
```

## Нюансы целых чисел

- `int` и `int32` — РАЗНЫЕ типы (нужно явное приведение).
- Нулевое значение: `0`.
- Переполнение: wraparound без ошибки.

```go
var x int32 = 100
var y int = 100
// x == y   // ошибка компиляции
int(x) == y // правильно

var z int8 = 127
z++
fmt.Println(z) // -128 (переполнение)
```

## Числа с плавающей точкой

```go
float32 // IEEE-754 32-bit
float64 // IEEE-754 64-bit
```

- Нулевое значение: `0.0`.
- По умолчанию литералы — `float64`.
- Нельзя сравнивать на равенство напрямую:

```go
var f float64 = 0.1 + 0.2
fmt.Println(f == 0.3) // false

const epsilon = 1e-9
if math.Abs(f-0.3) < epsilon {}
```

## Комплексные числа

```go
complex64   // float32 real + float32 imag
complex128  // float64 real + float64 imag

c := 1 + 2i
fmt.Println(real(c), imag(c)) // 1 2
```

## Размеры типов (64-bit)

```
int8/uint8:   1 byte
int16/uint16: 2 bytes
int32/uint32: 4 bytes
int64/uint64: 8 bytes
int/uint:     8 bytes (platform dependent)
float32:      4 bytes
float64:      8 bytes
complex64:    8 bytes
complex128:  16 bytes
```

## См. также

- [Как компьютер хранит числа](../../../../other/03_number_representation.md) — двоичное представление, дополнительный код, IEEE 754, переполнение.
