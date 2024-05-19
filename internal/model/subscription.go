package model

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	Email string `gorm:"unique"`
}
