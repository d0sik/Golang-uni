package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RateResponse struct {
	Rate  float64 `json:"rate"`
	Error string  `json:"error"`
}

type ExchangeService struct {
	BaseURL string
	Client  *http.Client
}

func NewExchangeService(url string) *ExchangeService {
	return &ExchangeService{
		BaseURL: url,
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (s *ExchangeService) GetRate(from, to string) (float64, error) {
	url := fmt.Sprintf("%s/convert?from=%s&to=%s", s.BaseURL, from, to)

	resp, err := s.Client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result RateResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("%s", result.Error)
	}

	return result.Rate, nil
}
