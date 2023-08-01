package coingecko

import (
	"io"
	"net/http"
)

func Ping() (string, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/ping")
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
