package request

type AuctionHistoryCreateRequest struct {
	AuctionID uint `json:"auction_id"`
	UserId    uint `json:"user_id"`
	Price     int  `json:"price"`
}
