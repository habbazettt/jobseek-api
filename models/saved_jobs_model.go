package models

import (
	"time"

	"gorm.io/gorm"
)

type SavedJob struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	FreelancerID uint           `gorm:"not null" json:"freelancer_id"`
	JobID        uint           `gorm:"not null" json:"job_id"`
	CreatedAt    time.Time      `json:"created_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
