package response

import (
	"go-fiber-gorm/model/entity"
	"time"
)

type User struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Address   string      `json:"address"`
	Phone     string      `json:"phone"`
	Role      entity.Role `json:"role"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
