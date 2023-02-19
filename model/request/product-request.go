package request

type ProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description" nullable:"true"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
}
