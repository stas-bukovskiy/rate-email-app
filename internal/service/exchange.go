package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"usd-uah-emal-subsriber/internal/config"
	"usd-uah-emal-subsriber/pkg/errs"
)

type ExchangeService struct {
	baseURL   string
	apiKey    string
	baseCcy   string
	targetCcy string
}

func NewExchangeService(conf config.ExchangeAPIConfig) *ExchangeService {
	return &ExchangeService{
		baseURL:   "http://apilayer.net/api/live",
		apiKey:    conf.AccessKey,
		baseCcy:   conf.BaseCcy,
		targetCcy: conf.TargetCcy,
	}
}

func (s *ExchangeService) Rate() (float64, error) {
	resp, err := http.Get(fmt.Sprintf("%s?source=%s&currencies=%s&access_key=%s&format=1", s.baseURL, s.baseCcy,
		s.targetCcy, s.apiKey))
	if err != nil {
		return 0, errs.F("failed to fetch rate: %w", err)
	}

	defer resp.Body.Close()
	var result RateResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, errs.F("failed to decode response: %w", err)
	}

	if !result.Success {
		return 0, errs.F("failed to fetch rate, error code: %d, error info: %s", result.Error.Code, result.Error.Info)
	}

	return result.Quotes.UAH, nil
}
