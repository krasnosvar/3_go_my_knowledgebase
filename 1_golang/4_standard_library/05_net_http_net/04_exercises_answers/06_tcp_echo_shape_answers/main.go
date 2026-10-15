// Упражнение 06: TCP echo-сервер на пакете net (Решение).
package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

// handleConn отвечает эхом на каждую строку клиента.
func handleConn(c net.Conn) {
	defer c.Close()
	sc := bufio.NewScanner(c)
	for sc.Scan() {
		if _, err := fmt.Fprintf(c, "echo: %s\n", sc.Text()); err != nil {
			return
		}
	}
}

func main() {
	fmt.Println("--- Проверка 06_tcp_echo_shape (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("FAIL: listen:", err)
		return
	}
	defer ln.Close()
	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				return // слушатель закрыт
			}
			go handleConn(c)
		}
	}()

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		fmt.Println("FAIL: dial:", err)
		return
	}
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(2 * time.Second))

	r := bufio.NewReader(conn)
	for _, msg := range []string{"hello", "go net"} {
		fmt.Fprintf(conn, "%s\n", msg)
		line, err := r.ReadString('\n')
		want := "echo: " + msg + "\n"
		check(err == nil && line == want, fmt.Sprintf("на %q ожидалось %q, получено %q (err=%v)", msg, want, line, err))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
