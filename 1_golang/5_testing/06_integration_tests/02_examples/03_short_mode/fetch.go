// Package fetch — пример: тест с реальным HTTP-сервером и режимом -short.
package fetch

import (
	"io"
	"net/http"
)

// Body скачивает тело ответа по url.
func Body(c *http.Client, url string) (string, error) {
	resp, err := c.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	return string(b), err
}
