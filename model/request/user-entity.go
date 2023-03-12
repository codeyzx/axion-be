package request

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	Users    Role = "Users"
	Admin    Role = "Admin"
	Operator Role = "Operator"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"-" gorm:"column:password"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	Role      Role           `json:"role"`
	Auctions  []Auction      `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
