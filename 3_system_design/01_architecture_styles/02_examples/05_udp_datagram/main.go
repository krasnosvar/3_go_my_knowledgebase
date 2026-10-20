// Пример 05: UDP — обмен одной датаграммой.
//
// Сервер поднимается в отдельной горутине на 127.0.0.1:0 (случайный
// порт), клиент отправляет одну датаграмму, сервер отвечает эхом,
// клиент печатает ответ — и программа завершается.
//
// Демонстрирует: net.ListenPacket / ReadFrom / WriteTo, отсутствие
// соединения (connectionless) и сохранение границ сообщений: одна
// отправка = одна датаграмма (в отличие от потока байтов TCP).
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// UDP-сокет сервера на случайном порту.
	serverConn, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("listen packet: %v", err)
	}
	defer serverConn.Close()

	addr := serverConn.LocalAddr()
	fmt.Printf("UDP-сервер слушает на %s\n", addr)

	done := make(chan struct{})

	// Сервер: принимает одну датаграмму и отвечает отправителю.
	go func() {
		defer close(done)

		buf := make([]byte, 1024)
		// ReadFrom возвращает данные и адрес отправителя (соединения нет).
		n, from, err := serverConn.ReadFrom(buf)
		if err != nil {
			log.Printf("server read: %v", err)
			return
		}
		fmt.Printf("Сервер получил датаграмму: %q от %s\n", buf[:n], from)

		// Отвечаем именно тому адресу, откуда пришёл запрос.
		if _, err := serverConn.WriteTo(buf[:n], from); err != nil {
			log.Printf("server write: %v", err)
		}
	}()

	// Клиент: net.Dial для UDP не открывает соединение, а лишь
	// фиксирует адрес назначения, позволяя использовать Write/Read.
	clientConn, err := net.Dial("udp", addr.String())
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer clientConn.Close()

	if _, err := clientConn.Write([]byte("hello udp")); err != nil {
		log.Fatalf("client write: %v", err)
	}

	buf := make([]byte, 1024)
	n, err := clientConn.Read(buf)
	if err != nil {
		log.Fatalf("client read: %v", err)
	}
	fmt.Printf("Клиент получил эхо: %q\n", buf[:n])

	<-done
	fmt.Println("Готово, программа завершается.")
}
