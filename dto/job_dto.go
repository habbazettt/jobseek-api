package dto

import "time"

// JobRequest digunakan untuk validasi input saat membuat/mengupdate pekerjaan
type JobRequest struct {
	Title           string    `json:"title" binding:"required"`
	Description     string    `json:"description" binding:"required"`
	Location        string    `json:"location" binding:"required"`
	Salary          int64     `json:"salary" binding:"required,min=0"`
	Currency        string    `json:"currency" binding:"required,oneof=IDR USD EUR"`
	JobType         string    `json:"job_type" binding:"required,oneof=full-time part-time freelance"`
	Category        string    `json:"category" binding:"required"`
	ExperienceLevel string    `json:"experience_level" binding:"required,oneof=junior mid senior"`
	Skills          []string  `json:"skills" binding:"required"`
	Deadline        time.Time `json:"deadline" binding:"required"`
}

// JobResponse digunakan untuk response API
type JobResponse struct {
	ID              uint      `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	CompanyID       uint      `json:"company_id"`
	Location        string    `json:"location"`
	Salary          int64     `json:"salary"`
	Currency        string    `json:"currency"`
	JobType         string    `json:"job_type"`
	Category        string    `json:"category"`
	ExperienceLevel string    `json:"experience_level"`
	Skills          []string  `json:"skills"`
	Deadline        time.Time `json:"deadline"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}

// JobFilterRequest digunakan untuk filtering & pagination di GetJobs()
type JobFilterRequest struct {
	Category        string `form:"category"`
	Location        string `form:"location"`
	ExperienceLevel string `form:"experience_level"`
	Page            int    `form:"page" binding:"required,min=1"`
	Limit           int    `form:"limit" binding:"required,min=1,max=50"`
}
