package response

import (
	"time"

	"gorm.io/gorm"
)

type AuctionHistory struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	AuctionID uint           `json:"-"`
	Auction   Auction        `json:"auction"`
	UserId    uint           `json:"-"`
	User      User           `json:"user"`
	Price     int            `json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
