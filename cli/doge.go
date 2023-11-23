package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type CoinGeckoResponse struct {
	Dogecoin struct {
		USD float64 `json:"usd"`
	} `json:"dogecoin"`
}

func getPrice() float64 {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=dogecoin&vs_currencies=usd"

	// HTTPリクエストを行う
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %s\n", err)
	}
	defer resp.Body.Close()

	// レスポンスのボディを読み込む
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s\n", err)
	}

	// JSONデータを解析
	var data CoinGeckoResponse
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalf("Error decoding JSON: %s\n", err)
	}

	// 結果を出力
	return data.Dogecoin.USD
}
