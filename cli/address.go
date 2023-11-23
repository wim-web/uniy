package main

import (
	"encoding/json"
	"fmt"
)

// https://github.com/unielon-org/unielon-api/blob/main/drc20.md

type AddressTickResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Tick string `json:"tick"`
		Amt  int    `json:"amt"`
	} `json:"data"`
	Total int `json:"total"`
}

func requestAddressTickApi(address, tick string) (*AddressTickResponse, error) {
	path := "/drc20/address/tick"
	requestBody, err := json.Marshal(map[string]string{
		"receive_address": address,
		"tick":            tick,
	})
	if err != nil {
		return nil, fmt.Errorf("Error in marshaling request body: %w", err)
	}

	response := &AddressTickResponse{}
	err = sendRequest(path, requestBody, &response)

	return response, err
}

func displayAddressTickResponse(response *AddressTickResponse) {
	fmt.Printf("Address Tick API Response:\nCode: %d\nMessage: %s\nTicker: %s, Amount: %d\n", response.Code, response.Msg, response.Data.Tick, response.Data.Amt)
	fmt.Println("----------")
}
