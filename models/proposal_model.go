package models

import (
	"time"

	"gorm.io/gorm"
)

type Proposal struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	JobID        uint           `gorm:"not null" json:"job_id"`
	Job          Job            `gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE" json:"job"`
	FreelancerID uint           `gorm:"not null" json:"freelancer_id"`
	Freelancer   User           `gorm:"foreignKey:FreelancerID;constraint:OnDelete:CASCADE" json:"freelancer"`
	CoverLetter  string         `gorm:"type:text;not null" json:"cover_letter"`
	BidAmount    int64          `gorm:"not null" json:"bid_amount"`
	Currency     string         `gorm:"type:varchar(10);not null;default:'IDR'" json:"currency"`
	Status       string         `gorm:"type:varchar(20);not null;default:'pending'" json:"status"` // pending, accepted, rejected
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
