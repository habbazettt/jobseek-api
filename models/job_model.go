package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Title           string         `gorm:"type:varchar(255);not null" json:"title"`
	Description     string         `gorm:"type:text;not null" json:"description"`
	CompanyID       uint           `gorm:"not null" json:"company_id"`
	Location        string         `gorm:"type:varchar(100);not null" json:"location"`
	Salary          int64          `gorm:"not null" json:"salary"`
	Currency        string         `gorm:"type:varchar(10);not null;default:'IDR'" json:"currency"`
	JobType         string         `gorm:"type:varchar(50);not null" json:"job_type"`
	Category        string         `gorm:"type:varchar(100);not null" json:"category"`
	ExperienceLevel string         `gorm:"type:varchar(50);not null" json:"experience_level"`
	Skills          []string       `gorm:"type:json;serializer:json" json:"skills"` // ✅ Menyimpan sebagai JSON
	Deadline        time.Time      `gorm:"not null" json:"deadline"`
	Status          string         `gorm:"type:varchar(20);not null;default:'open'" json:"status"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
