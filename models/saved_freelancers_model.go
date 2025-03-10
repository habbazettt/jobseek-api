package models

import (
	"time"

	"gorm.io/gorm"
)

type SavedFreelancer struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CompanyID    uint           `gorm:"not null" json:"company_id"`
	FreelancerID uint           `gorm:"not null" json:"freelancer_id"`
	CreatedAt    time.Time      `json:"created_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
