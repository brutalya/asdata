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

func GetCoinList() (string, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
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

func GetCoinMarketChart(coinId string, vsCurrency string, days string) (string, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/" + coinId + "/market_chart?vs_currency=" + vsCurrency + "&days=" + days)
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

func GetCoinMarket(coinId string, vsCurrency string) (string, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=" + vsCurrency + "&ids=" + coinId + "&order=market_cap_desc&per_page=100&page=1&sparkline=false&locale=en")
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
