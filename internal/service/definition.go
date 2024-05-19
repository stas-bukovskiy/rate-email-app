package service

import (
	"context"

	"usd-uah-emal-subsriber/internal/model"
	"usd-uah-emal-subsriber/internal/repository"
)

type (
	SubscriptionRepository interface {
		Exist(ctx context.Context, opts ...repository.Option) (bool, error)
		Create(ctx context.Context, sub *model.Subscription) error
		Subscriptions(ctx context.Context) ([]model.Subscription, error)
	}
)

type RateResponse struct {
	Success   bool   `json:"success"`
	Terms     string `json:"terms"`
	Privacy   string `json:"privacy"`
	Timestamp int    `json:"timestamp"`
	Source    string `json:"source"`
	Quotes    struct {
		UAH float64 `json:"USDUAH"`
	} `json:"quotes"`
	Error struct {
		Code string `json:"code"`
		Info string `json:"info"`
	}
}
