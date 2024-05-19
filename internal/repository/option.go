package repository

import "gorm.io/gorm"

type Option func(q *gorm.DB) *gorm.DB

func SubscriptionWithEmail(email string) Option {
	return func(q *gorm.DB) *gorm.DB {
		return q.Where("email = ?", email)
	}
}
