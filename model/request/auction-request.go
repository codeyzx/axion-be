package request

import (
	"axion/model/entity"
)

type AuctionCreateRequest struct {
	Name         string `json:"name" form:"name" validate:"required"`
	LastPrice    int    `json:"lastPrice"`
	Status       string `json:"status"`
	BiddersCount int    `json:"biddersCount"`
	UserId       uint   `json:"userId" form:"userId" validate:"required"`
	EndAt        string `json:"endAt"`
	ProductName  string `json:"productName" validate:"required"`
	Description  string `json:"description"`
	Price        int    `json:"price" validate:"required"`
	Image        string `json:"image"`
	// Product      Product `json:"product" validate:"required"`
}

type AuctionUpdateRequest struct {
	Name      string        `json:"name"`
	LastPrice int           `json:"lastPrice"`
	UserId    uint          `json:"userId"`
	Status    entity.Status `json:"status"`
	EndAt     string        `json:"endAt"`
}
