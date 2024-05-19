package cron

import (
	"context"
	"usd-uah-emal-subsriber/internal/model"
)

type (
	SenderService interface {
		SendEmail(to []string, subject, body string) error
	}

	ExchangeService interface {
		Rate() (float64, error)
	}

	SubscriptionService interface {
		Subscriptions(ctx context.Context) ([]model.Subscription, error)
	}
)
