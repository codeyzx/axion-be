package request

import (
	"time"

	"gorm.io/gorm"
)

type AuctionHistoryCreateRequest struct {
	AuctionID uint `json:"auctionId"`
	UserId    uint `json:"userId"`
	Price     int  `json:"price"`
}

type AuctionHistoryUpdateRequest struct {
	Price int `json:"price"`
}

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
