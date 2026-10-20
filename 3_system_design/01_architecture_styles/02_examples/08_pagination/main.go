// Пример 08: пагинация коллекции через query-параметры limit/offset.
//
// Сервер (httptest, случайный порт на 127.0.0.1) хранит 137 элементов
// и отдаёт их страницами: GET /items?limit=20&offset=40. В ответе —
// сама страница плюс метаданные (limit, offset, total, next). Клиент
// проходит несколько страниц и завершается.
//
// Демонстрирует: r.URL.Query(), разбор и валидацию limit/offset,
// вычисление ссылки next, формирование пагинированного JSON-ответа.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"time"
)

const total = 137 // всего элементов в "базе"

// page — форма пагинированного ответа.
type page struct {
	Items  []int  `json:"items"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Total  int    `json:"total"`
	Next   string `json:"next,omitempty"` // ссылка на следующую страницу (пусто на последней)
}

// parseIntParam читает целочисленный query-параметр с дефолтом и нижней
// границей.
func parseIntParam(q url.Values, name string, def, min int) int {
	s := q.Get(name)
	if s == "" {
		return def
	}
	v, err := strconv.Atoi(s)
	if err != nil || v < min {
		return def
	}
	return v
}

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	limit := parseIntParam(q, "limit", 20, 1)  // по умолчанию 20, минимум 1
	offset := parseIntParam(q, "offset", 0, 0) // по умолчанию 0, минимум 0

	if limit > 100 {
		limit = 100 // защита: не отдаём слишком большие страницы
	}

	// Вырезаем окно [offset, offset+limit) из "базы" [0..total).
	items := make([]int, 0, limit)
	for i := offset; i < offset+limit && i < total; i++ {
		items = append(items, i)
	}

	p := page{Items: items, Limit: limit, Offset: offset, Total: total}
	// Ссылка next есть, только если дальше ещё что-то осталось.
	if offset+limit < total {
		p.Next = fmt.Sprintf("/items?limit=%d&offset=%d", limit, offset+limit)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(p)
}

func main() {
	srv := httptest.NewServer(http.HandlerFunc(handler))
	defer srv.Close()
	fmt.Printf("Сервер с пагинацией слушает на %s (total=%d)\n\n", srv.URL, total)

	client := &http.Client{Timeout: 5 * time.Second}

	// Пройдём коллекцию страницами по 50, идя по ссылке next.
	next := "/items?limit=50&offset=0"
	pageNum := 1
	for next != "" {
		p := fetch(client, srv.URL+next)
		fmt.Printf("Страница %d: offset=%d limit=%d, элементов=%d (первый=%d, последний=%d)\n",
			pageNum, p.Offset, p.Limit, len(p.Items), first(p.Items), last(p.Items))
		next = p.Next
		pageNum++
	}

	fmt.Println("\nПоследняя страница достигнута (next пуст). Готово, программа завершается.")
}

func fetch(c *http.Client, url string) page {
	resp, err := c.Get(url)
	if err != nil {
		log.Fatalf("GET %s: %v", url, err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	var p page
	if err := json.Unmarshal(b, &p); err != nil {
		log.Fatalf("decode: %v", err)
	}
	return p
}

func first(xs []int) int {
	if len(xs) == 0 {
		return -1
	}
	return xs[0]
}

func last(xs []int) int {
	if len(xs) == 0 {
		return -1
	}
	return xs[len(xs)-1]
}
