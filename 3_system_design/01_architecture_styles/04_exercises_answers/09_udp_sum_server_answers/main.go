// Упражнение 09: UDP-сервер, суммирующий числа (Решение).
package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func runSumServer(conn net.PacketConn, done chan<- struct{}) {
	defer close(done)

	buf := make([]byte, 1024)
	n, addr, err := conn.ReadFrom(buf)
	if err != nil {
		log.Printf("server read: %v", err)
		return
	}

	sum := 0
	for _, field := range strings.Fields(string(buf[:n])) {
		v, err := strconv.Atoi(field)
		if err != nil {
			// Нечисловые токены пропускаем.
			continue
		}
		sum += v
	}

	if _, err := conn.WriteTo([]byte(strconv.Itoa(sum)), addr); err != nil {
		log.Printf("server write: %v", err)
	}
}

func main() {
	fmt.Println("--- Проверка 09_udp_sum_server (Решение) ---")

	serverConn, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("listen packet: %v", err)
	}
	defer serverConn.Close()

	done := make(chan struct{})
	go runSumServer(serverConn, done)

	client, err := net.Dial("udp", serverConn.LocalAddr().String())
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer client.Close()

	if _, err := client.Write([]byte("3 4 10")); err != nil {
		log.Fatalf("client write: %v", err)
	}

	_ = client.SetReadDeadline(time.Now().Add(2 * time.Second))

	buf := make([]byte, 1024)
	n, err := client.Read(buf)
	if err != nil {
		fmt.Printf("FAIL: ошибка чтения ответа: %v\n", err)
		fmt.Println("FAIL")
		return
	}

	got := strings.TrimSpace(string(buf[:n]))
	<-done
	if got == "17" {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: got=%q, want=%q\n", got, "17")
		fmt.Println("FAIL")
	}
}
