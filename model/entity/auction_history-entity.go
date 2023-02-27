package entity

import (
	"time"

	"gorm.io/gorm"
)

// AuctionHistory model info
// @Description Auction history information
// @Description with auction id, user id, price, created at, updated at, and deleted at
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
