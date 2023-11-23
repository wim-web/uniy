package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const BASE_URL = "https://unielon.com/v3"

// method: POST
func sendRequest(apiPath string, requestBody []byte, response interface{}) error {
	// HTTPクライアントの作成
	client := &http.Client{}

	url, err := url.JoinPath(BASE_URL, apiPath)
	if err != nil {
		return fmt.Errorf("Error in joining url: %w", err)
	}

	// POSTリクエストの作成
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("Error in creating request: %w", err)
	}

	// リクエストヘッダーの設定
	req.Header.Set("Content-Type", "application/json")

	// リクエストの実行
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error in sending request: %w", err)
	}
	defer resp.Body.Close()

	// レスポンスデータの読み込み
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("Error in decoding response: %w", err)
	}

	return nil
}
