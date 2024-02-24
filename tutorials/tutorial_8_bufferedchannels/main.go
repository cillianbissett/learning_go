package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_PRICE float32 = 100.0

func main() {
	var priceChannel = make(chan string)
	var websites = []string{"amazon", "ebay", "bestbuy", "newegg", "walmart"}
	for i := range websites {
		go checkPrice(websites[i], priceChannel)
	}
	sendMessage(priceChannel)
}

func checkPrice(website string, priceChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		price := rand.Float32() * 20
		if price < MAX_PRICE {
			priceChannel <- website
			break
		}
	}
}

func sendMessage(priceChannel chan string) {
	// and if we were scanning multiple sites, could use a select statement
	fmt.Printf("\nFound a deal on %s\n", <-priceChannel)
}
