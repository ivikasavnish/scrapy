package utils

import (
	"io"
	"net/http"
	"time"
)

func FetchPage(url string) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
