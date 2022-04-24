package entities

import (
	time "time"

	gorm "gorm.io/gorm"
)

type Meme struct {
	PhotoId    int    `json:"photo_id"`
	PhotoOwner int    `json:"photo_owner"`
	PhotoUrl   string `json:"photo_url"`

	AuthorId int    `json:"-"`
	Author   Author `gorm:"foreignKey:AuthorId;references:ID" json:"author"`

	Likes int `json:"likes"`
	Prio  int `json:"-" gorm:"default:0"`

	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
