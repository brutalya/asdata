package coingecko

import (
	"io"
	"net/http"
)

func GetResponse(url string) (string, error) {
	resp, err := http.Get(url)
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

// PING
func Ping() (string, error) {
	return GetResponse("https://api.coingecko.com/api/v3/ping")
}

// GENERAL
func GetCoinList() (string, error) {
	return GetResponse("https://api.coingecko.com/api/v3/coins/list")
}

func GetExchangeList() (string, error) {
	return GetResponse("https://api.coingecko.com/api/v3/exchanges")
}

// COINS
func GetCoin(coinId string) (string, error) {
	return GetResponse("https://api.coingecko.com/api/v3/coins/" + coinId)
}

func GetCoinHistory(coinId string, date string) (string, error) {
	return GetResponse("https://api.coingecko.com/api/v3/coins/" + coinId + "/history?date=" + date)
}

func GetCoinMarket(coinId string, vsCurrency string) (string, error) {
	return GetResponse("https://api.coingecko.com/api/v3/coins/markets?vs_currency=" + vsCurrency + "&ids=" + coinId + "&order=market_cap_desc&per_page=100&page=1&sparkline=false&locale=en")
}

func GetCoinMarketChart(coinId string, vsCurrency string, days string) (string, error) {
	return GetResponse("https://api.coingecko.com/api/v3/coins/" + coinId + "/market_chart?vs_currency=" + vsCurrency + "&days=" + days)
}

func GetCoinMarketHistory(coinId string, date string) (string, error) {
	return GetResponse("https://api.coingecko.com/api/v3/coins/" + coinId + "/market_chart?vs_currency=usd&days=" + date)
}

func GetCoinTickers(coinId string) (string, error) {
	return GetResponse("https://api.coingecko.com/api/v3/coins/" + coinId + "/tickers")
}

func GetCoinStatusUpdates(coinId string) (string, error) {
	return GetResponse("https://api.coingecko.com/api/v3/coins/" + coinId + "/status_updates")
}

// CONTRACTS
func GetCoinContract(coinId string) (string, error) {
	return GetResponse("https://api.coingecko.com/api/v3/coins/" + coinId + "/contract")
}

func GetCoinContractMarketChart(coinId string, vsCurrency string, days string) (string, error) {
	return GetResponse("https://api.coingecko.com/api/v3/coins/" + coinId + "/contract/" + vsCurrency + "/market_chart?vs_currency=" + vsCurrency + "&days=" + days)
}
