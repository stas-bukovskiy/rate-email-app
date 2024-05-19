package repository

import (
	"context"

	"gorm.io/gorm"

	"usd-uah-emal-subsriber/internal/model"
	"usd-uah-emal-subsriber/pkg/errs"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (s *SubscriptionRepository) Exist(ctx context.Context, opts ...Option) (bool, error) {
	stmt := s.db.WithContext(ctx).Model(&model.Subscription{})
	for _, opt := range opts {
		stmt = opt(stmt)
	}

	var count int64
	if err := stmt.Count(&count).Error; err != nil {
		return false, errs.F("failed to check if subscription exist: %w", err)
	}

	return count > 0, nil
}

func (s *SubscriptionRepository) Create(ctx context.Context, sub *model.Subscription) error {
	if err := s.db.WithContext(ctx).Create(sub).Error; err != nil {
		return errs.F("failed to create subscription: %w", err)
	}
	return nil
}

func (s *SubscriptionRepository) Subscriptions(ctx context.Context) ([]model.Subscription, error) {
	var subs []model.Subscription
	if err := s.db.WithContext(ctx).Find(subs).Error; err != nil {
		return nil, errs.F("failed to query subs: %w", err)
	}

	return subs, nil
}
