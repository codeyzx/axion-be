package entity

import (
	"time"

	"gorm.io/gorm"
)

type History struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Log       string         `json:"log"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
