package request

import (
	"axion/model/entity"
)

type AuctionCreateRequest struct {
	LastPrice    int     `json:"last_price"`
	Status       string  `json:"status"`
	BiddersCount int     `json:"bidders_count"`
	UserId       uint    `json:"user_id" form:"user_id" validate:"required"`
	Product      Product `json:"product" validate:"required"`
	EndAt        string  `json:"end_at"`
}

type AuctionUpdateRequest struct {
	LastPrice int           `json:"last_price"`
	UserId    uint          `json:"user_id"`
	Status    entity.Status `json:"status"`
	EndAt     string        `json:"end_at"`
}
