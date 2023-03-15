package response

import (
	"axion/model/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Auction struct {
	ID             uint                    `json:"id" gorm:"primaryKey"`
	Name           string                  `json:"name"`
	LastPrice      int                     `json:"last_price"`
	Status         entity.Status           `json:"status"`
	BiddersCount   int                     `json:"bidders_count"`
	ProductID      uuid.UUID               `json:"-" gorm:"primaryKey"`
	Product        Product                 `json:"product"`
	AuctionHistory []entity.AuctionHistory `json:"auction_history"`
	UserId         uint                    `json:"-"`
	User           User                    `json:"user"`
	BidderId       uint                    `json:"-"`
	Bidder         User                    `json:"bidder"`
	EndAt          string                  `json:"end_at"`
	CreatedAt      time.Time               `json:"created_at"`
	UpdatedAt      time.Time               `json:"updated_at"`
	DeletedAt      gorm.DeletedAt          `json:"-" gorm:"index,column:deleted_at"`
}
