package main

import (
	"fmt"
	"net/http/httptest"
)

func main() {
	req := httptest.NewRequest("GET", "/users?page=2", nil)
	fmt.Println(req.URL.Query().Get("page"))
}
