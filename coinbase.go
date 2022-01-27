package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl string = "https://api.coinbase.com/v2/prices/spot?currency="

type CoinbaseResponse struct {
	SpotPrice *SpotPrice `json:"data"`
}

type SpotPrice struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

func CallCoinbase(currency string) (*SpotPrice, error) {
	resp, err := http.Get(baseUrl + currency)
	if err != nil {
		// Deferring closes is bad, since they can error and that'll leave the error uncaught
		return nil, CloseAndCompactErrors(err, resp.Body.Close())
	}

	body, err := io.ReadAll(resp.Body)
	if err := CloseAndCompactErrors(err, resp.Body.Close()); err != nil {
		return nil, err
	}

	var data CoinbaseResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return &SpotPrice{}, err
	}
	return data.SpotPrice, nil
}

func CloseAndCompactErrors(upstream error, close error) error {
	if upstream == nil && close == nil {
		return nil
	}

	if upstream != nil && close == nil {
		return upstream
	}

	if upstream == nil && close != nil {
		return close
	}

	return fmt.Errorf("received error %s while closing after upstream error %s", close.Error(), upstream.Error())
}
