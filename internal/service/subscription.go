package service

import (
	"context"
	"net/http"
	"usd-uah-emal-subsriber/internal/localecode"
	"usd-uah-emal-subsriber/internal/model"
	"usd-uah-emal-subsriber/internal/repository"
	"usd-uah-emal-subsriber/pkg/errs"
)

type SubscriptionService struct {
	subRepo SubscriptionRepository
}

func NewSubscriptionService(subRepo SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{subRepo: subRepo}
}

func (s *SubscriptionService) CreateSubscription(ctx context.Context, email string) error {
	exist, err := s.subRepo.Exist(ctx, repository.SubscriptionWithEmail(email))
	if err != nil {
		return errs.F("failed to check if subscription exist, email: %s, error: %w", email, err)
	}

	if exist {
		return errs.F("subscription already exist, email: %s", email).SetCode(localecode.CodeEmailExist).
			SetStatusCode(http.StatusConflict)
	}

	sub := &model.Subscription{
		Email: email,
	}

	if err = s.subRepo.Create(ctx, sub); err != nil {
		return errs.F("failed to create subscription, email: %s, error: %w", email, err)
	}

	return nil
}

func (s *SubscriptionService) Subscriptions(ctx context.Context) ([]model.Subscription, error) {
	subs, err := s.subRepo.Subscriptions(ctx)
	if err != nil {
		return nil, errs.F("failed to get subscriptions: %w", err)
	}

	return subs, nil
}
