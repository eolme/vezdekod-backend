package entities

import (
	time "time"

	gorm "gorm.io/gorm"
)

type Author struct {
	UserId int `json:"user_id"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	ScreenName string `json:"screen_name"`

	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
