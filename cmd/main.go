package main

import (
	"fmt"
	"log"

	"github.com/brutalya/asdata/pkg/coingecko"
)

func main() {
	response, err := coingecko.Ping()
	if err != nil {
		log.Fatalf("Error pinging API: %v", err)
	}

	fmt.Println(response)

	res, e := coingecko.GetCoinMarketChart("bitcoin", "usd", "1")
	if e != nil {
		log.Fatalf("Error pinging API: %v", err)
	}

	fmt.Println(res)
}
