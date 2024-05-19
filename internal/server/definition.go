package server

import "context"

type (
	ExchangeService interface {
		Rate() (float64, error)
	}
	SubscriptionService interface {
		CreateSubscription(ctx context.Context, email string) error
	}
)
