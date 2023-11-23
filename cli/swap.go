package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// https://github.com/unielon-org/unielon-api/blob/main/swap.md

// Reserves

type ReserveData struct {
	LiquidityTotal string `json:"liquidity_total"`
	Reserve0       string `json:"reserve0"`
	Reserve1       string `json:"reserve1"`
	Tick0          string `json:"tick0"`
	Tick1          string `json:"tick1"`
}

func (r ReserveData) AmountOut() (float64, error) {
	fee := 0.03
	// Reserve0をfloat64に変換
	reserve0, err := strconv.ParseFloat(r.Reserve0, 64)
	if err != nil {
		return 0, fmt.Errorf("Error in parsing Reserve0: %w", err)
	}

	// Reserve1をfloat64に変換
	reserve1, err := strconv.ParseFloat(r.Reserve1, 64)
	if err != nil {
		return 0, fmt.Errorf("Error in parsing Reserve1: %w", err)
	}

	return (reserve0 / reserve1) * (1 - fee), nil
}

type SwapReservesResponse struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  ReserveData `json:"data"`
	Total int         `json:"total"`
}

func requestGetReserves(tick0, tick1 string) (*SwapReservesResponse, error) {
	path := "/swap/getreserves"
	requestBody, err := json.Marshal(map[string]string{
		"tick0": tick0,
		"tick1": tick1,
	})
	if err != nil {
		return nil, fmt.Errorf("Error in marshaling request body: %w", err)
	}

	response := &SwapReservesResponse{}
	err = sendRequest(path, requestBody, response)

	return response, err
}

func displayGetReservesResponse(response *SwapReservesResponse) {
	fmt.Printf("Get Reserves API Response:\nCode: %d\nMessage: %s\nTotal: %d\n", response.Code, response.Msg, response.Total)
	data := response.Data
	fmt.Printf("Liquidity Total: %s, Reserve0: %s, Reserve1: %s, Tick0: %s, Tick1: %s\n", data.LiquidityTotal, data.Reserve0, data.Reserve1, data.Tick0, data.Tick1)

	fmt.Println("----------")
}

// price
// no document

type TickerData struct {
	Tick      string  `json:"tick"`
	LastPrice float64 `json:"last_price"`
}

type SwapPriceResponse struct {
	Code  int          `json:"code"`
	Msg   string       `json:"msg"`
	Data  []TickerData `json:"data"`
	Total int          `json:"total"`
}

func requestSwapPriceApi() (*SwapPriceResponse, error) {
	path := "/swap/price"

	response := &SwapPriceResponse{}
	err := sendRequest(path, nil, response)

	return response, err
}

func displaySwapPriceResponse(response *SwapPriceResponse, filter string) {
	fmt.Printf("Swap Price API Response:\nCode: %d\nMessage: %s\nTotal Tickers: %d\n", response.Code, response.Msg, response.Total)
	for _, ticker := range response.Data {
		if filter != "" && ticker.Tick != filter {
			continue
		}
		fmt.Printf("Ticker: %s, Last Price: %f\n", ticker.Tick, ticker.LastPrice)
	}
	fmt.Println("----------")
}

// order history
// no doc

type SwapOrderResponse struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  []OrderData `json:"data"`
	Total int         `json:"total"`
}

func (r *SwapOrderResponse) CompleteOrder() []OrderData {
	var orders []OrderData

	for _, d := range r.Data {
		if d.OrderStatus == 0 {
			orders = append(orders, d)
		}
	}

	return orders
}

type OrderData struct {
	Tick0       string `json:"tick0"`
	Tick1       string `json:"tick1"`
	Amt0        int    `json:"amt0"`
	Amt1        int    `json:"amt1"`
	OrderStatus int    `json:"order_status"`
}

func CalculateAverageCost(orders []OrderData, tick string) float64 {
	var totalAmt0, totalAmt1 int

	for _, order := range orders {
		if order.Tick0 == "WDOGE(WRAPPED-DOGE)" {
			totalAmt0 += order.Amt0
			totalAmt1 += order.Amt1
		}
	}

	if totalAmt1 == 0 {
		return 0
	}

	return float64(totalAmt0) / float64(totalAmt1)
}

func requestSwapOrderApi(address string) (*SwapOrderResponse, error) {
	path := "/swap/order"

	requestBody, err := json.Marshal(map[string]string{
		"op":             "swap",
		"holder_address": address,
	})
	if err != nil {
		return nil, fmt.Errorf("Error in marshaling request body: %w", err)
	}

	response := &SwapOrderResponse{}
	err = sendRequest(path, requestBody, response)

	return response, err
}
