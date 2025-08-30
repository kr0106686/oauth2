package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `json:"id" gorm:"primarykey"`
	ProviderID string         `json:"provider_id"`
	Provider   string         `json:"provider"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Picture    string         `json:"picture"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
