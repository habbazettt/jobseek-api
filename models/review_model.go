package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	ReviewerID uint           `gorm:"not null" json:"reviewer_id"`
	ReviewedID uint           `gorm:"not null" json:"reviewed_id"`
	Rating     float64        `gorm:"not null" json:"rating"`
	Comment    string         `gorm:"type:text;not null" json:"comment"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
