// Упражнение 07: простой строчный протокол поверх TCP.
//
// Задача: реализовать handleConn — обработчик одного TCP-соединения,
// который читает от клиента одну строку (до '\n'), приводит её к
// ВЕРХНЕМУ регистру и отправляет обратно, тоже завершая '\n'.
//
// Сервер уже поднят на 127.0.0.1:0, клиент уже написан — ваша задача
// только в handleConn. TCP это поток байтов, поэтому разграничивайте
// сообщения по '\n'.
//
// Запуск: go run ./03_exercises/07_tcp_line_protocol/main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Держим strings живым в заготовке (пригодится в решении).
var _ = strings.ToUpper

// handleConn читает от клиента одну строку, преобразует её в верхний
// регистр и отправляет обратно (с завершающим '\n').
func handleConn(conn net.Conn) {
	defer conn.Close()
	// TODO:
	//  1) прочитать строку до '\n' через
	//   bufio.NewReader(conn).ReadString('\n');
	//  2) обрезать перевод строки, применить strings.ToUpper;
	//  3) записать результат обратно: conn.Write([]byte(result + "\n")).
}

func main() {
	fmt.Println("--- Проверка 07_tcp_line_protocol ---")

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	defer ln.Close()

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		handleConn(conn)
	}()

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("hello\n")); err != nil {
		log.Fatalf("client write: %v", err)
	}

	reply, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("FAIL: ошибка чтения ответа: %v\n", err)
		fmt.Println("FAIL")
		return
	}

	got := strings.TrimRight(reply, "\n")
	if got == "HELLO" {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: got=%q, want=%q\n", got, "HELLO")
		fmt.Println("FAIL")
	}
}
