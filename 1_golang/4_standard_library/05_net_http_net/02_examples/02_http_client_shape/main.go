package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}))
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil { // сначала проверяем ошибку: при err != nil resp == nil
		fmt.Println("request failed:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}
