package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:current_timestamp()"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}

func NewUser(username string, password string) *User {
	userId, _ := uuid.NewRandom()
	newUser := User{
		ID:        userId,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
	return &newUser
}
