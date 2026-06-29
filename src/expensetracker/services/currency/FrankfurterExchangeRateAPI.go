package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type frankfurterResponse struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

type FrankfurterExchangeRateAPI struct{}

func (api FrankfurterExchangeRateAPI) GetExchangeRate(from, to string) (float64, error) {
	endpoint := fmt.Sprintf(
		"https://api.frankfurter.dev/v1/latest?base=%s&symbols=%s",
		url.QueryEscape(from),
		url.QueryEscape(to),
	)

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(endpoint)
	if err != nil {
		return 0, fmt.Errorf("fetch exchange rate: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("exchange rate API returned status %d", resp.StatusCode)
	}

	var data frankfurterResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, fmt.Errorf("decode exchange rate response: %w", err)
	}

	rate, ok := data.Rates[to]
	if !ok {
		return 0, fmt.Errorf("no rate for %s -> %s", from, to)
	}

	return rate, nil
}
