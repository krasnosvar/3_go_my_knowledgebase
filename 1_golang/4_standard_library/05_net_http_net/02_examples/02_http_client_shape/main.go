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
	resp, _ := http.Get(ts.URL)
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}
