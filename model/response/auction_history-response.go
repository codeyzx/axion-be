package response

import (
	"time"

	"gorm.io/gorm"
)

type AuctionHistory struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	AuctionID   uint           `json:"auction_id"`
	AuctionName string         `json:"auction_name"`
	UserId      uint           `json:"user_id"`
	UserName    string         `json:"user_name"`
	Price       int            `json:"price"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
