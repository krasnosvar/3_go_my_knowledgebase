// Упражнение 04: sync.Pool для буферов.
//
// Задача: реализуйте getBuf/putBuf поверх bufPool и render(name):
//   - getBuf достаёт *bytes.Buffer из пула;
//   - putBuf обязательно сбрасывает буфер (Reset) перед возвратом в пул —
//     иначе следующий пользователь получит чужие данные;
//   - render берёт буфер, пишет "Hello, <name>!", возвращает строку и
//     возвращает буфер в пул.
//
// Помните: пул может быть очищен на любом GC — это кэш, а не хранилище.
//
// Запуск: go run ./03_exercises/04_pool_usage/main.go
package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

func getBuf() *bytes.Buffer {
	// TODO: верните буфер из пула.
	return new(bytes.Buffer)
}

func putBuf(b *bytes.Buffer) {
	// TODO: сбросьте буфер и верните его в пул.
	bufPool.Put(b)
}

func render(name string) string {
	// TODO: getBuf, запись "Hello, <name>!", копия строки, putBuf.
	return ""
}

func main() {
	fmt.Println("--- Проверка 04_pool_usage ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(render("Go") == "Hello, Go!", fmt.Sprintf("render(\"Go\"): ожидалось \"Hello, Go!\", получено %q", render("Go")))
	check(render("Pool") == "Hello, Pool!", "повторный render не должен содержать хвост прошлого вызова")

	dirty := getBuf()
	dirty.WriteString("secret")
	putBuf(dirty)
	again := getBuf()
	check(again.Len() == 0, fmt.Sprintf("буфер из пула должен быть пустым, а в нём %q", again.String()))
	putBuf(again)

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
