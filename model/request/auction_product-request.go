package request

type Product struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Price       int    `json:"price" validate:"required"`
	Image       string `json:"image"`
}
