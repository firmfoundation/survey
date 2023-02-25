package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Indicator struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	SurveyID  uuid.UUID `json:"survey_id" gorm:"type:uuid;default:uuid_generate_v4();not null"`
	Survey    Survey    `json:"survey"`
	Weight    float64   `json:"weight"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (i *Indicator) CreateIndicator(db *gorm.DB) (*Indicator, error) {

	var err error
	err = db.Debug().Create(&i).Error
	if err != nil {
		return &Indicator{}, err
	}
	return i, nil
}

func (q *Indicator) GetAllIndicators(db *gorm.DB) (*[]Indicator, error) {
	var err error
	indicator := []Indicator{}
	err = db.Debug().Limit(100).Find(&indicator).Error
	if err != nil {
		return &[]Indicator{}, err
	}
	return &indicator, err
}
