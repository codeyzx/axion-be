package request

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
	Name           string           `json:"name"`
	LastPrice      int              `json:"last_price"`
	Status         Status           `json:"status"`
	BiddersCount   int              `json:"bidders_count"`
	ProductID      uuid.UUID        `json:"-"`
	Product        Products         `json:"product"`
	AuctionHistory []AuctionHistory `json:"-"`
	UserId         uint             `json:"-"`
	User           User             `json:"user"`
	BidderId       uint             `json:"-"`
	Bidder         User             `json:"bidder"`
	EndAt          string           `json:"end_at"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	DeletedAt      gorm.DeletedAt   `json:"-" gorm:"index,column:deleted_at"`
}

type Products struct {
	ID          uuid.UUID      `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	Image       string         `json:"image"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
