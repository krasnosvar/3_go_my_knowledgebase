// Упражнение 05: UDP-сервер, суммирующий числа (хардкор).
//
// Задача: реализовать runSumServer — UDP-сервер, который принимает
// ОДНУ датаграмму со строкой чисел через пробел (например "3 4 10"),
// суммирует их и отправляет отправителю строку с результатом ("17").
// После обработки одной датаграммы сервер завершается.
//
// Сервер запускается в горутине; клиент и самопроверка уже написаны.
// Помните: UDP сохраняет границы сообщений — одна датаграмма = одна
// строка чисел, разбирать поток по '\n' не нужно.
//
// Запуск: go run ./03_exercises/05_udp_sum_server/main.go
package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

// Держим strconv живым в заготовке.
var _ = strconv.Atoi

// runSumServer читает одну датаграмму из conn, суммирует числа из неё
// и отправляет строку-результат обратно отправителю. По завершении
// закрывает канал done.
func runSumServer(conn net.PacketConn, done chan<- struct{}) {
	defer close(done)
	// TODO:
	//  1) прочитать датаграмму: n, addr, err := conn.ReadFrom(buf);
	//  2) поля := strings.Fields(string(buf[:n]));
	//  3) просуммировать числа через strconv.Atoi (нечисловые пропускать
	//     или считать ошибкой — на ваше усмотрение);
	//  4) отправить strconv.Itoa(sum) обратно: conn.WriteTo(..., addr).
}

func main() {
	fmt.Println("--- Проверка 05_udp_sum_server ---")

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

	// Дедлайн, чтобы нереализованный сервер не подвесил проверку,
	// а честно упал в FAIL по таймауту чтения.
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
