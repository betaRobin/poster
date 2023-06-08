package entity

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
