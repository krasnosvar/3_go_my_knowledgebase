package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	IsPremium bool   `json:"is_premium,omitempty"` // опустить, если false/пусто
}

func main() {
	// 1. Маршалинг (Struct -> JSON)
	u1 := User{ID: 1, Username: "alex_go", IsPremium: true}
	jsonData1, _ := json.Marshal(u1)
	fmt.Printf("User 1 JSON: %s\n", jsonData1)

	u2 := User{ID: 2, Username: "den_go", IsPremium: false}
	jsonData2, _ := json.Marshal(u2)
	fmt.Printf("User 2 JSON (c omitempty): %s\n", jsonData2)

	// 2. Унмаршалинг (JSON -> Struct)
	rawJSON := `{"id": 42, "username": "senior_dev", "is_premium": true}`
	var u3 User
	_ = json.Unmarshal([]byte(rawJSON), &u3)
	fmt.Printf("Восстановленная структура: %+v\n", u3)
}
