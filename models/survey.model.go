package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Survey struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Survey) SaveSurvey(db *gorm.DB) (*Survey, error) {

	var err error
	err = db.Debug().Create(&s).Error
	if err != nil {
		return &Survey{}, err
	}
	return s, nil
}
