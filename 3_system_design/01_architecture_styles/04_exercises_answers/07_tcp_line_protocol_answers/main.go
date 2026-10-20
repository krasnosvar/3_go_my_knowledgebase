// Упражнение 07: строчный протокол поверх TCP (Решение).
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	line, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Printf("server read: %v", err)
		return
	}
	result := strings.ToUpper(strings.TrimRight(line, "\n"))
	if _, err := conn.Write([]byte(result + "\n")); err != nil {
		log.Printf("server write: %v", err)
	}
}

func main() {
	fmt.Println("--- Проверка 07_tcp_line_protocol (Решение) ---")

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
