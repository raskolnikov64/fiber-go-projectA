package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserEntity struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Address   string         `json:"Address"`
	Phone     string         `json:"Phone"`
	Email     string         `json:"Email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
