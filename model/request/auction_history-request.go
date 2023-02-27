package request

type AuctionHistoryCreateRequest struct {
	AuctionID uint `json:"auctionId"`
	UserId    uint `json:"userId"`
	Price     int  `json:"price"`
}

type AuctionHistoryUpdateRequest struct {
	Price int `json:"price"`
}
