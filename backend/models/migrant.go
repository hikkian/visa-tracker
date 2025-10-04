package models

import (
	"time"
)

type Migrant struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	FullName    string     `json:"full_name"`
	Passport    string     `json:"passport"`
	Nationality string     `json:"nationality"`
	EntryDate   time.Time  `json:"entry_date"`
	StayType    string     `json:"stay_type"`
	VisaExpiry  *time.Time `json:"visa_expiry,omitempty"`
	AllowedDays *int       `json:"allowed_days,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
