// Пример 04: TCP echo-сервер и клиент.
//
// Сервер поднимается в отдельной горутине на 127.0.0.1:0 (случайный
// свободный порт), клиент подключается, отправляет строку, получает
// её обратно (эхо), печатает результат — и программа завершается.
//
// Демонстрирует: net.Listen("tcp", ...), Accept, соединение как поток
// байтов, разграничение сообщений через '\n' (bufio.ReadString).
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Слушаем TCP на случайном порту локального интерфейса.
	// ":0" -> ОС сама выбирает свободный порт (нет конфликтов).
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	defer ln.Close()

	addr := ln.Addr().String()
	fmt.Printf("TCP-сервер слушает на %s\n", addr)

	// Сигнал о том, что сервер обработал одно соединение и завершился.
	done := make(chan struct{})

	// Сервер в отдельной горутине: принимает одно соединение,
	// читает строку и отправляет её обратно.
	go func() {
		defer close(done)

		conn, err := ln.Accept()
		if err != nil {
			log.Printf("server accept: %v", err)
			return
		}
		defer conn.Close()

		// TCP — поток байтов, поэтому читаем до разделителя '\n'.
		line, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("server read: %v", err)
			return
		}
		fmt.Printf("Сервер получил: %q\n", line)

		// Эхо: отправляем обратно ту же строку.
		if _, err := conn.Write([]byte(line)); err != nil {
			log.Printf("server write: %v", err)
		}
	}()

	// Клиент: подключается к серверу, шлёт строку, читает эхо.
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer conn.Close()

	msg := "hello tcp\n"
	if _, err := conn.Write([]byte(msg)); err != nil {
		log.Fatalf("client write: %v", err)
	}

	reply, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatalf("client read: %v", err)
	}
	fmt.Printf("Клиент получил эхо: %q\n", reply)

	// Дожидаемся завершения серверной горутины и выходим.
	<-done
	fmt.Println("Готово, программа завершается.")
}
