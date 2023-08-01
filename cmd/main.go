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
}
