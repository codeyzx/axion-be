package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	Open   Status = "Open"
	Closed Status = "Closed"
)

type Auction struct {
	ID             uint             `json:"id" gorm:"primaryKey"`
	LastPrice      int              `json:"last_price"`
	Status         Status           `json:"status"`
	BiddersCount   int              `json:"bidders_count"`
	ProductID      uuid.UUID        `json:"product_id" gorm:"unique"`
	Product        Product          `json:"-"`
	AuctionHistory []AuctionHistory `json:"auction_history"`
	UserId         uint             `json:"user_id"`
	User           User             `json:"-"`
	EndAt          string           `json:"end_at"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	DeletedAt      gorm.DeletedAt   `json:"-" gorm:"index,column:deleted_at"`
}
