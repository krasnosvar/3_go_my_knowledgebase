// Package userclient — код под тестом для упражнения 02.
package userclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// ErrNotFound — пользователь не найден (HTTP 404).
var ErrNotFound = errors.New("user not found")

// User — ответ API.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Client — HTTP-клиент к сервису пользователей.
type Client struct {
	BaseURL string
	HTTP    *http.Client
}

// GetUser выполняет GET {BaseURL}/users/{id}.
func (c *Client) GetUser(ctx context.Context, id int) (User, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/users/%d", c.BaseURL, id), nil)
	if err != nil {
		return User{}, err
	}
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var u User
		if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
			return User{}, fmt.Errorf("decode user: %w", err)
		}
		return u, nil
	case http.StatusNotFound:
		return User{}, ErrNotFound
	default:
		return User{}, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}
}
