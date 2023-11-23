package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var funcMap = map[string]func(){
	"tick-info": tickInfo,
	"orders":    orders,
	"buy-avg":   buyAverage,
}

func buyAverage() {
	address := os.Args[2]
	tick := os.Args[3]

	res, err := requestSwapOrderApi(address)
	if err != nil {
		log.Fatalln(err)
	}

	orders := res.CompleteOrder()

	price := CalculateAverageCost(orders, tick)

	fmt.Println(price)
}

func orders() {
	address := os.Args[2]
	res, err := requestSwapOrderApi(address)
	if err != nil {
		log.Fatalln(err)
	}

	b, err := json.Marshal(res.CompleteOrder())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(string(b))
}

func tickInfo() {
	tick := os.Args[2]
	res, err := requestSwapPriceApi()
	if err != nil {
		log.Fatalln(err)
	}

	wdogePrice := getPrice()

	var t TickerData

	for _, ticker := range res.Data {
		if ticker.Tick == tick {
			t = ticker
			break
		}
	}

	fmt.Println(tick)
	fmt.Println("$D:", t.LastPrice)
	fmt.Println("$ :", t.LastPrice*wdogePrice)
	fmt.Println("¥ :", t.LastPrice*wdogePrice*150)
}

func main() {
	if f, ok := funcMap[os.Args[1]]; ok {
		f()
	} else {
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: go run [arg]")
	fmt.Println("Where [arg] can be one of the following:")
	for key := range funcMap {
		fmt.Printf("- %s\n", key)
	}
}
