package main

import (
	"fmt"
)

type Request struct {
	Path  string
	Token string
}

type Response struct {
	Status int
	Body   string
}

type HandlerFunc func(req Request) Response

// AuthMiddleware принимает секретный ключ авторизации и возвращает замыкание,
// которое проверяет токен перед вызовом основного обработчика.
func AuthMiddleware(secretToken string, next HandlerFunc) HandlerFunc {
	return func(req Request) Response {
		if req.Token != secretToken {
			return Response{Status: 401, Body: "Unauthorized"}
		}
		// Передаем управление следующему обработчику
		return next(req)
	}
}

func main() {
	// Основной обработчик секретных данных
	secretHandler := func(req Request) Response {
		return Response{Status: 200, Body: "Секретные данные пользователя: {ID: 42, Balance: 1000}"}
	}

	// Оборачиваем обработчик авторизацией
	protectedHandler := AuthMiddleware("valid-session-token-123", secretHandler)

	// Запрос без токена
	resp1 := protectedHandler(Request{Path: "/api/profile", Token: "wrong-token"})
	fmt.Printf("Запрос 1 (плохой токен): Status %d, Body: %q\n", resp1.Status, resp1.Body)

	// Запрос с хорошим токеном
	resp2 := protectedHandler(Request{Path: "/api/profile", Token: "valid-session-token-123"})
	fmt.Printf("Запрос 2 (хороший токен): Status %d, Body: %q\n", resp2.Status, resp2.Body)
}
