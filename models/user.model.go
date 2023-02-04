package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FullName  string    `gorm:"size:255;not null;unique" json:"full_name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Role      int       `json:"role"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (q *User) CreateUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&q).Error
	if err != nil {
		return &User{}, err
	}
	return q, nil
}
